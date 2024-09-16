package models

import (
	"time"

	"github.com/chlyNiklas/lou-taylor-api/api"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Place      string
	Title      string
	Descripion string
	ImagePath  string
	Date       time.Time
}

func (e *Event) ToApiEvent() *api.Event {
	return &api.Event{
		Description: &e.Descripion,
		Id:          e.ID,
		ImageUrl:    e.ImagePath,
		Place:       e.Place,
		Time:        &e.Date,
		Title:       e.Title,
	}

}
