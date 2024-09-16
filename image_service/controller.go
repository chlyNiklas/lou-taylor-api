package image_service

import (
	"crypto/md5"
	"encoding/hex"
	"image"
	"log"
	"os"
	"path"

	"github.com/devedge/imagehash"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"golang.org/x/image/draw"
)

func (m *Manager) SaveImage(img image.Image) (filename string, err error) {

	// rescale to max width
	if img.Bounds().Dx() > m.cfg.MaxWith {
		img = rescale(img, m.cfg.MaxWith)
		log.Printf("rescaled image")
	}

	// the filename gets constructed using a content aware hash of the SaveImage
	// this hash gets then md5 hashed again because i don't really trust it.
	// dHash: https://www.hackerfactor.com/blog/index.php?/archives/529-Kind-of-Like-That.html
	if ih, err := imagehash.Dhash(img, 128); err != nil {
		return "", err
	} else {
		h := md5.Sum(ih)
		filename = hex.EncodeToString(h[:]) + ".webp"
	}

	file, err := os.Create(path.Join(m.cfg.SavePath, filename))
	if err != nil {
		return
	}
	o, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, m.cfg.Quality)
	if err != nil {
		return
	}
	err = webp.Encode(file, img, o)
	if err != nil {
		return
	}

	return
}

func (m *Manager) Read(filename string) (file *os.File, size int64, err error) {
	filePath := path.Join(m.cfg.SavePath, filename)

	file, err = os.Open(filePath)
	if err != nil {
		return
	}
	stat, err := os.Stat(filePath)
	if err != nil {
		return
	}

	size = stat.Size()
	return
}

func (m *Manager) Delete(filename string) error {
	return nil
}

func rescale(original image.Image, width int) image.Image {

	// Define the new size
	height := int(float64(width) / float64(original.Bounds().Dx()) * float64(original.Bounds().Dy()))
	scaled := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.BiLinear.Scale(scaled, scaled.Bounds(), original, original.Bounds(), draw.Over, nil)

	return scaled
}
