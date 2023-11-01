package service

import (
	"context"
	"sendo/internal/file/service/requests"
)

type FileUseCase interface {
	UploadImage(ctx context.Context, fileImage requests.FileImage) (*string, error)
	UploadMultipleImage(ctx context.Context, files requests.FileImageMultiple) (*[]string, error)
	RemoveImage(ctx context.Context, imageName string) (bool, error)
}

type FileRepository interface {
}
