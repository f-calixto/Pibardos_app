package calendar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/coding-kiko/calendar_service/pkg/errors"
	"github.com/coding-kiko/calendar_service/pkg/log"
	"github.com/gorilla/mux"
)

type handlers struct {
	service Service
	logger  log.Logger
}

type Handlers interface {
	NewEvent(w http.ResponseWriter, r *http.Request)
	JoinEvent(w http.ResponseWriter, r *http.Request)
	GetEvents(w http.ResponseWriter, r *http.Request)
	CancelEvent(w http.ResponseWriter, r *http.Request)
	UpdateEvent(w http.ResponseWriter, r *http.Request)

	MethodNotAllowedHandler() http.Handler
}

func NewHandler(service Service, logger log.Logger) Handlers {
	return &handlers{
		service: service,
		logger:  logger,
	}
}

func (h *handlers) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// extract user id passed by context by jwtMiddleware
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	eventId := mux.Vars(r)["event_id"]
	groupId := mux.Vars(r)["group_id"]

	req := UpdateEventRequest{}
	json.NewDecoder(r.Body).Decode(&req)
	req.GroupId = groupId
	req.UserId = userId
	req.EventId = eventId

	event, err := h.service.UpdateEvent(req)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(event)
}

func (h *handlers) CancelEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// extract user id passed by context by jwtMiddleware
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	eventId := mux.Vars(r)["event_id"]

	req := CancelEventRequest{
		UserId:  userId,
		EventId: eventId,
	}
	event, err := h.service.CancelEvent(req)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(event)
}

func (h *handlers) GetEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	groupId := mux.Vars(r)["group_id"]

	events, err := h.service.GetEvents(groupId)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(events)
}

func (h *handlers) JoinEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// extract user id passed by context by jwtMiddleware
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))

	req := JoinEventRequest{
		UserId:  userId,
		EventId: mux.Vars(r)["event_id"],
	}

	event, err := h.service.JoinEvent(req)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(event)
}

func (h *handlers) NewEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// extract user id passed by context by jwtMiddleware
	userId := fmt.Sprintf("%v", r.Context().Value(UserIdKey{}))
	groupId := mux.Vars(r)["group_id"]

	req := NewEventRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(400)
		time.Sleep(1 * time.Millisecond)
		return
	}
	req.CreatorId = userId
	req.GroupId = groupId

	event, err := h.service.NewEvent(req)
	if err != nil {
		statusCode, resp := errors.CreateResponse(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
		time.Sleep(1 * time.Millisecond)
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(event)
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
