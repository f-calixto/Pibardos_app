package group

import (
	// std lib
	"encoding/json"
	"fmt"
	"net/http"

	// Internal
	"github.com/coding-kiko/group_service/pkg/errors"
	"github.com/coding-kiko/group_service/pkg/log"
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
	CreateGroup(w http.ResponseWriter, r *http.Request)
	GetGroup(w http.ResponseWriter, r *http.Request)
	UpdateGroup(w http.ResponseWriter, r *http.Request)
	UpdateGroupAvatar(w http.ResponseWriter, r *http.Request)
	GenerateAccessCode(w http.ResponseWriter, r *http.Request)
	JoinGroup(w http.ResponseWriter, r *http.Request)
	// GetGroupMembers(w http.ResponseWriter, r *http.Request)

	MethodNotAllowedHandler() http.Handler
}

func (h *handlers) JoinGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// extract user id passed by context by jwtMiddleware
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	req := AccessCode{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	req.UserId = userId

	group, err := h.service.JoinGroup(req)
	if err != nil {
		h.logger.Error("handlers.go", "JoinGroup", err.Error())
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(group)
}

func (h *handlers) GenerateAccessCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// extract user id passed by context by jwtMiddleware
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))
	id := mux.Vars(r)["id"]

	req := AccessCodeRequest{
		UserId:  userId,
		GroupId: id,
	}
	code, err := h.service.GenerateAccessCode(req)
	if err != nil {
		h.logger.Error("handlers.go", "GenerateAccessCode", err.Error())
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(code)
}

func (h *handlers) UpdateGroupAvatar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	// if file is invalid => file == nil: I can validate this way, for now I won use the error
	file, _, _ := r.FormFile("avatar")
	req := FileRequest{
		Id:   id,
		File: file,
	}
	group, err := h.service.UpdateGroupAvatar(req)
	if err != nil {
		h.logger.Error("handlers.go", "UpdateGroupAvatar", err.Error())
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(group)
}

func (h *handlers) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	req := UpdateGroupRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	req.Id = &id

	user, err := h.service.UpdateGroup(req)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}

func (h *handlers) GetGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	group, err := h.service.GetGroup(id)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(group)
}

func (h *handlers) CreateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// extract user id passed by context by jwtMiddleware
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	// if file is invalid => file == nil: I can validate this way, for now I won use the error
	file, _, _ := r.FormFile("avatar")
	req := CreateGroupRequest{
		Name:     r.FormValue("name"),
		Country:  r.FormValue("country"),
		Admin_id: userId,
		File:     file,
	}

	group, err := h.service.CreateGroup(req)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(group)
}

// func (h *handlers) GetGroup(w http.ResponseWriter, r *http.Request) {

// }

// func (h *handlers) GenerateAccessCode(w http.ResponseWriter, r *http.Request) {

// }

// func (h *handlers) JoinGroup(w http.ResponseWriter, r *http.Request) {

// }

// func (h *handlers) UpdateGroup(w http.ResponseWriter, r *http.Request) {

// }

// override default gorilla method not allowed handler
func (h *handlers) MethodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		statusCode, resp := errors.CreateResponse(errors.NewMethodNotAllowed("method not allowed"))
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
	})
}
