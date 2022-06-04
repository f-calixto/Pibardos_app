package calendar

import (
	// std lib
	"time"

	// internal
	"github.com/coding-kiko/calendar_service/pkg/errors"
	"github.com/coding-kiko/calendar_service/pkg/log"

	// third party
	"github.com/google/uuid"
)

type service struct {
	repository Repository
	logger     log.Logger
}

type Service interface {
	NewEvent(req NewEventRequest) (Event, error)
	JoinEvent(req JoinEventRequest) (Event, error)
	GetEvents(groupId string) ([]Event, error)
	CancelEvent(req CancelEventRequest) (Event, error)
	UpdateEvent(req UpdateEventRequest) (Event, error)
}

func NewService(repository Repository, logger log.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (s *service) UpdateEvent(req UpdateEventRequest) (Event, error) {
	// check if the creator is updating
	err := s.repository.IsCreator(req.EventId, req.UserId)
	if err != nil {
		return Event{}, err
	}

	// if any date update is present
	if req.Start_date != nil || req.End_date != nil {
		var start, end string

		// if both dates are updated
		if req.Start_date != nil && req.End_date != nil {
			start, end = *req.Start_date, *req.End_date
		}
		// if only start date is present
		if req.Start_date != nil && req.End_date == nil {
			currentEnd := s.repository.GetEndDate(req.EventId)
			start, end = *req.Start_date, currentEnd
		}
		// if only end date is present
		if req.Start_date == nil && req.End_date != nil {
			currentStart := s.repository.GetStartDate(req.EventId)
			start, end = currentStart, *req.End_date
		}

		// check if dates input by user make sense
		if err := CheckEndDateValidity(end, start); err != nil {
			return Event{}, err
		}
		// check for event colissions
		err := s.repository.CheckEventCollision(req.GroupId, req.EventId, start, end)
		if err != nil {
			return Event{}, err
		}
	}

	events, err := s.repository.UpdateEvent(req)
	if err != nil {
		return Event{}, err
	}
	return events, nil
}

func (s *service) CancelEvent(req CancelEventRequest) (Event, error) {
	// check if the creator is cancelling
	err := s.repository.IsCreator(req.EventId, req.UserId)
	if err != nil {
		return Event{}, err
	}

	events, err := s.repository.CancelEvent(req)
	if err != nil {
		return Event{}, err
	}
	return events, nil
}

func (s *service) GetEvents(groupId string) ([]Event, error) {
	events, err := s.repository.GetEvents(groupId)
	if err != nil {
		return []Event{}, err
	}
	return events, nil
}

func (s *service) JoinEvent(req JoinEventRequest) (Event, error) {
	// check if user has already joined in the activity
	err := s.repository.CheckForUserInTheEvent(req.EventId, req.UserId)
	if err != nil {
		return Event{}, err
	}

	event, err := s.repository.JoinEvent(req)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

func (s *service) NewEvent(req NewEventRequest) (Event, error) {
	var eventId = uuid.New().String()
	var endDate = req.End_date

	//if no end date is specified make it last 1 hour
	if endDate == "" {
		endDate = AddOneHour(req.Start_date)
	}

	// check if dates input by user make sense
	if err := CheckEndDateValidity(endDate, req.Start_date); err != nil {
		return Event{}, err
	}

	// check for event colissions
	err := s.repository.CheckEventCollision(req.GroupId, eventId, req.Start_date, endDate)
	if err != nil {
		return Event{}, err
	}

	// create new event
	event := Event{
		Id:         eventId,
		GroupId:    req.GroupId,
		CreatorId:  req.CreatorId,
		Title:      req.Title,
		Start_date: req.Start_date,
		End_date:   endDate,
		Cancelled:  false,
		Guest_list: []string{req.CreatorId},
	}
	err = s.repository.NewEvent(event)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

// check if a endDate 'e' is equal or before startDate 's' - receives dates in iso format
func CheckEndDateValidity(e string, s string) error {
	// convert string to time
	E, _ := time.Parse(isoFormat, e)
	S, _ := time.Parse(isoFormat, s)

	// calculate time since date - rounded to the minute
	sinceA := time.Since(E).Round(time.Minute)
	sinceB := time.Since(S).Round(time.Minute)

	if sinceA >= sinceB {
		return errors.NewInvalidDate("end date must be after start date")
	}
	return nil
}

// creates an end date that is one hour later than the start date
func AddOneHour(startDate string) string {
	t, _ := time.Parse(isoFormat, startDate)
	endDate := t.Add(time.Duration(1) * time.Hour)
	return endDate.Format(isoFormat)
}
