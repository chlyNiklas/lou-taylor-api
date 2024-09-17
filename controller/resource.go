package controller

import (
	"image"
	"os"

	"github.com/chlyNiklas/lou-taylor-api/config"
	"github.com/chlyNiklas/lou-taylor-api/models"
	"github.com/google/uuid"
)

type DataBase interface {
	GetAllEvents() ([]*models.Event, error)
	GetAllFutureEvents() ([]*models.Event, error)
	GetAllPastEvents() ([]*models.Event, error)
	GetEventById(id uuid.UUID) (*models.Event, error)
	DeleteEventById(id uuid.UUID) error
	WriteEvent(event *models.Event) error
	SaveEvent(event *models.Event) error
}

type ImageStore interface {
	Read(filename string) (file *os.File, size int64, err error)
	Delete(filename string) error
	SaveImage(img image.Image) (filename string, err error)
}

// Service implements StrictServerInterface from api.
type Service struct {
	cfg *config.Config
	db  DataBase
	img ImageStore
}

// New returns a pointer to a new service.
func New(cfg *config.Config, img ImageStore, db DataBase) *Service {
	return &Service{
		cfg: cfg,
		db:  db,
		img: img,
	}
}
