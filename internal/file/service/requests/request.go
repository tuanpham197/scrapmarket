package requests

import (
	"mime/multipart"
)

type FileImage struct {
	File *multipart.FileHeader `form:"file" binding:"required" swaggerignore:"true"`
}

type FileImageMultiple struct {
	Files []*multipart.FileHeader `form:"files[]" validate:"image" swaggerignore:"true"`
}

type RequestDelete struct {
	Path string `json:"path" binding:"required"`
}
