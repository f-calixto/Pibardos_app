package group

// this file contains request, responses and entity definitions used in the whole package

import (
	// std lib
	"mime/multipart"

	// third party
	"github.com/dgrijalva/jwt-go"
)

// response body placeholder for successful responses
type HttpSuccessResponseBody struct {
	StatusCode int64         `json:"statusCode"`
	Data       []interface{} `json:"data"`
}

// dummy struct to avoid passing string as key in the http request context
// used in middleware.go
type UserIdKey struct {
}

// used to parse jwt payload
// used in middleware.go
type Claims struct {
	userId string
	jwt.StandardClaims
}

type Group struct {
	Id                     string
	Name                   string
	Size                   int64
	Admin_id               string
	Country                string
	Access_code            string
	Access_code_issue_time int64
	Avatar_route           string
	Created_at             string
}

type CreateGroupRequest struct {
	name     string
	admin_id string
	country  string
	file     multipart.File
}

type CreateGroupResponse struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Size         int64  `json:"size"`
	Admin_id     string `json:"admin_id"`
	Country      string `json:"country"`
	Access_code  string `json:"access_code"`
	Avatar_route string `json:"avatar"`
	Created_at   string `json:"created_at"`
}

type JoinGroupRequest struct {
}

type JoinGroupResponse struct {
}
