package server_service

import (
	"context"

	"github.com/chlyNiklas/lou-taylor-api/api"
	"github.com/chlyNiklas/lou-taylor-api/utils"
)

// PostAuthLogin checks if the credentials match for the admin user. Then generates a jwt response.
func (s *Service) PostAuthLogin(ctx context.Context, request api.PostAuthLoginRequestObject) (api.PostAuthLoginResponseObject, error) {

	if request.Body.Password != s.cfg.Admin.Password || request.Body.Username != s.cfg.Admin.Name {
		return api.PostAuthLogin401Response{}, nil
	}

	token, err := utils.CreateJWT(request.Body.Username, []string{"admin"}, s.cfg.JWTSecret)
	// Implementation here
	return api.PostAuthLogin200JSONResponse{
		Token: &token,
	}, err
}
