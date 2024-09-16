// Package image_service provides functionality for saving and retrieving images.
package image_service

type ImageConfig struct {
	Quality  float32 // Quality defines the quality of the stored webp Values from 1-100
	MaxWith  int     // MaxWith defines the maximum with an image gets stored with
	SavePath string  // SavePath specifies the directory where the image will be saved.
}

// Manager saves images, and returns them
type Manager struct {
	cfg *ImageConfig
}

// New creates a new instance of Manager with the provided image configuration.
func New(cfg *ImageConfig) *Manager {
	return &Manager{
		cfg: cfg,
	}
}
