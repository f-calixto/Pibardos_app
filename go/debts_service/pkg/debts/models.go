package debts

// dummy struct to avoid passing string as key in the http request context
// used in middleware.go/handlers.go
type UserIdKey struct{}

type Debt struct {
	Id          string `json:"id,omitempty"`
	GroupId     string `json:"group_id,omitempty"`
	LenderId    string `json:"lender_id"`
	BorrowerId  string `json:"borrower_id"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Amount      int64  `json:"amount"`
	Status      int64  `json:"status"`
}

type DebtRequest struct {
	GroupId     string `json:"group_id,omitempty"`
	LenderId    string `json:"lender_id"`
	BorrowerId  string `json:"borrower_id"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Amount      int64  `json:"amount"`
}

type GetDebtsRequest struct {
	GroupId string
	UserId  string
}

// used for accepting, rejecting or canceling a debt
type PatchDebtRequest struct {
	DebtId string
	UserId string
}
