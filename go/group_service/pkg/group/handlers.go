package group

import (
	// std lib
	"encoding/json"
	"net/http"

	// Internal
	"github.com/coding-kiko/group_service/pkg/log"
)

type handlers struct {
	service Service
	logger  log.Logger
}

type Handlers interface {
	// GetGroup(w http.ResponseWriter, r *http.Request)
	CreateGroup(w http.ResponseWriter, r *http.Request)
	// GenerateAccessCode(w http.ResponseWriter, r *http.Request)
	// JoinGroup(w http.ResponseWriter, r *http.Request)
	// UpdateGroup(w http.ResponseWriter, r *http.Request)
}

func NewHandler(service Service, logger log.Logger) Handlers {
	return &handlers{
		service: service,
		logger:  logger,
	}
}

func (h *handlers) CreateGroup(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(UserIdKey{})
	// me paso el error por el culo. Si el archivo es nulo o invalido: file == nil, puedo validarlo asi
	file, _, _ := r.FormFile("avatar")
	req := CreateGroupRequest{
		name:     r.FormValue("name"),
		country:  r.FormValue("country"),
		admin_id: userId.(string),
		file:     file,
	}
	resp, err := h.service.CreateGroup(req)
	if err != nil {
		return
	}
	httpSuccessResponse := HttpSuccessResponseBody{}
	httpSuccessResponse.StatusCode = 201
	httpSuccessResponse.Data = append(httpSuccessResponse.Data, resp)
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(httpSuccessResponse)
}

// func (h *handlers) GetGroup(w http.ResponseWriter, r *http.Request) {

// }

// func (h *handlers) GenerateAccessCode(w http.ResponseWriter, r *http.Request) {

// }

// func (h *handlers) JoinGroup(w http.ResponseWriter, r *http.Request) {

// }

// func (h *handlers) UpdateGroup(w http.ResponseWriter, r *http.Request) {

// }
