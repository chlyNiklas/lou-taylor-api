package database

import (
	"github.com/chlyNiklas/lou-taylor-api/model"
	"github.com/google/uuid"
)

func (d *DB) GetAllEvents() (events []*model.Event, err error) {
	err = d.conn.Find(&events).Error
	return
}
func (d *DB) GetAllFutureEvents() (events []*model.Event, err error) {
	err = d.conn.Where("date > NOW()").Find(&events).Error
	return
}
func (d *DB) GetAllPastEvents() (events []*model.Event, err error) {
	err = d.conn.Where("date =< NOW()").Find(&events).Error
	return
}
func (d *DB) GetEventById(id uuid.UUID) (event *model.Event, err error) {
	err = d.conn.First(&event, id).Error
	return
}
func (d *DB) DeleteEventById(id uuid.UUID) error {
	return d.conn.Delete(&model.Event{}, id).Error
}

func (d *DB) WriteEvent(event *model.Event) (err error) {
	return d.conn.Create(event).Error
}
func (d *DB) SaveEvent(event *model.Event) (err error) {
	return d.conn.Save(event).Error
}
