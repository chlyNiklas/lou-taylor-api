package database

import (
	"github.com/chlyNiklas/lou-taylor-api/models"
	"github.com/google/uuid"
)

func (d *DB) GetAllEvents() (events []*models.Event, err error) {
	err = d.conn.Find(&events).Error
	return
}
func (d *DB) GetAllFutureEvents() (events []*models.Event, err error) {
	err = d.conn.Where("date > NOW()").Find(&events).Error
	return
}
func (d *DB) GetAllPastEvents() (events []*models.Event, err error) {
	err = d.conn.Where("date =< NOW()").Find(&events).Error
	return
}
func (d *DB) GetEventById(id uuid.UUID) (event *models.Event, err error) {
	err = d.conn.First(&event, id).Error
	return
}
func (d *DB) DeleteEventById(id uuid.UUID) error {
	return d.conn.Delete(&models.Event{}, id).Error
}

func (d *DB) WriteEvent(event *models.Event) (err error) {
	return d.conn.Create(event).Error
}
func (d *DB) SaveEvent(event *models.Event) (err error) {
	return d.conn.Save(event).Error
}
