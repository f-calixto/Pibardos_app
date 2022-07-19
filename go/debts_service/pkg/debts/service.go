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
	CreateDebt(req DebtRequest) (Debt, error)
	RejectDebt(req PatchDebtRequest) (Debt, error)
	AcceptDebt(req PatchDebtRequest) (Debt, error)
	CancelDebt(req PatchDebtRequest) (Debt, error)
	GetSentDebts(req GetDebtsRequest) ([]Debt, error)
	GetReceivedDebts(req GetDebtsRequest) ([]Debt, error)
}

func (s *service) CreateDebt(req DebtRequest) (Debt, error) {
	debt := Debt{
		Id:          uuid.NewString(),
		GroupId:     req.GroupId,
		LenderId:    req.LenderId,
		BorrowerId:  req.BorrowerId,
		Status:      2,
		Description: req.Description,
		Date:        req.Date,
		Amount:      req.Amount,
	}

	err := s.repository.CreateDebt(debt)
	if err != nil {
		return Debt{}, err
	}
	return debt, nil
}

func (s *service) RejectDebt(req PatchDebtRequest) (Debt, error) {
	debt, err := s.repository.RejectDebt(req)
	if err != nil {
		return Debt{}, err
	}
	return debt, nil
}

func (s *service) AcceptDebt(req PatchDebtRequest) (Debt, error) {
	debt, err := s.repository.AcceptDebt(req)
	if err != nil {
		return Debt{}, err
	}
	return debt, nil
}

func (s *service) CancelDebt(req PatchDebtRequest) (Debt, error) {
	debt, err := s.repository.CancelDebt(req)
	if err != nil {
		return Debt{}, err
	}
	return debt, nil
}

func (s *service) GetSentDebts(req GetDebtsRequest) ([]Debt, error) {
	sent, err := s.repository.GetSentDebts(req)
	if err != nil {
		return []Debt{}, err
	}
	return sent, nil
}

func (s *service) GetReceivedDebts(req GetDebtsRequest) ([]Debt, error) {
	received, err := s.repository.GetReceivedDebts(req)
	if err != nil {
		return []Debt{}, err
	}
	return received, nil
}
