package image_service

type ImageConfig struct {
	Quality  float32
	MaxWith  int
	SavePath string
}

// Manager saves images, and returns them
type Manager struct {
	cfg *ImageConfig
}

func New(cfg *ImageConfig) *Manager {
	return &Manager{
		cfg: cfg,
	}
}
