package server_service

import (
	"github.com/chlyNiklas/lou-taylor-api/api"
	"github.com/chlyNiklas/lou-taylor-api/config"
	"github.com/chlyNiklas/lou-taylor-api/models"
)

// compiletime check for impl StrictServerInterface
var _ api.StrictServerInterface = (*Service)(nil)

type DataBase interface {
	GetAllEvents() ([]*models.Event, error)
	GetAllFutureEvents() ([]*models.Event, error)
	GetAllPastEvents() ([]*models.Event, error)
}

// Service implements StrictServerInterface from api.
type Service struct {
	cfg *config.Config
	db  DataBase
}

// New returns a pointer to a new service.
func New(cfg *config.Config, db DataBase) *Service {
	return &Service{
		cfg: cfg,
		db:  db,
	}
}
