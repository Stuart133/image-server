package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/stuart133/image-server/pkg/image"
)

func resizeHandler(w http.ResponseWriter, req *http.Request) {
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

	// Parse resize paramters
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

	resized, err := image.Resize(file, height, width)
	if err != nil {
		writeBadRequest(w, fmt.Sprintf("could not resize image: %s", err))
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(resized)
}
