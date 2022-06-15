package debts

import (
	// std lib

	// internal
	"github.com/coding-kiko/debts_service/pkg/log"

	// third party
	"github.com/google/uuid"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(repository Repository, logger log.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

type Service interface {
	CreateRequest(req DebtRequest) (DebtRequest, error)
	AcceptDebt(id string) (DebtRequest, error)
	RejectDebt(id string) (DebtRequest, error)
	GetSentRequests(req GetRequestsRequest) ([]DebtRequest, error)
	GetReceivedRequests(req GetRequestsRequest) ([]DebtRequest, error)
	GetGroupDebts(groupId string) ([]Borrower, error)
	CancelDebt(req CancelDebtRequest) (Borrower, error)
}

func (s *service) CancelDebt(req CancelDebtRequest) (Borrower, error) {
	debt, err := s.repository.CancelDebt(req)
	if err != nil {
		return Borrower{}, err
	}
	return debt, nil
}

func (s *service) GetGroupDebts(groupId string) ([]Borrower, error) {
	debts, err := s.repository.GetGroupDebts(groupId)
	if err != nil {
		return []Borrower{}, err
	}
	return debts, nil
}

func (s *service) GetSentRequests(req GetRequestsRequest) ([]DebtRequest, error) {
	sent, err := s.repository.GetSentRequests(req)
	if err != nil {
		return []DebtRequest{}, err
	}
	return sent, nil
}

func (s *service) GetReceivedRequests(req GetRequestsRequest) ([]DebtRequest, error) {
	received, err := s.repository.GetReceivedRequests(req)
	if err != nil {
		return []DebtRequest{}, err
	}
	return received, nil
}

func (s *service) RejectDebt(id string) (DebtRequest, error) {
	debtRequest, err := s.repository.RejectDebt(id)
	if err != nil {
		return DebtRequest{}, err
	}
	return debtRequest, nil
}

func (s *service) AcceptDebt(id string) (DebtRequest, error) {

	debtRequest, err := s.repository.AcceptDebt(id)
	if err != nil {
		return DebtRequest{}, err
	}

	balanceReq := BalanceDebtRequest{
		GroupId:    debtRequest.GroupId,
		LenderId:   debtRequest.LenderId,
		BorrowerId: debtRequest.BorrowerId,
		Amount:     debtRequest.Amount,
	}
	err = s.repository.BalanceDebt(balanceReq)
	if err != nil {
		return DebtRequest{}, err
	}
	return debtRequest, nil
}

func (s *service) CreateRequest(req DebtRequest) (DebtRequest, error) {
	req.Id = uuid.NewString()
	req.Status = 2 // status pending

	err := s.repository.CreateRequest(req)
	if err != nil {
		return DebtRequest{}, err
	}
	return req, nil
}
