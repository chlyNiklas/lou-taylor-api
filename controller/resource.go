package controller

import (
	"image"
	"os"

	"github.com/chlyNiklas/lou-taylor-api/config"
	"github.com/chlyNiklas/lou-taylor-api/model"
	"github.com/google/uuid"
)

type DataBase interface {
	GetAllEvents() ([]*model.Event, error)
	GetAllFutureEvents() ([]*model.Event, error)
	GetAllPastEvents() ([]*model.Event, error)
	GetEventById(id uuid.UUID) (*model.Event, error)
	DeleteEventById(id uuid.UUID) error
	WriteEvent(event *model.Event) error
	SaveEvent(event *model.Event) error
}

type ImageStore interface {
	Read(filename string) (file *os.File, size int64, err error)
	Delete(filename string) error
	SaveImage(img image.Image) (filename string, err error)
}

type AuthService interface {
	Login(user *model.User) (token string, err error)
}

// Service implements StrictServerInterface from api.
type Service struct {
	cfg  *config.Config
	db   DataBase
	img  ImageStore
	auth AuthService
}

// New returns a pointer to a new service.
func New(cfg *config.Config, db DataBase, img ImageStore, auth AuthService) *Service {
	return &Service{
		cfg:  cfg,
		db:   db,
		img:  img,
		auth: auth,
	}
}
