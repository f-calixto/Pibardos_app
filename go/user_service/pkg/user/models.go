package user

import (
	// std lib
	"mime/multipart"

	// third party
	"github.com/golang-jwt/jwt"
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

type User struct {
	Id         string `json:"id,omitempty"`
	Username   string `json:"username,omitempty"`
	Email      string `json:"email,omitempty"`
	Country    string `json:"country,omitempty"`
	Birthdate  string `json:"birthdate,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Status     string `json:"status,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
}

type FileRequest struct {
	Id   string
	File multipart.File
}

type UserIdKey struct {
}

// used to parse jwt payload
// used in middleware.go
type Claims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
