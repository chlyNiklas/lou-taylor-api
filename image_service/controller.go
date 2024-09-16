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

// SaveImage processes, resizes, and saves the given image to disk as a WebP file.
// The image is rescaled to the maximum width specified in the Manager's configuration.
// The filename is generated using a content-aware hash (dHash) that is further hashed with MD5.
//
// Parameters:
//   - img: the image.Image object to be saved.
//
// Returns:
//   - filename: the generated filename for the saved image.
//   - err: any error encountered during the process.
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

// Read retrieves a saved image file and its size in bytes from the configured save path.
//
//	If there is an error, it will be of type \*PathError.
//
// Parameters:
//   - filename: the name of the file to read.
//
// Returns:
//   - file: a pointer to the opened file.
//   - size: the size of the file in bytes.
//   - err: any error encountered during the file access.
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

// Delete removes the specified image file from the configured save path.
// If there is an error, it will be of type \*PathError.
func (m *Manager) Delete(filename string) error {
	return os.Remove(path.Join(m.cfg.SavePath, filename))
}

// rescale resizes an image to the specified width while maintaining the aspect ratio.
// It uses bilinear interpolation for resizing.
func rescale(original image.Image, width int) image.Image {

	// Define the new size
	height := int(float64(width) / float64(original.Bounds().Dx()) * float64(original.Bounds().Dy()))
	scaled := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.BiLinear.Scale(scaled, scaled.Bounds(), original, original.Bounds(), draw.Over, nil)

	return scaled
}
