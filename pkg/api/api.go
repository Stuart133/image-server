package api

import "net/http"

type API struct {
}

// Register API handlers
func RegisterHandlers() {
	http.HandleFunc("/resize", resizeHandler)
	http.HandleFunc("/rotate", rotateHandler)
}
