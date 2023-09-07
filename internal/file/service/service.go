package service

import (
	"context"
	"fmt"
	"sendo/internal/file/service/requests"
	"sendo/pkg/common"
	"strings"
)

type service struct {
	repository FileRepository
}

func NewService(repository FileRepository) service {
	return service{repository: repository}
}

// Upload image
// @Summary      Upload image
// @Description  Upload image
// @Param 		 request body requests.FileImage true "upload param"
// @Tags         File
// @Produce      json
// @Success		 200	{object} string
// @Failure		 400	{object} error
// @Router       /files/image [post]
func (s service) UploadImage(ctx context.Context, fileImage requests.FileImage) (*string, error) {
	err := requests.ValidateImageUpload(fileImage.File)
	if err != nil {
		return nil, err
	}
	file, err := common.UploadImage(fileImage.File, "category")
	if err != nil {
		return nil, err
	}
	return &file, nil
}

// Remove image
// @Summary      Remove image
// @Description  Remove image
// @Param 		 request body string true "upload param"
// @Tags         File
// @Produce      json
// @Success		 200	{object} bool
// @Failure		 400	{object} error
// @Router       /files/image [delete]
func (s service) RemoveImage(ctx context.Context, path string) (bool, error) {
	params := strings.Split(path, "/")
	var bucket, objectName string = "", ""
	if len(params) > 1 {
		bucket = params[0]
		objectName = params[1]
	}

	if objectName == "" {
		return false, nil
	}
	result := common.RemoveImage(objectName, bucket)
	return result, nil
}

// Upload multiple image
// @Summary      Upload multiple image
// @Description  Upload multiple image
// @Param 		 request body requests.FileImageMultiple true "upload param"
// @Tags         File
// @Produce      json
// @Success		 200	{object} bool
// @Failure		 400	{object} error
// @Failure		 422	{object} error
// @Router       /files/upload-multiple [POST]
func (s service) UploadMultipleImage(ctx context.Context, files requests.FileImageMultiple) (*[]string, error) {
	err := requests.ValidateStructUploadMultilple(files.Files)
	if err != nil {
		return nil, err
	}

	result, err := common.UploadMultipleImage(files.Files, "images")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(result)

	return result, nil
}
