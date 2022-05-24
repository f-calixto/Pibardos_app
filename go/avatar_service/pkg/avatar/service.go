package avatar

import (
	// std lib
	"bytes"
	"image"
	"image/jpeg"

	// Internal
	"github.com/coding-kiko/avatar_service/pkg/log"

	// third party

	"golang.org/x/image/draw"
)

const (
	smallImageSize  int = 80
	mediumImageSize int = 300
	bigImageSize    int = 800
)

type service struct {
	avatarStorage AvatarStorage
	logger        log.Logger
}

type Service interface {
	ProcessAvatar(file File) error
}

func NewService(avatarStorage AvatarStorage, logger log.Logger) Service {
	return &service{
		avatarStorage: avatarStorage,
		logger:        logger,
	}
}

// handle avatar resizing and process to prepare for storage
func (s *service) ProcessAvatar(file File) error {
	s.logger.Info("service.go", "ProcessAvatar", file.Filename)

	// Resize to smaller file size and call storage method
	smallData, err := Resize(smallImageSize, file.Data)
	if err != nil {
		s.logger.Error("service.go", "ProcessAvatar", "error resizing to small image")
		return err
	}
	smallFile := File{Data: smallData, Filename: file.Filename, Size: smallImageSize}
	err = s.avatarStorage.StoreAvatar(smallFile)
	if err != nil {
		return err
	}

	// Resize to medium file size and call storage method
	mediumData, err := Resize(mediumImageSize, file.Data)
	if err != nil {
		s.logger.Error("service.go", "ProcessAvatar", "error resizing to medium image")
		return err
	}
	mediumFile := File{Data: mediumData, Filename: file.Filename, Size: mediumImageSize}
	err = s.avatarStorage.StoreAvatar(mediumFile)
	if err != nil {
		return err
	}

	// Resize to smaller file size and call storage method
	bigData, err := Resize(bigImageSize, file.Data)
	if err != nil {
		s.logger.Error("service.go", "ProcessAvatar", "error resizing to big image")
		return err
	}
	bigFile := File{Data: bigData, Filename: file.Filename, Size: bigImageSize}
	err = s.avatarStorage.StoreAvatar(bigFile)
	if err != nil {
		return err
	}

	return nil
}

// resize image sent as []byte to width (size) - maintain ratio
func Resize(size int, data []byte) ([]byte, error) {
	// convert []byte to image.Image
	img, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// Resize
	dst := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.ApproxBiLinear.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)

	// convert image.Image back to []byte
	newBuf := new(bytes.Buffer)
	jpeg.Encode(newBuf, dst, nil)
	resizedData := newBuf.Bytes()

	return resizedData, nil
}
