package authentication

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func validateAndExpectError(t *testing.T, expected error, token string, secret string) {
	_, _, err := validateJWT(token, secret)

	if !errors.Is(err, expected) {
		t.Errorf("expected err: %v but got %v", expected, err)
	}
}

func TestJWT(t *testing.T) {

	secret := "my secret"
	user := "häns"

	t.Run("matches fields", func(t *testing.T) {

		abilities := []string{"köchen", "jöchen"}
		validperiod := time.Minute

		jwt, err := createJWT(user, abilities, secret, validperiod)

		if err != nil {
			t.Errorf("Expected err == nil but got: %v", err)
		}

		gotUser, gotAbilities, err := validateJWT(jwt, secret)
		if err != nil {
			t.Errorf("Expected err == nil but got: %v", err)
		}
		if gotUser != user {
			t.Errorf("expected user: %s but got %s", user, gotUser)
		}

		if !reflect.DeepEqual(abilities, gotAbilities) {
			t.Errorf("expected abilities: %v but got %v", abilities, gotAbilities)
		}
	})

	t.Run("expired token", func(t *testing.T) {
		validperiod := time.Millisecond

		jwt, err := createJWT(user, []string{}, secret, validperiod)
		if err != nil {
			t.Errorf("Expected err == nil but got: %v", err)
		}

		time.Sleep(time.Second)

		_, _, err = validateJWT(jwt, secret)
		if !errors.Is(err, errExpiredJWT) {
			t.Errorf("expected err: %v but got %v", errExpiredJWT, err)
		}
	})

	t.Run("wrongly signed", func(t *testing.T) {
		token, err := createJWT("name", []string{}, "wrong secret", time.Hour)
		if err != nil {
			t.Errorf("Expected err == nil but got: %v", err)
		}

		_, _, err = validateJWT(token, secret)
		if !errors.Is(err, jwt.ErrSignatureInvalid) {
			t.Errorf("expected err: %v but got %v", jwt.ErrSignatureInvalid, err)
		}

	})

	t.Run("missing field name", func(t *testing.T) {
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"abilities":       []string{},
			"expiration_date": time.Now().Add(time.Hour).Unix(),
		}).SignedString(secret)

		if err != nil {
			t.Errorf("Expected err == nil but got: %v", err)
		}

		validateAndExpectError(t, errInvalidJWT, token, secret)
	})
	t.Run("missing field abilities", func(t *testing.T) {
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name":            "hans",
			"expiration_date": time.Now().Add(time.Hour).Unix(),
		}).SignedString(secret)

		if err != nil {
			t.Errorf("Expected err == nil but got: %v", err)
		}

		validateAndExpectError(t, errInvalidJWT, token, secret)

	})
	t.Run("missing field expiration_date", func(t *testing.T) {
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name":      "hans",
			"abilities": []string{},
		}).SignedString(secret)

		if err != nil {
			t.Errorf("Expected err == nil but got: %v", err)
		}

		validateAndExpectError(t, errInvalidJWT, token, secret)

	})
	t.Run("no cliams", func(t *testing.T) {
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, nil).SignedString(secret)

		if err != nil {
			t.Errorf("Expected err == nil but got: %v", err)
		}

		validateAndExpectError(t, errInvalidJWT, token, secret)

	})
}
