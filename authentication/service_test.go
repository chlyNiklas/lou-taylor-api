package authentication

import (
	"errors"
	"testing"
	"time"

	"github.com/chlyNiklas/lou-taylor-api/model"
)

var stCfg *Config = &Config{
	Admin: &model.User{
		Name:     "Hans",
		Password: "Password",
	},
	JWTSecret:   "asdf",
	ValidPeriod: time.Minute,
}

func TestLogin_OK(t *testing.T) {

	auth := New(stCfg, nil)

	token, err := auth.Login(stCfg.Admin)

	if err != nil {
		t.Errorf("expected err to be nil but was: %v", err)
	}

	_, _, err = validateJWT(token, stCfg.JWTSecret)

	if err != nil {
		t.Errorf("expected err from validate to be nil but was: %v", err)
	}
}
func TestLogin_WrongCredentials(t *testing.T) {

	auth := New(stCfg, nil)

	_, err := auth.Login(&model.User{Name: "hans", Password: "wrong"})

	if !errors.Is(err, ErrInvalidCredentials) {
		t.Errorf("expected err to be  %v but was: %v", ErrInvalidCredentials, err)
	}

}
