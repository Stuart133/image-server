package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/h2non/bimg"
)

func main() {
	http.HandleFunc("/resize", resize)
	http.HandleFunc("/rotate", rotate)

	http.ListenAndServe(":8000", nil)
}

func rotate(w http.ResponseWriter, req *http.Request) {
	// Read the multipart form & extract the image file
	file, err := read_multipart_file(req)
	if err != nil {
		writeBadResponse(w, err.Error())
	}
	defer file.Close()

	// Read rotation paramter
	rotation, err := strconv.Atoi(req.FormValue("rotation"))
	if err != nil {
		writeBadResponse(w, fmt.Sprintf("could not parse rotation: %s", err))
		return
	}

	// Read the image file & copy into an image buffer
	buf := bytes.Buffer{}
	io.Copy(&buf, file)
	image_buffer := bimg.NewImage(buf.Bytes())

	// Do the actual resizing
	rotated, err := image_buffer.Rotate(bimg.Angle(rotation))
	if err != nil {
		writeBadResponse(w, fmt.Sprintf("could not rotate image: %s", err))
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(rotated)
}

func resize(w http.ResponseWriter, req *http.Request) {
	// Read the multipart form & extract the image file
	file, err := read_multipart_file(req)
	if err != nil {
		writeBadResponse(w, err.Error())
	}
	defer file.Close()

	// Read resize paramters
	width, err := strconv.Atoi(req.FormValue("width"))
	if err != nil {
		writeBadResponse(w, fmt.Sprintf("could not parse output image width: %s", err))
		return
	}
	height, err := strconv.Atoi(req.FormValue("height"))
	if err != nil {
		writeBadResponse(w, fmt.Sprintf("could not parse output image height: %s", err))
		return
	}

	// Read the image file & copy into an image buffer
	buf := bytes.Buffer{}
	io.Copy(&buf, file)
	image_buffer := bimg.NewImage(buf.Bytes())

	// Do the actual resizing
	resized, err := image_buffer.Resize(width, height)
	if err != nil {
		writeBadResponse(w, fmt.Sprintf("could not resize image: %s", err))
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(resized)
}

func read_multipart_file(req *http.Request) (multipart.File, error) {
	err := req.ParseMultipartForm(1 << 31)
	if err != nil {
		return nil, fmt.Errorf("could not parse multipart form: %s", err)
	}

	// Read image file
	file, _, err := req.FormFile("image")
	if err != nil {
		return nil, fmt.Errorf("could not parse image file from multipart form: %s", err)
	}

	return file, nil
}

func writeBadResponse(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))
}
