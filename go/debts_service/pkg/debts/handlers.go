package debts

import (
	// std lib
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	// internal
	"github.com/coding-kiko/debts_service/pkg/errors"
	"github.com/coding-kiko/debts_service/pkg/log"

	// third party
	"github.com/gorilla/mux"
)

type handlers struct {
	service Service
	logger  log.Logger
}

func NewHandler(service Service, logger log.Logger) Handlers {
	return &handlers{
		service: service,
		logger:  logger,
	}
}

type Handlers interface {
	CreateDebt(w http.ResponseWriter, r *http.Request)
	AcceptDebt(w http.ResponseWriter, r *http.Request)
	RejectDebt(w http.ResponseWriter, r *http.Request)
	GetSentDebts(w http.ResponseWriter, r *http.Request)
	GetReceivedDebts(w http.ResponseWriter, r *http.Request)
	CancelDebt(w http.ResponseWriter, r *http.Request)

	MethodNotAllowedHandler() http.Handler
}

func (h *handlers) CreateDebt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	groupId := mux.Vars(r)["group_id"]
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	req := DebtRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(400)
		time.Sleep(1 * time.Millisecond)
		return
	}
	req.GroupId = groupId

	// check if user requesting debts is authenicated user
	if userId != req.LenderId {
		statusCode, resp := errors.CreateResponse(errors.NewUnauthorized("cannot create debt request: user id and lender id do not match"))
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	debtRequest, err := h.service.CreateDebt(req)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(debtRequest)
}

func (h *handlers) CancelDebt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	requestId := mux.Vars(r)["request_id"]
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	req := PatchDebtRequest{
		RequestId: requestId,
		UserId:    userId,
	}

	debt, err := h.service.CancelDebt(req)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(debt)
}

func (h *handlers) RejectDebt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	requestId := mux.Vars(r)["request_id"]
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	req := PatchDebtRequest{
		RequestId: requestId,
		UserId:    userId,
	}

	debt, err := h.service.RejectDebt(req)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(debt)
}

func (h *handlers) AcceptDebt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	requestId := mux.Vars(r)["request_id"]
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	req := PatchDebtRequest{
		RequestId: requestId,
		UserId:    userId,
	}

	debt, err := h.service.AcceptDebt(req)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(debt)
}

func (h *handlers) GetReceivedDebts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	groupId := mux.Vars(r)["group_id"]
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	req := GetDebtsRequest{
		GroupId: groupId,
		UserId:  userId,
	}
	received, err := h.service.GetReceivedDebts(req)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(received)
}

func (h *handlers) GetSentDebts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	groupId := mux.Vars(r)["group_id"]
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	req := GetDebtsRequest{
		GroupId: groupId,
		UserId:  userId,
	}
	received, err := h.service.GetSentDebts(req)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(received)
}

// override default gorilla method not allowed handler
func (h *handlers) MethodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		statusCode, resp := errors.CreateResponse(errors.NewMethodNotAllowed("method not allowed"))
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
	})
}
