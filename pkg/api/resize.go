package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/h2non/bimg"
)

func resizeHandler(w http.ResponseWriter, req *http.Request) {
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

	// Read resize paramters
	width, err := strconv.Atoi(req.FormValue("width"))
	if err != nil {
		writeBadRequest(w, fmt.Sprintf("could not parse output image width: %s", err))
		return
	}
	height, err := strconv.Atoi(req.FormValue("height"))
	if err != nil {
		writeBadRequest(w, fmt.Sprintf("could not parse output image height: %s", err))
		return
	}

	// Read the image file & copy into an image buffer
	buf := bytes.Buffer{}
	io.Copy(&buf, file)
	image_buffer := bimg.NewImage(buf.Bytes())

	// Do the actual resizing
	resized, err := image_buffer.Resize(width, height)
	if err != nil {
		writeBadRequest(w, fmt.Sprintf("could not resize image: %s", err))
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(resized)
}
