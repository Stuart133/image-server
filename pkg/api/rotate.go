package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/h2non/bimg"
)

func rotateHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}

	// Read the multipart form & extract the image file
	file, err := read_multipart_file(req)
	if err != nil {
		writeBadRequest(w, err.Error())
		return
	}
	defer file.Close()

	// Read rotation paramter
	rotation, err := strconv.Atoi(req.FormValue("rotation"))
	if err != nil {
		writeBadRequest(w, fmt.Sprintf("could not parse rotation: %s", err))
		return
	}

	// Read the image file & copy into an image buffer
	buf := bytes.Buffer{}
	io.Copy(&buf, file)
	image_buffer := bimg.NewImage(buf.Bytes())

	// Do the actual resizing
	rotated, err := image_buffer.Rotate(bimg.Angle(rotation))
	if err != nil {
		writeBadRequest(w, fmt.Sprintf("could not rotate image: %s", err))
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(rotated)
}
