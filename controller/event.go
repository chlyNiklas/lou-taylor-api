package controller

import (
	"context"

	"github.com/chlyNiklas/lou-taylor-api/api"
	"github.com/chlyNiklas/lou-taylor-api/model"
	"github.com/chlyNiklas/lou-taylor-api/utils"
	"gorm.io/gorm"
)

func (s *Service) GetEvents(ctx context.Context, request api.GetEventsRequestObject) (api.GetEventsResponseObject, error) {

	var getEvents func() ([]*model.Event, error)
	switch utils.DSNTE((*string)(request.Params.Status)) {
	case "past":
		getEvents = s.db.GetAllPastEvents
	case "future":
		getEvents = s.db.GetAllFutureEvents
	default:
		getEvents = s.db.GetAllEvents
	}

	events, err := getEvents()
	if err != nil {
		return nil, err
	}

	apiEvent := utils.Map(
		events,
		func(e *model.Event) api.Event {
			return e.ToApiEvent()
		})

	return api.GetEvents200JSONResponse(apiEvent), err
}

func (s *Service) PostEvents(ctx context.Context, request api.PostEventsRequestObject) (api.PostEventsResponseObject, error) {

	event, err := model.FromApiPostEvent(request.Body)
	if err != nil {
		if err == model.ErrInvalidMissingField {
			return api.PostEvents400Response{}, nil
		}
		return nil, err
	}

	err = s.db.WriteEvent(event)

	// Implementation here
	return api.PostEvents201JSONResponse(event.ToApiEvent()), err
}

func (s *Service) DeleteEventsEventId(ctx context.Context, request api.DeleteEventsEventIdRequestObject) (api.DeleteEventsEventIdResponseObject, error) {

	err := s.db.DeleteEventById(request.EventId)

	if err == gorm.ErrRecordNotFound {
		return api.DeleteEventsEventId404Response{}, nil
	}
	// Implementation here
	return api.DeleteEventsEventId204Response{}, nil
}

func (s *Service) GetEventsEventId(ctx context.Context, request api.GetEventsEventIdRequestObject) (response api.GetEventsEventIdResponseObject, err error) {

	if event, err := s.db.GetEventById(request.EventId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return api.GetEventsEventId404Response{}, nil
		}
		return nil, err
	} else {
		response = api.GetEventsEventId200JSONResponse(event.ToApiEvent())
	}
	// Implementation here
	return
}

func (s *Service) PutEventsEventId(ctx context.Context, request api.PutEventsEventIdRequestObject) (api.PutEventsEventIdResponseObject, error) {
	event, err := model.FromApiPutEvent(request.Body)
	if err != nil {
		if err == model.ErrInvalidMissingField {
			return api.PutEventsEventId400Response{}, nil
		}
		return nil, err
	}

	err = s.db.SaveEvent(event)

	return api.PutEventsEventId200JSONResponse(event.ToApiEvent()), err
}
