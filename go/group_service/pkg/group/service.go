package group

import (
	// std lib
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	// Internal
	"github.com/coding-kiko/group_service/pkg/errors"
	"github.com/coding-kiko/group_service/pkg/log"

	// third party
	"github.com/google/uuid"
)

const (
	defaultGroupAvatar = "default_group_avatar.jpeg"
	isoFormat          = "2006-01-02T15:04:05.999999999Z"
)

type service struct {
	repository     Repository
	rabbitProducer RabbitProducer
	logger         log.Logger
}

type Service interface {
	CreateGroup(req CreateGroupRequest) (Group, error)
	GetGroup(id string) (Group, error)
	UpdateGroup(req UpdateGroupRequest) (Group, error)
	UpdateGroupAvatar(req FileRequest) (Group, error)
	GenerateAccessCode(req AccessCodeRequest) (AccessCode, error)
	JoinGroup(req AccessCode) (Group, error)
	GetGroupMembers(req GetGroupMembersRequest) ([]User, error)
}

func NewService(repository Repository, rabbitProducer RabbitProducer, logger log.Logger) Service {
	return &service{
		repository:     repository,
		rabbitProducer: rabbitProducer,
		logger:         logger,
	}
}

func (s *service) GetGroupMembers(req GetGroupMembersRequest) ([]User, error) {
	members, err := s.repository.GetGroupMembers(req)
	if err != nil {
		return []User{}, err
	}
	return members, nil
}

func (s *service) JoinGroup(req AccessCode) (Group, error) {
	group, err := s.repository.JoinGroup(req)
	if err != nil {
		return Group{}, err
	}
	return group, nil
}

func (s *service) GenerateAccessCode(req AccessCodeRequest) (AccessCode, error) {
	code := AccessCode{
		AccessCode: NewAccessCode(),
		GroupId:    req.GroupId,
		UserId:     req.UserId,
		Expiration: time.Now().Add(time.Duration(30) * time.Minute).Unix(),
	}

	code, err := s.repository.StoreAccessCode(code)
	if err != nil {
		return AccessCode{}, err
	}
	return code, nil
}

func (s *service) UpdateGroupAvatar(req FileRequest) (Group, error) {
	file, err := ValidateFile(req.File)
	if err != nil {
		return Group{}, err
	}

	newAvatar := NewAvatar(file, req.Id)

	// send new avatar - via rabbitmq
	err = s.rabbitProducer.AvatarQueue(newAvatar)
	if err != nil {
		s.logger.Error("rabbitmq.go", "AvatarQueue", err.Error())
		return Group{}, err
	}

	newAvatarreq := UpdateAvatarRequest{Id: req.Id, Avatar: newAvatar.Name}
	group, err := s.repository.UpdateAvatar(newAvatarreq)
	if err != nil {
		s.logger.Error("repository.go", "InsertUser", err.Error())
		return Group{}, err
	}
	return group, nil
}

func (s *service) UpdateGroup(req UpdateGroupRequest) (Group, error) {
	group, err := s.repository.UpdateGroup(req)
	if err != nil {
		return Group{}, err
	}
	return group, nil
}

func (s *service) GetGroup(id string) (Group, error) {
	group, err := s.repository.GetGroup(id)
	if err != nil {
		return Group{}, nil
	}
	return group, nil
}

func (s *service) CreateGroup(req CreateGroupRequest) (Group, error) {
	var newAvatar string
	var id string = uuid.New().String()

	if req.File == nil {
		newAvatar = defaultGroupAvatar
	} else {
		file, err := ValidateFile(req.File)
		if err != nil {
			return Group{}, err
		}
		file = NewAvatar(file, id)

		// send new avatar - via rabbitmq
		err = s.rabbitProducer.AvatarQueue(file)
		if err != nil {
			s.logger.Error("rabbitmq.go", "AvatarQueue", err.Error())
			return Group{}, err
		}

		newAvatar = file.Name
	}

	group := Group{
		Id:                          id,
		Name:                        req.Name,
		Size:                        1,
		Country:                     strings.ToUpper(req.Country),
		Admin_id:                    req.Admin_id,
		Access_code:                 NewAccessCode(),
		Access_code_expiration_time: time.Now().Add(time.Duration(30) * time.Minute).Unix(),
		Avatar:                      newAvatar,
		Created_at:                  time.Now().UTC().Format(isoFormat),
		Description:                 "",
	}

	resp, err := s.repository.CreateGroup(group)
	if err != nil {
		s.logger.Error("repository.go", "CreateGroup", err.Error())
		return Group{}, err
	}
	return resp, nil
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
	if mimeType != "image/jpg" && mimeType != "image/jpeg" {
		return File{}, errors.NewFileError("file must be jpg or jpeg")
	}

	return File{
		Data:     data,
		MimeType: mimeType,
	}, nil
}

func NewAvatar(file File, id string) File {
	// parse extension
	ext := strings.Split(file.MimeType, "/")[1]

	// hash id and generate file name: hashedId.ext
	hashedId := md5.Sum([]byte(id))
	name := fmt.Sprintf("%s.%s", fmt.Sprintf("%x", hashedId), ext)
	file.Name = name

	return file
}

// Generate an access code that lasts 30 minutes
func NewAccessCode() string {
	bytes := make([]byte, 3)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return strings.ToUpper(hex.EncodeToString(bytes))
}
