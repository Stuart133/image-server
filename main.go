package main

import (
	"net/http"

	"github.com/stuart133/image-server/pkg/api"
)

func main() {
	api.RegisterHandlers()

	http.ListenAndServe(":8000", nil)
}
