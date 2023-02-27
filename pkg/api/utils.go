package api

import (
	"fmt"
	"mime/multipart"
	"net/http"
)

// Parse a multipart form from the request & extract a file
func readMultipartFile(req *http.Request, filename string) (multipart.File, error) {
	err := req.ParseMultipartForm(1 << 31)
	if err != nil {
		return nil, fmt.Errorf("could not parse multipart form: %s", err)
	}

	file, _, err := req.FormFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not parse image file from multipart form: %s", err)
	}

	return file, nil
}

func writeBadRequest(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))
}

func writeMethodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method not allowed"))
}
