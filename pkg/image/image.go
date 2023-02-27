package image

import (
	"bytes"
	"io"

	"github.com/h2non/bimg"
)

// Read an image, resize it & return the encoded bytes
func Resize(r io.Reader, height, width int) ([]byte, error) {
	buf := bytes.Buffer{}
	io.Copy(&buf, r)
	image_buffer := bimg.NewImage(buf.Bytes())

	resized, err := image_buffer.Resize(height, width)
	if err != nil {
		return nil, err
	}

	return resized, nil
}

// Read an image, rotate it & return the encoded bytes
func Rotate(r io.Reader, rotationAngle int) ([]byte, error) {
	buf := bytes.Buffer{}
	io.Copy(&buf, r)
	image_buffer := bimg.NewImage(buf.Bytes())

	rotated, err := image_buffer.Rotate(bimg.Angle(rotationAngle))
	if err != nil {
		return nil, err
	}

	return rotated, nil
}
