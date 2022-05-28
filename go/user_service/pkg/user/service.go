package user

import (
	// std lib
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"

	// Internal
	"github.com/coding-kiko/user_service/pkg/errors"
	"github.com/coding-kiko/user_service/pkg/log"
)

type service struct {
	repository     Repository
	rabbitProducer RabbitProducer
	logger         log.Logger
}

const (
	defaultUserAvatar = "default_user_avatar.jpeg"
)

type Service interface {
	UpsertUser(req UpsertUserRequest) (User, error)
	UpdateUserAvatar(req FileRequest) (User, error)
	GetUser(id string) (User, error)
	GetUserGroups(id string) ([]Group, error)
}

func NewService(repository Repository, rabbitProducer RabbitProducer, logger log.Logger) Service {
	return &service{
		rabbitProducer: rabbitProducer,
		repository:     repository,
		logger:         logger,
	}
}

func (s *service) GetUser(id string) (User, error) {
	user, err := s.repository.GetUser(id)
	if err != nil {
		s.logger.Error("repository.go", "GetUser", err.Error())
		return User{}, err
	}
	return user, nil
}

func (s *service) GetUserGroups(id string) ([]Group, error) {
	groups, err := s.repository.GetUserGroups(id)
	if err != nil {
		s.logger.Error("repository.go", "GetUserGroups", err.Error())
		return []Group{}, err
	}
	return groups, nil
}

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
			s.logger.Error("repository.go", "InsertUser", err.Error())
			return User{}, err
		}
		return user, nil
	}
	// case update user
	user, err := s.repository.UpdateUser(req)
	if err != nil {
		s.logger.Error("repository.go", "UpdateUser", err.Error())
		return User{}, err
	}
	return user, nil
}

// Update user avatar by id
func (s *service) UpdateUserAvatar(req FileRequest) (User, error) {

	// check is file is valid
	file, err := ValidateFile(req.File)
	if err != nil {
		return User{}, err
	}

	// generate new avatar ready to be stored
	newAvatar := NewAvatar(file, req.Id)

	// send new avatar - via rabbitmq
	err = s.rabbitProducer.AvatarQueue(newAvatar)
	if err != nil {
		s.logger.Error("rabbitmq.go", "AvatarQueue", err.Error())
		return User{}, err
	}

	newAvatarreq := UpdateAvatarRequest{Id: req.Id, Avatar: newAvatar.Name}
	user, err := s.repository.UpdateAvatar(newAvatarreq)
	if err != nil {
		s.logger.Error("repository.go", "InsertUser", err.Error())
		return User{}, err
	}
	return user, nil
}

// Check if file is valid image and extract image type
func ValidateFile(file multipart.File) (File, error) {
	// check if any file was passed
	if file == nil {
		return File{}, errors.NewFileError("invalid file")
	}

	// check if the file is an image
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return File{}, errors.NewFileError("error reading file")
	}
	mimeType := http.DetectContentType(data)
	if mimeType != "image/png" && mimeType != "image/jpg" && mimeType != "image/jpeg" && mimeType != "application/octet-stream" {
		return File{}, errors.NewFileError("file must be an image")
	}

	return File{
		Data:     data,
		MimeType: mimeType,
	}, nil
}

func NewAvatar(file File, id string) File {
	// parse extension
	ext := strings.Split(file.MimeType, "/")[1]
	if ext == "octet-stream" {
		ext = "heic"
	}

	// hash id and generate file name: hashedId.ext
	hashedId := md5.Sum([]byte(id))
	name := fmt.Sprintf("%s.%s", fmt.Sprintf("%x", hashedId), ext)
	file.Name = name

	return file
}
