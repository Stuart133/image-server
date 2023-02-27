package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/stuart133/image-server/pkg/image"
)

func rotateHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}

	file, err := readMultipartFile(req, "image")
	if err != nil {
		writeBadRequest(w, err.Error())
		return
	}
	defer file.Close()

	// Parse the rotation paramter
	rotation, err := strconv.Atoi(req.FormValue("rotation"))
	if err != nil {
		writeBadRequest(w, fmt.Sprintf("could not parse rotation: %s", err))
		return
	}

	rotated, err := image.Rotate(file, rotation)
	if err != nil {
		writeBadRequest(w, fmt.Sprintf("could not rotate image: %s", err))
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(rotated)
}
