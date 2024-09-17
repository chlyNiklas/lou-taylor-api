package authentication

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
)

var cfg = &Config{
	JWTSecret:   []byte("secret"),
	ValidPeriod: time.Hour,
}

func specWithPath(ps []struct {
	string
	*openapi3.PathItem
}) *openapi3.T {

	paths := &openapi3.Paths{}

	for _, path := range ps {
		paths.Set(path.string, path.PathItem)
	}

	return &openapi3.T{
		Components: &openapi3.Components{
			SecuritySchemes: openapi3.SecuritySchemes{
				"BearerAuth": &openapi3.SecuritySchemeRef{
					Value: &openapi3.SecurityScheme{
						Type:         "http",
						Scheme:       "bearer",
						BearerFormat: "JWT", // JWT format for the bearer token
					},
				},
			},
		},
		Paths: paths,
	}
}

func specWithSecurityForPath(path string, security *openapi3.SecurityRequirements) *openapi3.T {

	return specWithPath([]struct {
		string
		*openapi3.PathItem
	}{
		{path,
			&openapi3.PathItem{
				Get: &openapi3.Operation{
					Summary:  "Health Check",
					Security: security,
				},

				Put: &openapi3.Operation{
					Summary:  "Health Check",
					Security: security,
				},

				Patch: &openapi3.Operation{
					Summary:  "Health Check",
					Security: security,
				},

				Post: &openapi3.Operation{
					Summary:  "Health Check",
					Security: security,
				},

				Delete: &openapi3.Operation{
					Summary:  "Health Check",
					Security: security,
				},
			},
		},
	})
}

func setToken(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
}

func testForAllMethods(t *testing.T, name string, test func(t *testing.T, method string)) {

	for _, method := range []string{"GET", "PUT", "PATCH", "POST", "DELETE"} {
		t.Run(name+"_"+method, func(t *testing.T) {
			test(t, method)
		})
	}
}

func TestAuthentication_SecuredNoAbilities(t *testing.T) {

	spec := specWithSecurityForPath("/health-check", &openapi3.SecurityRequirements{
		{
			"BearerAuth": []string{}, // Security requirement for bearer authentication
		},
	})

	testForAllMethods(t, "missing token -> Forbidden", func(t *testing.T, method string) {

		req, _ := http.NewRequest(method, "/health-check", nil)
		rr := httptest.NewRecorder()

		m := New(cfg, spec)

		htt := m.Authentication(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { return }))

		htt.ServeHTTP(rr, req)

		expected := "403 Forbidden"
		if rr.Result().Status != expected {
			t.Errorf("Expected %s got: %s", expected, rr.Result().Status)
		}

	})
	testForAllMethods(t, "with token on existing path -> OK", func(t *testing.T, method string) {

		req, _ := http.NewRequest(method, "/health-check", nil)
		rr := httptest.NewRecorder()

		token, _ := createJWT("admin", []string{}, cfg.JWTSecret, time.Hour)

		setToken(req, token)

		// run
		m := New(cfg, spec)
		htt := m.Authentication(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) }))
		htt.ServeHTTP(rr, req)

		if rr.Result().Status != "200 OK" {
			t.Errorf("Expected 403 got: %s", rr.Result().Status)
		}

	})
	testForAllMethods(t, "with token on notexisting path -> OK", func(t *testing.T, method string) {

		req, _ := http.NewRequest(method, "/wrong-path", nil)
		rr := httptest.NewRecorder()

		token, _ := createJWT("admin", []string{}, cfg.JWTSecret, time.Hour)

		setToken(req, token)

		// run
		m := New(cfg, spec)
		htt := m.Authentication(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNotFound) }))
		htt.ServeHTTP(rr, req)

		expected := "404 Not Found"
		if rr.Result().Status != expected {
			t.Errorf("Expected %s got: %s", expected, rr.Result().Status)
		}

	})

}

func TestMatchPath(t *testing.T) {
	tests := []struct {
		a     string
		b     string
		match bool
	}{
		{
			a:     "/asdf/{joggölid}/asdf",
			b:     "/asdf/ruedi",
			match: false,
		},
		{
			a:     "/asdf/{joggölid}",
			b:     "/asdf/ruedi",
			match: true,
		},
		{
			a:     "/asdf/{joggölid}",
			b:     "/asdf/ruedi/asdf",
			match: false,
		},
		{
			a:     "/asdf/{joggölid}/asdf/{someId}",
			b:     "/asdf/ruedi/asdf/3",
			match: true,
		},
	}

	for _, test := range tests {
		if matchPath(test.a, test.b) != test.match {
			t.Errorf("Expected %s & %s to %v", test.a, test.b, test.match)
		}
	}
}
