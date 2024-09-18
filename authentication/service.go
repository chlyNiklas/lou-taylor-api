package authentication

import (
	"errors"

	"github.com/chlyNiklas/lou-taylor-api/model"
)

var ErrInvalidCredentials = errors.New("either username or password is wrong")

// PostAuthLogin checks if the credentials match for the admin user. Then generates a jwt response.
func (s *Service) Login(user *model.User) (token string, err error) {

	if user.Password != s.cfg.Admin.Password || user.Name != s.cfg.Admin.Name {
		return "", ErrInvalidCredentials
	}

	token, err = createJWT(user.Name, []string{"admin"}, s.cfg.JWTSecret, s.cfg.ValidPeriod)

	return
}
