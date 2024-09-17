package models

import (
	"errors"
	"time"

	"github.com/chlyNiklas/lou-taylor-api/api"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrInvalidMissingField = errors.New("fields missing")

type Event struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Place      string
	Title      string
	Descripion string
	ImageLink  string
	Date       time.Time
}

func (e *Event) ToApiEvent() api.Event {
	return api.Event{
		Description: &e.Descripion,
		Id:          e.ID,
		ImageUrl:    e.ImageLink,
		Place:       e.Place,
		Time:        &e.Date,
		Title:       e.Title,
	}
}

func FromApiPostEvent(e *api.PostEventsJSONRequestBody) (event *Event, err error) {

	if e.Title == "" || e.Place == "" || e.Time.IsZero() {
		err = ErrInvalidMissingField
	}

	event = &Event{
		Title: e.Title,
		Place: e.Place,
		Date:  e.Time,
	}

	if e.Description != nil {
		event.Descripion = *e.Description
	}

	if e.Image != nil {
		event.ImageLink = *e.Image
	}

	return
}

func FromApiPutEvent(e *api.PutEventsEventIdJSONRequestBody) (event *Event, err error) {

	if e.Title == "" || e.Place == "" || e.Time.IsZero() {
		err = ErrInvalidMissingField
	}

	event = &Event{
		Title: e.Title,
		Place: e.Place,
		Date:  e.Time,
	}

	if e.Description != nil {
		event.Descripion = *e.Description
	}

	if e.Image != nil {
		event.ImageLink = *e.Image
	}
	return

}
