package debts

import (
	// std lib
	"database/sql"

	// internal

	"github.com/coding-kiko/debts_service/pkg/errors"
	"github.com/coding-kiko/debts_service/pkg/log"
)

var (
	getDebtQuery    = `SELECT FROM debts WHERE id = $1`
	createDebtQuery = `INSERT INTO debts(id, group_id, lender_id, borrower_id, date, 
						  	 description, amount, status) Values($1, $2, $3, $4, $5, $6, $7, $8)`
	patchDebtQuery = `UPDATE debts SET status = $1 WHERE id = $2 AND borrower_id = $3
						AND status = 2 RETURNING *`
	cancelDebtQuery = `UPDATE debts SET status = $1 WHERE id = $2 AND lender_id = $3
						AND status = 2 RETURNING *`
	getReceivedQuery = `SELECT FROM debts WHERE borrower_id = $1 AND group_id = $2`
	getSentQuery     = `SELECT FROM debts WHERE lender_id = $1 AND group_id = $2`
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
	CreateDebt(debt Debt) error
	AcceptDebt(req PatchDebtRequest) (Debt, error)
	RejectDebt(req PatchDebtRequest) (Debt, error)
	CancelDebt(req PatchDebtRequest) (Debt, error)
	GetSentDebts(req GetDebtsRequest) ([]Debt, error)
	GetReceivedDebts(req GetDebtsRequest) ([]Debt, error)
}

func (r *repo) CreateDebt(debt Debt) error {
	_, err := r.db.Exec(createDebtQuery, debt.Id, debt.GroupId, debt.LenderId, debt.BorrowerId, debt.Date, debt.Description, debt.Amount, debt.Status)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) AcceptDebt(req PatchDebtRequest) (Debt, error) {
	var updatedDebt Debt

	err := r.db.QueryRow(patchDebtQuery, 1, req.DebtId, req.UserId).Scan(&updatedDebt.Id, &updatedDebt.GroupId, &updatedDebt.LenderId, &updatedDebt.BorrowerId, &updatedDebt.Date, &updatedDebt.Description, &updatedDebt.Amount, &updatedDebt.Status)
	if err != nil {
		return Debt{}, errors.NewUnauthorized("User accepting the request is not the real borrower")
	}

	return updatedDebt, nil
}

func (r *repo) RejectDebt(req PatchDebtRequest) (Debt, error) {
	var updatedDebt Debt

	err := r.db.QueryRow(patchDebtQuery, 0, req.DebtId, req.UserId).Scan(&updatedDebt.Id, &updatedDebt.GroupId, &updatedDebt.LenderId, &updatedDebt.BorrowerId, &updatedDebt.Date, &updatedDebt.Description, &updatedDebt.Amount, &updatedDebt.Status)
	if err != nil {
		return Debt{}, errors.NewUnauthorized("User rejecting the request is not the real borrower")
	}

	return updatedDebt, nil
}

func (r *repo) CancelDebt(req PatchDebtRequest) (Debt, error) {
	var updatedDebt Debt

	err := r.db.QueryRow(cancelDebtQuery, 3, req.DebtId, req.UserId).Scan(&updatedDebt.Id, &updatedDebt.GroupId, &updatedDebt.LenderId, &updatedDebt.BorrowerId, &updatedDebt.Date, &updatedDebt.Description, &updatedDebt.Amount, &updatedDebt.Status)
	if err != nil {
		return Debt{}, errors.NewUnauthorized("User accepting the request is not the real borrower")
	}

	return updatedDebt, nil
}

func (r *repo) GetSentDebts(req GetDebtsRequest) ([]Debt, error) {
	var sent []Debt

	rows, err := r.db.Query(getSentQuery, req.UserId, req.GroupId)
	if err != nil {
		return []Debt{}, errors.NewNotFound("No sent debts")
	}
	defer rows.Close()

	for rows.Next() {
		var debt Debt

		err := rows.Scan(&debt.Id, &debt.GroupId, &debt.LenderId, &debt.BorrowerId, &debt.Date, &debt.Description, &debt.Amount, &debt.Status)
		if err != nil {
			return []Debt{}, err
		}
		sent = append(sent, debt)
	}
	return sent, nil
}

func (r *repo) GetReceivedDebts(req GetDebtsRequest) ([]Debt, error) {
	var received []Debt

	rows, err := r.db.Query(getReceivedQuery, req.UserId, req.GroupId)
	if err != nil {
		return []Debt{}, errors.NewNotFound("No sent debts")
	}
	defer rows.Close()

	for rows.Next() {
		var debt Debt

		err := rows.Scan(&debt.Id, &debt.GroupId, &debt.LenderId, &debt.BorrowerId, &debt.Date, &debt.Description, &debt.Amount, &debt.Status)
		if err != nil {
			return []Debt{}, errors.NewNotFound("unexpected error scanning rows")
		}
		received = append(received, debt)
	}
	return received, nil
}
