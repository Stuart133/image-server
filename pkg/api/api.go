package api

import "net/http"

// Register API handlers
func RegisterHandlers() {
	http.HandleFunc("/resize", resizeHandler)
	http.HandleFunc("/rotate", rotateHandler)
}
