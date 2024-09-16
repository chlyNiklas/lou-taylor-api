package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidJWT = errors.New("invalid JWT")
var ErrExpiredJWT = errors.New("expired JWT")

func CreateJWT(username string, abilities []string, secret []byte) (token string, err error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":            username,
		"abilities":       abilities,
		"expiration_date": time.Now().Add(time.Hour * 3).Unix(),
	}).SignedString(secret)

}

func ValidateJWT(tokenString string, secret []byte) (user string, abilities []string, err error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})
	if err != nil {
		return
	}

	// extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = ErrInvalidJWT
		return
	}

	// extract user
	user, ok = claims["name"].(string)
	if !ok {
		err = ErrInvalidJWT
		return
	}

	exp_date, ok := claims["expiration_date"].(float64)
	if !ok {
		err = ErrInvalidJWT
		return
	}

	if int64(exp_date) < time.Now().Unix() {
		err = ErrExpiredJWT
		return
	}

	// extract abilities
	abilitiesSlice, ok := claims["abilities"].([]any)
	if !ok {
		err = ErrInvalidJWT
		return
	}

	abilities = Map(abilitiesSlice, func(ab any) string {
		ability, o := ab.(string)
		if !o {
			ok = false
		}
		return ability
	})

	if !ok {
		err = ErrInvalidJWT
		return
	}

	return
}
