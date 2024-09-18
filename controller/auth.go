package controller

import (
	"context"

	"github.com/chlyNiklas/lou-taylor-api/api"
	"github.com/chlyNiklas/lou-taylor-api/authentication"
	"github.com/chlyNiklas/lou-taylor-api/model"
)

// PostAuthLogin checks if the credentials match for the admin user. Then generates a jwt response.
func (s *Service) PostAuthLogin(ctx context.Context, request api.PostAuthLoginRequestObject) (api.PostAuthLoginResponseObject, error) {

	user := &model.User{}

	user.Name = request.Body.Username
	user.Password = request.Body.Password

	token, err := s.auth.Login(user)

	if err == authentication.ErrInvalidCredentials {
		return api.PostAuthLogin401Response{}, nil
	}

	// Implementation here
	return api.PostAuthLogin200JSONResponse{Token: &token}, nil
}
