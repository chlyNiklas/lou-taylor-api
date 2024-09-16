package database

import "github.com/chlyNiklas/lou-taylor-api/models"

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
