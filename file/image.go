package file

import (
	"bytes"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/devasiajoseph/webapp/core"
	"github.com/google/uuid"
)

const (
	MB = 1 << 20
)

type ImageData struct {
	ImageID       int       `db:"image_id" json:"image_id"`
	Filename      string    `db:"file_name" json:"file_name"`
	FileSize      int64     `db:"file_size" json:"file_size"`
	Height        int       `db:"height" json:"height"`
	Width         int       `db:"width" json:"width"`
	OriginalImage string    `db:"original_image" json:"original_image"`
	UploadedTime  time.Time `json:"UploadedTime"`
	MaxUploadSize int64
}

func SaveFile(file multipart.File, filePath string) error {
	f, err := os.OpenFile(core.AbsolutePath(filePath), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return err
	}
	io.Copy(f, file)
	f.Close()

	return nil
}

var ImageuploadPath = "static/uploads/images"

func GetFileSize(file multipart.File) (int64, error) {
	// Read the contents of the file into a buffer
	buffer := bytes.NewBuffer(nil)
	_, err := io.Copy(buffer, file)
	if err != nil {
		log.Println("Error calculating file size")
		return 0, err
	}

	// Get the size of the buffer, which is the file size in bytes
	fileSize := int64(buffer.Len())

	return fileSize, nil
}

func ValidImageType(contentType string) bool {
	return contentType == "image/png" || contentType == "image/jpeg" || contentType == "image/gif" || contentType == "image/webp"
}

func ValidUploadSize(file multipart.File, fs int64) bool {
	size, err := GetFileSize(file)
	if err != nil {
		return false
	}
	return size <= fs
}

func ExtractExtension(contentType string) (string, error) {
	switch {
	case strings.Contains(contentType, "jpeg"):
		return ".jpg", nil
	case strings.Contains(contentType, "png"):
		return ".png", nil
	case strings.Contains(contentType, "gif"):
		return ".gif", nil
	case strings.Contains(contentType, "bmp"):
		return ".bmp", nil
	case strings.Contains(contentType, "webp"):
		return ".webp", nil
	default:
		return "", errors.New("unknown content type")
	}
}

func ToWebp() {

}

func (imgd *ImageData) ProcessUpload(w http.ResponseWriter, r *http.Request, id string) error {

	file, info, err := r.FormFile(id)
	if err != nil {
		log.Println("Error getting uploaded image")
		return err
	}
	defer file.Close()
	contentType := info.Header.Get("Content-Type")
	if !ValidImageType(contentType) {
		return errors.New("unknown image type")
	}
	if !ValidUploadSize(file, imgd.FileSize*MB) {
		return errors.New("file too large")
	}

	imgd.Filename = uuid.NewString() + ".webp"
	ext, err := ExtractExtension(contentType)
	if err != nil {
		return err
	}
	imgd.OriginalImage = uuid.NewString() + ext

	/*img, _, err := image.Decode(file)

	if err != nil {
		return err
	}*/
	return err
}
