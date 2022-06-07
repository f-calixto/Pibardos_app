package user

import (
	// std lib
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	// Internal
	"github.com/coding-kiko/user_service/pkg/errors"
	"github.com/coding-kiko/user_service/pkg/log"

	// third party
	"github.com/gorilla/mux"
)

type handlers struct {
	service Service
	logger  log.Logger
}

type Handlers interface {
	UpdateUser(w http.ResponseWriter, r *http.Request)
	UpdateUserAvatar(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUserGroups(w http.ResponseWriter, r *http.Request)

	MethodNotAllowedHandler() http.Handler
}

func NewHandler(service Service, logger log.Logger) Handlers {
	return &handlers{
		service: service,
		logger:  logger,
	}
}

func (h *handlers) GetUserGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// extract user id passed by context by jwtMiddleware
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	// check if id path var is the same as jwt authenticated user id
	if userId != mux.Vars(r)["id"] {
		h.logger.Error("handlers.go", "GetUserGroups", "malformed id in path param")
		statusCode, resp := errors.CreateResponse(errors.NewJwtAuthorization("jwt id and path id do not match"))
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	groups, err := h.service.GetUserGroups(userId)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(groups)
}

func (h *handlers) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// extract user id passed by context by jwtMiddleware
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	// check if id path var is the same as jwt authenticated user id
	/*if userId != mux.Vars(r)["id"] {
		h.logger.Error("handlers.go", "GetUser", "malformed id in path param")
		statusCode, resp := errors.CreateResponse(errors.NewJwtAuthorization("jwt id and path id do not match"))
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		return
	}*/

	user, err := h.service.GetUser(userId)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}

func (h *handlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// extract user id passed by context by jwtMiddleware
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	// check if id path var is the same as jwt authenticated user id
	if userId != mux.Vars(r)["id"] {
		h.logger.Error("handlers.go", "UpdateUser", "malformed id in path param")
		statusCode, resp := errors.CreateResponse(errors.NewJwtAuthorization("jwt id and path id do not match"))
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	req := UpdateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(400)
		time.Sleep(1 * time.Millisecond)
		return
	}
	upsertReq := UpsertUserRequest{
		Id:        &userId,
		Birthdate: req.Birthdate,
		Country:   req.Country,
		Status:    req.Status,
	}
	user, err := h.service.UpsertUser(upsertReq)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}

func (h *handlers) UpdateUserAvatar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// extract user id passed by context by jwtMiddleware
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	// check if id path var is the same as jwt authenticated user id
	if userId != mux.Vars(r)["id"] {
		h.logger.Error("handlers.go", "UpdateUserAvatar", "malformed id in path param")
		statusCode, resp := errors.CreateResponse(errors.NewJwtAuthorization("jwt id and path id do not match"))
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	// if file is invalid => file == nil: I can validate this way, for now I won use the error
	file, _, _ := r.FormFile("avatar")
	req := FileRequest{
		Id:   userId,
		File: file,
	}
	user, err := h.service.UpdateUserAvatar(req)
	if err != nil {
		h.logger.Error("handlers.go", "UpdateUserAvatar", err.Error())
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
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
