package server_service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"reflect"

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
	var title string
	for {
		part, err := request.Body.NextPart()
		if err == io.EOF {
			break // End of parts
		}
		switch part.FormName() {
		case "title":
			// Read the title part
			titleBytes, err := io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			title = string(titleBytes)

		case "image":
			// Save the uploaded image
			imageFilePath, err = saveUploadedFile(part, "uploads")
			if err != nil {
				http.Error(w, "Error saving image", http.StatusBadRequest)
				return
			}

		default:
			// Ignore unexpected form fields
			continue
		}
	}

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

func saveUploadedFile(part *multipart.Part, uploadDir string) (string, error) {
	// Ensure upload directory exists
	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Create a new file in the upload directory
	fileName := filepath.Base(part.FileName())
	filePath := filepath.Join(uploadDir, fileName)

	outFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	// Copy the uploaded file data into the new file
	_, err = io.Copy(outFile, part)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
