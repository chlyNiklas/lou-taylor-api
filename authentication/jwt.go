package authentication

import (
	"errors"
	"fmt"
	"time"

	"github.com/chlyNiklas/lou-taylor-api/utils"
	"github.com/golang-jwt/jwt/v5"
)

var errInvalidJWT = errors.New("invalid JWT")
var errExpiredJWT = errors.New("expired JWT")

func createJWT(name string, abilities []string, secret string, vialid time.Duration) (token string, err error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":            name,
		"abilities":       abilities,
		"expiration_date": time.Now().Add(vialid).Unix(),
	}).SignedString([]byte(secret))
}

func validateJWT(tokenString string, secret string) (user string, abilities []string, err error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return
	}

	// extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errInvalidJWT
		return
	}

	// extract user
	user, ok = claims["name"].(string)
	if !ok {
		err = errInvalidJWT
		return
	}

	exp_date, ok := claims["expiration_date"].(float64)
	if !ok {
		err = errInvalidJWT
		return
	}

	if int64(exp_date) < time.Now().Unix() {
		err = errExpiredJWT
		return
	}

	// extract abilities
	abilitiesSlice, ok := claims["abilities"].([]any)
	if !ok {
		err = errInvalidJWT
		return
	}

	abilities = utils.Map(abilitiesSlice, func(ab any) string {
		ability, o := ab.(string)
		if !o {
			ok = false
		}
		return ability
	})

	if !ok {
		err = errInvalidJWT
		return
	}

	return
}
