package calendar

const isoFormat = "2006-01-02T15:04:05.999999999Z"

// dummy struct to avoid passing string as key in the http request context
// used in middleware.go/handlers.go
type UserIdKey struct {
}

type Event struct {
	Id         string   `json:"id"`
	GroupId    string   `json:"group_id"`
	Title      string   `json:"title"`
	Start_date string   `json:"start_date"`
	End_date   string   `json:"end_date"`
	Guest_list []string `json:"guest_list"`
	Cancelled  bool     `json:"cancelled"`
	CreatorId  string   `json:"creator_id"`
}

type NewEventRequest struct {
	GroupId    string
	Title      string `json:"title"`
	Start_date string `json:"start_date"`
	End_date   string `json:"end_date"`
	CreatorId  string
}

type JoinEventRequest struct {
	UserId  string
	EventId string
}

type CancelEventRequest struct {
	UserId  string
	EventId string
}

type UpdateEventRequest struct {
	GroupId    string
	UserId     string
	EventId    string
	Title      *string `json:"title,omitempty"`
	Start_date *string `json:"start_date,omitempty"`
	End_date   *string `json:"end_date,omitempty"`
}
