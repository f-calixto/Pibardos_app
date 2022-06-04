package calendar

import (
	// std lib
	"database/sql"
	"fmt"
	"strings"
	"time"

	// internal
	"github.com/coding-kiko/calendar_service/pkg/errors"
	"github.com/coding-kiko/calendar_service/pkg/log"

	// Third Party
	"github.com/lib/pq"
)

var (
	newEventQuery = `INSERT INTO calendar_events(id, group_id, title, start_date, end_date, creator_id, cancelled, guest_list)
					Values($1, $2, $3, $4, $5, $6, $7, $8)`
	checkDuplicateGuestQuery = `SELECT guest_list FROM calendar_events WHERE id = $1`
	joinEventQuery           = `UPDATE calendar_events SET guest_list = array_append(guest_list, $1) WHERE id = $2`
	getEventQuery            = `SELECT id, group_id, title, start_date, end_date, creator_id, cancelled, guest_list FROM calendar_events WHERE id = $1`
	getAllEventsQuery        = `SELECT id, group_id, title, start_date, end_date, creator_id, cancelled, guest_list FROM calendar_events WHERE group_id = $1`
	cancelEventQuery         = `UPDATE calendar_events SET cancelled = true WHERE id = $1`
	getCreatorIdQuery        = `SELECT creator_id FROM calendar_events WHERE id = $1`
	getStartDateQuery        = `SELECT start_date FROM calendar_events WHERE id = $1`
	getEndDateQuery          = `SELECT end_date FROM calendar_events WHERE id = $1`
	baseUpdateQuery          = `UPDATE calendar_events SET `
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

type Repository interface {
	// primary - endpoint functions
	NewEvent(event Event) error
	JoinEvent(req JoinEventRequest) (Event, error)
	GetEvents(groupId string) ([]Event, error)
	CancelEvent(req CancelEventRequest) (Event, error)
	UpdateEvent(req UpdateEventRequest) (Event, error)

	// secondary - helper functions that still need to query
	IsCreator(eventId string, userId string) error
	GetEvent(eventId string) (Event, error)
	CheckForUserInTheEvent(eventId string, userId string) error
	CheckEventCollision(groupId string, eventId string, startDate string, endDate string) error
	GetStartDate(eventId string) string
	GetEndDate(eventId string) string
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: logger,
	}
}

func (r *repo) UpdateEvent(req UpdateEventRequest) (Event, error) {

	// get custom query for the updating parameters
	updateQuery, args, err := PatchQueryConstructor(req)
	if err != nil {
		return Event{}, err
	}

	// update
	_, err = r.db.Exec(updateQuery, args...)
	if err != nil {
		return Event{}, errors.NewNotFound("event not found")
	}

	// get updated event
	updatedEvent, err := r.GetEvent(req.EventId)
	if err != nil {
		return Event{}, err
	}

	return updatedEvent, nil
}

func (r *repo) CancelEvent(req CancelEventRequest) (Event, error) {

	// cancel
	_, err := r.db.Exec(cancelEventQuery, req.EventId)
	if err != nil {
		return Event{}, errors.NewNotFound("event not found")
	}

	// get updated event
	updatedEvent, err := r.GetEvent(req.EventId)
	if err != nil {
		return Event{}, err
	}

	return updatedEvent, nil
}

func (r *repo) GetEvents(groupId string) ([]Event, error) {
	var events []Event

	rows, err := r.db.Query(getAllEventsQuery, groupId)
	if err != nil {
		return []Event{}, errors.NewNotFound("no group found")
	}
	defer rows.Close()

	for rows.Next() {
		event := Event{}
		err := rows.Scan(&event.Id, &event.GroupId, &event.Title, &event.Start_date, &event.End_date, &event.CreatorId, &event.Cancelled, pq.Array(&event.Guest_list))
		if err != nil {
			return []Event{}, errors.NewNotFound()
		}
		events = append(events, event)
	}
	r.logger.Info("repository.go", "GetEvents", "group events retrieved successfuly")
	return events, nil
}

func (r *repo) JoinEvent(req JoinEventRequest) (Event, error) {

	// join event
	_, err := r.db.Exec(joinEventQuery, req.UserId, req.EventId)
	if err != nil {
		return Event{}, errors.NewNotFound("event not found")
	}

	// get updated event
	updatedEvent, err := r.GetEvent(req.EventId)
	if err != nil {
		return Event{}, err
	}

	return updatedEvent, nil
}

// create new event for group if there are no conflicts
func (r *repo) NewEvent(event Event) error {

	// create new event
	_, err := r.db.Exec(newEventQuery, event.Id, event.GroupId, event.Title, event.Start_date, event.End_date, event.CreatorId, event.Cancelled, pq.Array(event.Guest_list))
	if err != nil {
		return errors.NewNotFound("event not found")
	}
	return nil
}

// chcek if user is already in the event
func (r *repo) CheckForUserInTheEvent(eventId string, userId string) error {
	var guests []string

	err := r.db.QueryRow(checkDuplicateGuestQuery, eventId).Scan(pq.Array(&guests))
	if err != nil {
		return errors.NewNotFound("event not found")
	}
	for _, guest := range guests {
		if guest == userId {
			return errors.NewAlreadyJoined("you have already joined the activity")
		}
	}
	return nil
}

// check if userId requesting the action is the creator of the event
func (r *repo) IsCreator(eventId string, userId string) error {
	var creatorId string

	err := r.db.QueryRow(getCreatorIdQuery, eventId).Scan(&creatorId)
	if err != nil {
		return errors.NewNotFound("event not found")
	}
	if creatorId != userId {
		return errors.NewUnauthorized("you are not the creator of the event")
	}
	return nil
}

// get single event based on its id
func (r *repo) GetEvent(eventId string) (Event, error) {
	event := Event{}
	err := r.db.QueryRow(getEventQuery, eventId).Scan(&event.Id, &event.GroupId, &event.Title, &event.Start_date, &event.End_date, &event.CreatorId, &event.Cancelled, pq.Array(&event.Guest_list))
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

// get startDate by event Id
func (r *repo) GetStartDate(eventId string) string {
	var startDate string

	_ = r.db.QueryRow(getStartDateQuery, eventId).Scan(&startDate)
	return startDate
}

// get endDate by event Id
func (r *repo) GetEndDate(eventId string) string {
	var endDate string

	_ = r.db.QueryRow(getEndDateQuery, eventId).Scan(&endDate)
	return endDate
}

// check for any event collision for a group calendar
func (r *repo) CheckEventCollision(groupId string, eventId string, startDate string, endDate string) error {
	events, err := r.GetEvents(groupId)
	if err != nil {
		return err
	}
	for _, e := range events {
		if e.Cancelled {
			continue
		}
		if err := CheckDateCollision(startDate, endDate, e.Start_date, e.End_date); err != nil {
			if e.Id == eventId { // collision updating the same date
				continue
			}
			return err
		}
	}
	return nil
}

// will stop working in 2038 or for events after that :/
// check if two date range collide in time
func CheckDateCollision(s1 string, e1 string, s2 string, e2 string) error {
	a, _ := time.Parse(isoFormat, s1)
	b, _ := time.Parse(isoFormat, e1)
	c, _ := time.Parse(isoFormat, s2)
	d, _ := time.Parse(isoFormat, e2)

	if (a.Unix() < d.Unix()) && (b.Unix() > c.Unix()) {
		return errors.NewOverlapingDate("there is already an event on that same time")
	}
	return nil
}

// creates event update query dynamically depending on the fields to be updated
func PatchQueryConstructor(req UpdateEventRequest) (string, []interface{}, error) {
	i := 1 // increments accordingly eith the number of args
	queryParts := make([]string, 0)
	args := make([]interface{}, 0)
	query := baseUpdateQuery

	if req.Start_date != nil {
		queryParts = append(queryParts, fmt.Sprintf("start_date = $%d", i))
		args = append(args, *req.Start_date)
		i++
	}

	if req.End_date != nil {
		queryParts = append(queryParts, fmt.Sprintf("end_date = $%d", i))
		args = append(args, *req.End_date)
		i++
	}

	if req.Title != nil {
		queryParts = append(queryParts, fmt.Sprintf("title = $%d", i))
		args = append(args, *req.Title)
		i++
	}

	if len(queryParts) == 0 {
		return "", nil, errors.NewInvalidUpdate()
	}
	query += strings.Join(queryParts, ", ") + fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, req.EventId)

	return query, args, nil
}
