package authentication

import (
	"errors"
	"testing"
	"time"

	"github.com/chlyNiklas/lou-taylor-api/model"
)

func TestLogin_OK(t *testing.T) {
	cfg := &Config{
		Admin: &model.User{
			Name:     "Hans",
			Password: "Password",
		},
		JWTSecret:   []byte("asdf"),
		ValidPeriod: time.Minute,
	}
	auth := New(cfg, nil)

	token, err := auth.Login(cfg.Admin)

	if err != nil {
		t.Errorf("expected err to be nil but was: %v", err)
	}

	_, _, err = validateJWT(token, cfg.JWTSecret)

	if err != nil {
		t.Errorf("expected err from validate to be nil but was: %v", err)
	}
}
func TestLogin_WrongCredentials(t *testing.T) {
	cfg := &Config{
		Admin: &model.User{
			Name:     "Hans",
			Password: "Password",
		},
		JWTSecret:   []byte("asdf"),
		ValidPeriod: time.Minute,
	}
	auth := New(cfg, nil)

	_, err := auth.Login(&model.User{Name: "hans", Password: "wrong"})

	if !errors.Is(err, ErrInvalidCredentials) {
		t.Errorf("expected err to be  %v but was: %v", ErrInvalidCredentials, err)
	}

}
