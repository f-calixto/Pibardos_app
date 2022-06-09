package debts

// dummy struct to avoid passing string as key in the http request context
// used in middleware.go/handlers.go
type UserIdKey struct{}

type DebtRequest struct {
	Id          string `json:"id,omitempty"`
	GroupId     string `json:"group_id,omitempty"`
	LenderId    string `json:"lender_id"`
	BorrowerId  string `json:"borrower_id"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Amount      int64  `json:"amount"`
	Status      int64  `json:"status"`
}

type BalanceDebtRequest struct {
	GroupId    string `json:"group_id,omitempty"`
	LenderId   string `json:"lender_id"`
	BorrowerId string `json:"borrower_id"`
	Amount     int64  `json:"amount"`
}

type GetRequestsRequest struct {
	GroupId string
	UserId  string
}

type Borrower struct {
	BorrowerId string `json:"borrower_id"`
	Debts      []Debt `json:"debts"`
}

type Debt struct {
	LenderId string `json:"lender_id"`
	Amount   int64  `json:"amount"`
}

type CancelDebtRequest struct {
	GroupId    string `json:",omitempty"`
	Amount     int64  `json:"amount"`
	LenderId   string `json:"lender_id"`
	BorrowerId string `json:"borrower_id"`
}
