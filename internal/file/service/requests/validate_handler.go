package requests

import (
	"mime/multipart"
	"net/http"

	"github.com/gabriel-vasile/mimetype"
)

type ValidateImageError struct {
	Code    int
	Message string
}

func (e *ValidateImageError) Error() string {
	return e.Message
}

func ValidateImageUpload(image *multipart.FileHeader) error {
	err := &ValidateImageError{
		Code:    http.StatusUnprocessableEntity,
		Message: "",
	}
	if image == nil {
		err.Message = "image file is required"
		return err
	}

	// Check that the image file is not too large.
	// fmt.Println("IMAGE SIZE: ====>>>", image.Size/1024/1024)
	if (image.Size / 1024 / 1024) > 5 {
		err.Message = "large file"
		return err
	}

	file, errOpen := image.Open()
	if errOpen != nil {
		err.Message = "open image fail"
		return err
	}

	// Check that the image file is an image.
	mime, errImage := mimetype.DetectReader(file)
	if errImage != nil {
		err.Message = "image invalid"
		return err
	}

	if mime.Is("image/png") || mime.Is("image/jpg") || mime.Is("image/jpeg") {
		return nil
	}

	// The image is valid!
	err.Message = "image invalid"
	return err

}

func ValidateStructUploadMultilple(files []*multipart.FileHeader) error {
	for _, file := range files {
		err := ValidateImageUpload(file)
		if err != nil {
			return err
		}
	}
	return nil
}
