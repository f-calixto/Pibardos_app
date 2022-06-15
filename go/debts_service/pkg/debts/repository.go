package debts

import (
	// std lib
	"database/sql"

	// internal
	"github.com/coding-kiko/debts_service/pkg/errors"
	"github.com/coding-kiko/debts_service/pkg/log"
)

var (
	getCurrentStatusQuery = `SELECT status FROM debt_requests WHERE id = $1`
	createRequestQuery    = `INSERT INTO debt_requests(id, group_id, lender_id, borrower_id, date, 
						  	 description, amount, status) Values($1, $2, $3, $4, $5, $6, $7, $8)`
	alterDebtRequestQuery = `UPDATE debt_requests SET status = $1 WHERE id = $2`
	getDebtRequestQuery   = `SELECT id, group_id, lender_id, borrower_id, date, description, amount, 
						   	  status FROM debt_requests WHERE id = $1`
	getDebtAmountQuery = `SELECT amount FROM debts WHERE group_id = $1 AND lender_id = $2 
					 	  AND borrower_id = $3`
	upsertDebtAmountQuery = `INSERT INTO debts(amount, group_id, lender_id, borrower_id) VALUES($1, $2, $3, $4)
							 ON CONFLICT ON CONSTRAINT debts_group_id_lender_id_borrower_id_key DO 
							 UPDATE SET amount = $1 WHERE debts.group_id = $2 AND debts.lender_id = $3 
							 AND debts.borrower_id = $4`
	getSentRequestsQuery = `SELECT id, group_id, lender_id, borrower_id, date, description, amount, status
							FROM debt_requests WHERE group_id = $1 AND lender_id = $2`
	getReceivedRequestsQuery = `SELECT id, group_id, lender_id, borrower_id, date, description, amount, status
								FROM debt_requests WHERE group_id = $1 AND borrower_id = $2`
	getGroupDebtsQuery = `SELECT borrower_id, lender_id, amount FROM debts WHERE group_id = $1 AND amount > 0`
	cancelDebtQuery    = `UPDATE debts SET amount = GREATEST(amount - $1, 0) WHERE group_id = $2 AND 
						  lender_id = $3 AND borrower_id = $4`
	getUpdatedDebtsQuery = `SELECT lender_id, amount FROM debts WHERE group_id = $1 
							AND amount > 0 AND borrower_id = $2`
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: logger,
	}
}

type Repository interface {
	CreateRequest(req DebtRequest) error
	AcceptDebt(id string) (DebtRequest, error)
	RejectDebt(id string) (DebtRequest, error)
	GetDebtRequest(id string) (DebtRequest, error)
	BalanceDebt(req BalanceDebtRequest) error
	GetSentRequests(req GetRequestsRequest) ([]DebtRequest, error)
	GetReceivedRequests(req GetRequestsRequest) ([]DebtRequest, error)
	GetGroupDebts(groupId string) ([]Borrower, error)
	CancelDebt(req CancelDebtRequest) (Borrower, error)
}

// cancel debt tottally or partially, if amount is bigger than debt it becomes 0
func (r *repo) CancelDebt(req CancelDebtRequest) (Borrower, error) {
	_, err := r.db.Exec(cancelDebtQuery, req.Amount, req.GroupId, req.LenderId, req.BorrowerId)
	if err != nil {
		return Borrower{}, errors.NewNotFound("request not found")
	}

	//get updated debts for the borrower
	rows, err := r.db.Query(getUpdatedDebtsQuery, req.GroupId, req.BorrowerId)
	if err != nil {
		return Borrower{}, errors.NewNotFound("borrower debts not found")
	}
	defer rows.Close()

	var debts Borrower
	debts.BorrowerId = req.BorrowerId
	for rows.Next() {
		var d Debt

		err := rows.Scan(&d.LenderId, &d.Amount)
		if err != nil {
			return Borrower{}, errors.NewNotFound("group debts not found")
		}

		debts.Debts = append(debts.Debts, d)
	}

	return debts, nil
}

func (r *repo) GetGroupDebts(groupId string) ([]Borrower, error) {
	var debts []Borrower
	var tmp = make(map[string][]Debt)

	rows, err := r.db.Query(getGroupDebtsQuery, groupId)
	if err != nil {
		return []Borrower{}, errors.NewNotFound("group debts not found")
	}
	defer rows.Close()

	for rows.Next() {
		var debt Debt
		var borrowerId string

		err := rows.Scan(&borrowerId, &debt.LenderId, &debt.Amount)
		if err != nil {
			return []Borrower{}, errors.NewNotFound("group debts not found")
		}
		// create map in the form of: map[borrower1: [debt1, debt2] borrower2: [debt1]]
		tmp[borrowerId] = append(tmp[borrowerId], debt)
	}

	// create group debts json struct
	for k, v := range tmp {
		borrower := Borrower{
			BorrowerId: k,
			Debts:      v,
		}
		debts = append(debts, borrower)
	}

	return debts, nil
}

func (r *repo) GetSentRequests(req GetRequestsRequest) ([]DebtRequest, error) {
	var sent []DebtRequest

	rows, err := r.db.Query(getSentRequestsQuery, req.GroupId, req.UserId)
	if err != nil {
		return []DebtRequest{}, errors.NewNotFound("user requests not found")
	}
	defer rows.Close()

	for rows.Next() {
		var req DebtRequest

		err := rows.Scan(&req.Id, &req.GroupId, &req.LenderId, &req.BorrowerId, &req.Date, &req.Description, &req.Amount, &req.Status)
		if err != nil {
			return []DebtRequest{}, errors.NewNotFound("user requests not found")
		}
		sent = append(sent, req)
	}
	return sent, nil
}

func (r *repo) GetReceivedRequests(req GetRequestsRequest) ([]DebtRequest, error) {
	var received []DebtRequest

	rows, err := r.db.Query(getReceivedRequestsQuery, req.GroupId, req.UserId)
	if err != nil {
		return []DebtRequest{}, errors.NewNotFound("user requests not found")
	}
	defer rows.Close()

	for rows.Next() {
		var req DebtRequest

		err := rows.Scan(&req.Id, &req.GroupId, &req.LenderId, &req.BorrowerId, &req.Date, &req.Description, &req.Amount, &req.Status)
		if err != nil {
			return []DebtRequest{}, errors.NewNotFound("user requests not found")
		}
		received = append(received, req)
	}
	return received, nil
}

// balance debt between lender and borrower
func (r *repo) BalanceDebt(req BalanceDebtRequest) error {
	var amount int64

	err := r.db.QueryRow(getDebtAmountQuery, req.GroupId, req.BorrowerId, req.LenderId).Scan(&amount)
	if err != nil {
		// debt was not created previously
		if err == sql.ErrNoRows {
			amount = 0
		} else {
			return errors.NewNotFound("debt not found")
		}
	}

	difference := amount - req.Amount
	if difference <= 0 {
		_, err := r.db.Exec(upsertDebtAmountQuery, 0, req.GroupId, req.BorrowerId, req.LenderId)
		if err != nil {
			return errors.NewNotFound("request not found")
		}

		_, err = r.db.Exec(upsertDebtAmountQuery, -difference, req.GroupId, req.LenderId, req.BorrowerId)
		if err != nil {
			return errors.NewNotFound("request not found")
		}
	}
	if difference > 0 {
		_, err := r.db.Exec(upsertDebtAmountQuery, difference, req.GroupId, req.BorrowerId, req.LenderId)
		if err != nil {
			return errors.NewNotFound("request not found")
		}
	}

	return nil
}

// change debt request staus to 0 (rejected)
func (r *repo) RejectDebt(id string) (DebtRequest, error) {
	var currentStatus int64

	err := r.db.QueryRow(getCurrentStatusQuery, id).Scan(&currentStatus)
	if err != nil {
		return DebtRequest{}, errors.NewNotFound("request not found")
	}
	if currentStatus != 2 {
		return DebtRequest{}, errors.NewInvalidUpdate("request was already accepted or rejected")
	}

	_, err = r.db.Exec(alterDebtRequestQuery, 0, id)
	if err != nil {
		return DebtRequest{}, errors.NewNotFound("request not found")
	}

	// get updated request
	debtRequest, err := r.GetDebtRequest(id)
	if err != nil {
		return DebtRequest{}, err
	}
	return debtRequest, nil
}

// change debt request staus to 1 (accepted)
func (r *repo) AcceptDebt(id string) (DebtRequest, error) {
	var currentStatus int64

	err := r.db.QueryRow(getCurrentStatusQuery, id).Scan(&currentStatus)
	if err != nil {
		return DebtRequest{}, errors.NewNotFound("request not found")
	}
	if currentStatus != 2 {
		return DebtRequest{}, errors.NewInvalidUpdate("request was already accepted or rejected")
	}

	_, err = r.db.Exec(alterDebtRequestQuery, 1, id)
	if err != nil {
		return DebtRequest{}, errors.NewNotFound("request not found")
	}

	// get updated request
	debtRequest, err := r.GetDebtRequest(id)
	if err != nil {
		return DebtRequest{}, err
	}
	return debtRequest, nil
}

// get single debt request
func (r *repo) GetDebtRequest(id string) (DebtRequest, error) {
	var debtRequest DebtRequest

	err := r.db.QueryRow(getDebtRequestQuery, id).Scan(&debtRequest.Id, &debtRequest.GroupId, &debtRequest.LenderId, &debtRequest.BorrowerId, &debtRequest.Date, &debtRequest.Description, &debtRequest.Amount, &debtRequest.Status)
	if err != nil {
		return DebtRequest{}, errors.NewNotFound("request not found")
	}
	return debtRequest, nil
}

func (r *repo) CreateRequest(req DebtRequest) error {
	_, err := r.db.Exec(createRequestQuery, req.Id, req.GroupId, req.LenderId, req.BorrowerId, req.Date, req.Description, req.Amount, req.Status)
	if err != nil {
		return err
	}
	return nil
}
