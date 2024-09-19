package authentication

import (
	"time"

	"github.com/chlyNiklas/lou-taylor-api/model"
	"github.com/getkin/kin-openapi/openapi3"
)

// Config holds all config values
type Config struct {
	Admin       *model.User   `toml:"admin" comment:"Credentials of admin user"`
	JWTSecret   string        `toml:"secret" comment:"secret to sign JWT with"`
	ValidPeriod time.Duration `toml:"valid_period" comment:"timeperiod the token is valid in nanoseconds"`
}

// Service provides all authentication methods & middlewares
type Service struct {
	cfg  *Config
	spec *openapi3.T
}

// New returns a new auth service
func New(cfg *Config, spec *openapi3.T) *Service {
	return &Service{
		cfg:  cfg,
		spec: spec,
	}
}
