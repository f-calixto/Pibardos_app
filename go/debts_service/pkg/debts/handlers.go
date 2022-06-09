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
	CreateRequest(w http.ResponseWriter, r *http.Request)
	AcceptRequest(w http.ResponseWriter, r *http.Request)
	RejectRequest(w http.ResponseWriter, r *http.Request)
	GetReceivedRequests(w http.ResponseWriter, r *http.Request)
	GetSentRequests(w http.ResponseWriter, r *http.Request)
	GetGroupDebts(w http.ResponseWriter, r *http.Request)
	CancelDebt(w http.ResponseWriter, r *http.Request)

	MethodNotAllowedHandler() http.Handler
}

func (h *handlers) CancelDebt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	groupId := mux.Vars(r)["group_id"]
	Id := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	req := CancelDebtRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(400)
		time.Sleep(1 * time.Millisecond)
		return
	}
	req.GroupId = groupId

	if req.BorrowerId != Id {
		statusCode, resp := errors.CreateResponse(errors.NewJwtAuthorization("authenticated user is not the same as borrower"))
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		return
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

func (h *handlers) GetGroupDebts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	groupId := mux.Vars(r)["group_id"]

	groupDebts, err := h.service.GetGroupDebts(groupId)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(groupDebts)
}

func (h *handlers) GetReceivedRequests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	groupId := mux.Vars(r)["group_id"]
	userId := mux.Vars(r)["user_id"]
	Id := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	// check if id path var is the same as jwt authenticated user id
	if userId != Id {
		h.logger.Error("handlers.go", "UpdateUserAvatar", "malformed id in path param")
		statusCode, resp := errors.CreateResponse(errors.NewJwtAuthorization("jwt id and path id do not match"))
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		return
	}

	req := GetRequestsRequest{
		GroupId: groupId,
		UserId:  userId,
	}
	received, err := h.service.GetReceivedRequests(req)
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

func (h *handlers) GetSentRequests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	groupId := mux.Vars(r)["group_id"]
	userId := mux.Vars(r)["user_id"]
	Id := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	// check if id path var is the same as jwt authenticated user id
	if userId != Id {
		h.logger.Error("handlers.go", "UpdateUserAvatar", "malformed id in path param")
		statusCode, resp := errors.CreateResponse(errors.NewJwtAuthorization("jwt id and path id do not match"))
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		return
	}

	req := GetRequestsRequest{
		GroupId: groupId,
		UserId:  userId,
	}

	sent, err := h.service.GetSentRequests(req)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(sent)
}

func (h *handlers) RejectRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	requestId := mux.Vars(r)["request_id"]

	debtRequest, err := h.service.RejectDebt(requestId)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(debtRequest)
}

func (h *handlers) AcceptRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	requestId := mux.Vars(r)["request_id"]

	debtRequest, err := h.service.AcceptDebt(requestId)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(debtRequest)
}

func (h *handlers) CreateRequest(w http.ResponseWriter, r *http.Request) {
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

	debtRequest, err := h.service.CreateRequest(req)
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
