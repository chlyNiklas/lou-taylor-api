package utils

import (
	"image"
	"io"

	// image decoders for DecodeImage
	_ "github.com/kolesa-team/go-webp/webp"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func DecodeImage(r io.Reader) (img image.Image, err error) {
	img, _, err = image.Decode(r)
	return
}
