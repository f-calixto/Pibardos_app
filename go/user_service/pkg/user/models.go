package user

import (
	// std lib
	"mime/multipart"
)

// Update or insert
type UpsertUserRequest struct {
	Id         *string `json:"id,omitempty"`
	Username   *string `json:"username,omitempty"`
	Email      *string `json:"email,omitempty"`
	Country    *string `json:"country,omitempty"`
	Birthdate  *string `json:"birthdate,omitempty"`
	Created_at *string `json:"created_at,omitempty"`
	Status     *string `json:"status,omitempty"`
}

type UpdateUserRequest struct {
	Country   *string `json:"country,omitempty"`
	Birthdate *string `json:"birthdate,omitempty"`
	Status    *string `json:"status,omitempty"`
}

type UpdateAvatarRequest struct {
	Id     string
	Avatar string
}

type GetUserRequest struct {
	Id string
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

type Group struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	Admin_id   string `json:"admin_id"`
	Country    string `json:"country"`
	Avatar     string `json:"avatar"`
	Created_at string `json:"created_at"`
}

type FileRequest struct {
	Id   string
	File multipart.File
}

type File struct {
	Name     string
	Data     []byte
	MimeType string
}

type UserIdKey struct {
}
