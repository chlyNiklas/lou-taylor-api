package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/chlyNiklas/lou-taylor-api/api"
	"github.com/chlyNiklas/lou-taylor-api/config"
	"github.com/chlyNiklas/lou-taylor-api/utils"
	"github.com/getkin/kin-openapi/openapi3"
)

type Middleware struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Middleware {
	return &Middleware{
		cfg: cfg,
	}
}

func (m *Middleware) Authentication(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// get swagger specification
		spec, err := api.GetSwagger()
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		// spec for current path
		path := matchPaths(spec.Paths, r.URL.Path)

		// if no matching path, serve
		if path == nil {
			next.ServeHTTP(w, r)
			return
		}

		// get operation
		var operation *openapi3.Operation
		switch r.Method {
		case "GET":
			operation = path.Get
		case "PUT":
			operation = path.Put
		case "POST":
			operation = path.Post
		case "DELETE":
			operation = path.Delete
		case "PATCH":
			operation = path.Patch
		}

		// if no Security serve
		if operation.Security == nil {
			next.ServeHTTP(w, r)
			return
		}

		jwt, err := extractToken(r)
		if err != nil {

			w.WriteHeader(http.StatusForbidden)
			return
		}

		// validate JWT
		_, abilies, err := utils.ValidateJWT(jwt, m.cfg.JWTSecret)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		for _, req := range *operation.Security {
			barear := req["bearerAuth"]
			if len(barear) == 0 {
				continue
			}
			if utils.MachesAny(barear, abilies) {
				next.ServeHTTP(w, r)
				return
			}

		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func extractToken(r *http.Request) (token string, err error) {
	tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

	if tokenHeader == "" {
		err = errors.New("missing token")
		return
	}

	token = strings.TrimPrefix(tokenHeader, "Bearer ")

	if token == "" {
		err = errors.New("malformed token")
	}
	return
}

func matchPaths(paths *openapi3.Paths, path string) *openapi3.PathItem {
	for _, value := range paths.InMatchingOrder() {
		if matchPath(value, path) {
			return paths.Find(value)
		}
	}
	return nil
}

func matchPath(a, b string) bool {
	ia := 0
	ib := 0

	capt := false

	a = a + "/"
	b = b + "/"

	for ia < len(a) || ib < len(b) {
		if ia >= len(a) || ib >= len(b) {
			return false
		}
		if capt {
			for a[ia] != '/' {
				ia++
			}
			for b[ib] != '/' {
				ib++
			}
			capt = false
		} else {
			if a[ia] == '{' {
				capt = true
				continue
			}

			if a[ia] != b[ib] {
				return false
			}
			ia++
			ib++

		}
	}
	return true
}
