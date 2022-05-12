package group

import (
	// std lib
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	// Internal
	"github.com/coding-kiko/group_service/pkg/log"

	// third party
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository Repository
	logger     log.Logger
}

type Service interface {
	// JoinGroup(req JoinGroupRequest) (JoinGroupResponse, error)
	CreateGroup(req CreateGroupRequest) (CreateGroupResponse, error)
}

func NewService(repository Repository, logger log.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (s *service) CreateGroup(req CreateGroupRequest) (CreateGroupResponse, error) {
	var id string = uuid.New().String()

	avatar_route, err := ProcessAndStoreAvatar(req.file, id)
	if err != nil {
		// manage error
		return CreateGroupResponse{}, err
	}
	group := Group{
		Id:                     id,
		Name:                   req.name,
		Size:                   1,
		Country:                req.country,
		Admin_id:               req.admin_id,
		Access_code:            NewAccessCode(),
		Access_code_issue_time: time.Now().Add(time.Duration(30) * time.Minute).Unix(),
		Avatar_route:           avatar_route,
		Created_at:             time.Now().Format("02/01/2006"),
	}
	resp, err := s.repository.CreateGroup(group)
	if err != nil {
		return CreateGroupResponse{}, err
	}
	return resp, nil
}

// func (s *service) JoinGroup(req JoinGroupRequest) (JoinGroupResponse, error) {

// }

// store avatar in a persistent volume mapped directory
// returns the relative image path
func ProcessAndStoreAvatar(file multipart.File, id string) (string, error) {
	var basePath string = "/static/"
	var defaultAvatarPath = "/static/default_group_avatar.jpg"

	if file == nil {
		return defaultAvatarPath, nil
	}
	// check if the file is an image
	data, err := ioutil.ReadAll(file)
	if err != nil {
		// manage error
		return "", err
	}
	mimeType := http.DetectContentType(data)
	if mimeType != "image/png" && mimeType != "image/jpg" && mimeType != "image/jpeg" && mimeType != "application/octet-stream" {
		return "", err
	}

	ext := strings.Split(mimeType, "/")[1]
	if ext == "octet-stream" {
		ext = "heic"
	}
	// hash id and store in form hashedId.ext
	hashedId, err := bcrypt.GenerateFromPassword([]byte(id), bcrypt.DefaultCost)
	if err != nil {
		// manage error
		return "", err
	}
	filePath := fmt.Sprintf("%s%s.%s", basePath, string(hashedId), ext)
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		// manage error
		return "", err
	}
	return filePath, nil
}

// Generate an access code that lasts 30 minutes
func NewAccessCode() string {
	bytes := make([]byte, 3)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return strings.ToUpper(hex.EncodeToString(bytes))
}
