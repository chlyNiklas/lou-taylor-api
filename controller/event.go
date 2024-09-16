package controller

import (
	"context"
	"github.com/chlyNiklas/lou-taylor-api/api"
	"github.com/chlyNiklas/lou-taylor-api/models"
	"github.com/chlyNiklas/lou-taylor-api/utils"
)

func (s *Service) GetEvents(ctx context.Context, request api.GetEventsRequestObject) (api.GetEventsResponseObject, error) {

	var getEvents func() ([]*models.Event, error)
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
		func(e *models.Event) api.Event {
			return *e.ToApiEvent()
		})

	return api.GetEvents200JSONResponse(apiEvent), err
}

func (s *Service) PostEvents(ctx context.Context, request api.PostEventsRequestObject) (api.PostEventsResponseObject, error) {

	// Implementation here
	return api.PostEvents400Response{}, nil
}

func (s *Service) DeleteEventsEventId(ctx context.Context, request api.DeleteEventsEventIdRequestObject) (api.DeleteEventsEventIdResponseObject, error) {
	// Implementation here
	return nil, nil
}

func (s *Service) GetEventsEventId(ctx context.Context, request api.GetEventsEventIdRequestObject) (api.GetEventsEventIdResponseObject, error) {
	// Implementation here
	return nil, nil
}

func (s *Service) PutEventsEventId(ctx context.Context, request api.PutEventsEventIdRequestObject) (api.PutEventsEventIdResponseObject, error) {
	// Implementation here
	return nil, nil
}
