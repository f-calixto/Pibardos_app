package avatar

import (
	// std lib
	"fmt"
	"io/ioutil"

	// Internal
	"github.com/coding-kiko/avatar_service/pkg/log"
)

type avatarStorage struct {
	directory string
	logger    log.Logger
}

type AvatarStorage interface {
	StoreAvatar(file File) error
}

func NewAvatarStorage(directory string, logger log.Logger) AvatarStorage {
	return &avatarStorage{
		directory: directory,
		logger:    logger,
	}
}

func (a *avatarStorage) StoreAvatar(file File) error {
	a.logger.Info("storage.go", "ProcessAvatar", fmt.Sprintf("%d/%s", file.Size, file.Filename))

	fileAbsolutePath := fmt.Sprintf("%s%d/%s", a.directory, file.Size, file.Filename)
	err := ioutil.WriteFile(fileAbsolutePath, file.Data, 0644)
	if err != nil {
		a.logger.Error("storage.go", "ProcessAvatar", err.Error())
		return err
	}
	return nil
}
