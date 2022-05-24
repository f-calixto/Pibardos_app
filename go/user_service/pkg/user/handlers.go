package user

import (
	// std lib
	"encoding/json"
	"fmt"
	"net/http"

	// Internal
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
}

func NewHandler(service Service, logger log.Logger) Handlers {
	return &handlers{
		service: service,
		logger:  logger,
	}
}

func (h *handlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// extract user id passed by context by jwtMiddleware
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	req := UpdateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
		w.WriteHeader(http.StatusNotFound)
		return
	}
	SuccessfulResponse(w, user)
}

func (h *handlers) UpdateUserAvatar(w http.ResponseWriter, r *http.Request) {
	// extract user id passed by context by jwtMiddleware
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	if userId != mux.Vars(r)["id"] {
		h.logger.Error("handlers.go", "UpdateUserAvatar", userId+" != "+mux.Vars(r)["id"])
		w.WriteHeader(http.StatusUnauthorized)
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	SuccessfulResponse(w, user)
}

func SuccessfulResponse(w http.ResponseWriter, user User) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
