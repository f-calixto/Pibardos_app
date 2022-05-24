package user

import (
	// std lib
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	// Internal
	"github.com/coding-kiko/user_service/pkg/log"
)

type service struct {
	repository     Repository
	rabbitProducer RabbitProducer
	logger         log.Logger
}

type Service interface {
	UpsertUser(req UpsertUserRequest) (User, error)
	UpdateUserAvatar(req FileRequest) (User, error)
}

func NewService(repository Repository, rabbitProducer RabbitProducer, logger log.Logger) Service {
	return &service{
		rabbitProducer: rabbitProducer,
		repository:     repository,
		logger:         logger,
	}
}

var (
	defaultUserAvatar = "default_user.jpg"
)

// receives requests of new users to insert or existing users fields to update - calls repo method accordingly
func (s *service) UpsertUser(req UpsertUserRequest) (User, error) {
	// case insert new user
	if req.Created_at != nil {
		user := User{
			Id:         *req.Id,
			Username:   strings.ToLower(*req.Username),
			Email:      strings.ToLower(*req.Email),
			Created_at: *req.Created_at,
			Country:    strings.ToUpper(*req.Country),
			Birthdate:  *req.Birthdate,
			Status:     "",
			Avatar:     defaultUserAvatar,
		}
		user, err := s.repository.InsertUser(user)
		if err != nil {
			return User{}, err
		}
		return user, nil
	}
	// case update user
	user, err := s.repository.UpdateUser(req)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *service) UpdateUserAvatar(req FileRequest) (User, error) {
	if req.File == nil {
		return User{}, errors.New("invalid file")
	}

	// check if the file is an image
	data, err := ioutil.ReadAll(req.File)
	if err != nil {
		return User{}, errors.New("error reading file")
	}
	mimeType := http.DetectContentType(data)
	if mimeType != "image/png" && mimeType != "image/jpg" && mimeType != "image/jpeg" && mimeType != "application/octet-stream" {
		return User{}, errors.New("file must be an image")
	}

	ext := strings.Split(mimeType, "/")[1]
	if ext == "octet-stream" {
		ext = "heic"
	}

	// hash id and store in form hashedId.ext
	hashedId := md5.Sum([]byte(req.Id))
	newAvatar := fmt.Sprintf("%s.%s", fmt.Sprintf("%x", hashedId), ext)

	err = s.rabbitProducer.AvatarQueue(data, newAvatar)
	if err != nil {
		s.logger.Error("service.go", "UpdateUserAvatar", "rabbit response error")
		return User{}, err
	}

	user := User{Id: req.Id, Avatar: newAvatar}
	user, err = s.repository.UpdateAvatar(user)
	if err != nil {
		s.logger.Error("service.go", "UpdateUserAvatar", "repository response error")
		return User{}, errors.New("error uploading file")
	}
	return user, nil
}
