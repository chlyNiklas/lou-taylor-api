package controller

import (
	"context"
	"path"

	"github.com/chlyNiklas/lou-taylor-api/api"
	"github.com/chlyNiklas/lou-taylor-api/utils"
)

// Write an image
// (GET /images)
func (s *Service) PostImages(ctx context.Context, request api.PostImagesRequestObject) (api.PostImagesResponseObject, error) {
	mp, err := request.Body.ReadForm(1000000)

	header := mp.File["image"][0]
	file, err := header.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	image, err := utils.DecodeImage(file)
	if err != nil {
		return nil, err
	}

	name, err := s.img.SaveImage(image)
	if err != nil {
		return nil, err
	}
	url := path.Join(s.cfg.BaseUrl, "images", name)

	return api.PostImages201JSONResponse{ImageUrl: &url}, nil
}

// Delete an image
// (DELETE /images/{imageId})
func (s *Service) DeleteImagesImageName(ctx context.Context, request api.DeleteImagesImageNameRequestObject) (api.DeleteImagesImageNameResponseObject, error) {
	err := s.img.Delete(request.ImageName)
	if err != nil {
		return api.DeleteImagesImageName404Response{}, nil
	}
	return api.DeleteImagesImageName204Response{}, nil
}

// Retrieve an image
// (GET /images/{imageId})
func (s *Service) GetImagesImageName(ctx context.Context, request api.GetImagesImageNameRequestObject) (api.GetImagesImageNameResponseObject, error) {

	img, size, err := s.img.Read(request.ImageName)
	response := api.GetImagesImageName200ImagewebpResponse{Body: img, ContentLength: size}

	return response, err
}
