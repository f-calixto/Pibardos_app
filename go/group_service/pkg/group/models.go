package group

// this file contains request, responses and entity definitions used in the whole package

import (
	// std lib
	"mime/multipart"
)

// dummy struct to avoid passing string as key in the http request context
// used in middleware.go/handlers.go
type UserIdKey struct {
}

type Group struct {
	Id                          string `json:"id"`
	Name                        string `json:"name"`
	Size                        int64  `json:"size"`
	Country                     string `json:"country"`
	Admin_id                    string `json:"admin_id"`
	Access_code                 string `json:"access_code,omitempty"`
	Access_code_expiration_time int64  `json:",omitempty"`
	Avatar                      string `json:"avatar"`
	Created_at                  string `json:"created_at"`
	Description                 string `json:"description"`
}

type CreateGroupRequest struct {
	Name     string
	Admin_id string
	Country  string
	File     multipart.File
}

type UpdateGroupRequest struct {
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Country     *string `json:"country,omitempty"`
	Description *string `json:"description,omitempty"`
	Admin_id    *string `json:"admin_id,omitempty"`
}

type File struct {
	Name     string
	Data     []byte
	MimeType string
}

type FileRequest struct {
	Id   string
	File multipart.File
}

type UpdateAvatarRequest struct {
	Id     string
	Avatar string
}

type AccessCodeRequest struct {
	GroupId string
	UserId  string
}

type AccessCode struct {
	GroupId    string `json:",omitempty"`
	UserId     string `json:",omitempty"`
	AccessCode string `json:"access_code"`
	Expiration int64  `json:",omitempty"`
}

type JoinGroupRequest struct {
}

type User struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Country    string `json:"country"`
	Birthdate  string `json:"birthdate"`
	Created_at string `json:"created_at"`
	Status     string `json:"status"`
	Avatar     string `json:"avatar"`
}

type GetGroupMembersRequest struct {
	Id     string
	Amount int
}
