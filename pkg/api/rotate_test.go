package api

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/v3/assert"
)

func Test_Rotate(t *testing.T) {
	tests := []struct {
		name string

		sendFile bool
		rotation string
		method   string

		wantCode int
	}{
		{
			name: "Good request",

			sendFile: true,
			rotation: "90",
			method:   "POST",

			wantCode: 200,
		},
		{
			name: "Missing file",

			sendFile: false,
			rotation: "90",
			method:   "POST",

			wantCode: 400,
		},
		{
			name: "Invalid rotation",

			sendFile: true,
			rotation: "hmm",
			method:   "POST",

			wantCode: 400,
		},
		{
			name: "Incorrect method",

			sendFile: false,
			method:   "GET",

			wantCode: 405,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create multipart file for testing
			buf := bytes.Buffer{}
			mw := multipart.NewWriter(&buf)

			// Create image part
			fw, err := mw.CreateFormFile("image", "balloons.jpg")
			assert.NilError(t, err)
			_, err = io.Copy(fw, mustOpen("../../data/balloons.jpg"))
			assert.NilError(t, err)

			// Add rotation parameter
			fw, err = mw.CreateFormField("rotation")
			fw.Write([]byte(tt.rotation))

			mw.Close()

			var req *http.Request
			if tt.sendFile {
				req = httptest.NewRequest(tt.method, "/resize", &buf)
				req.Header.Set("Content-Type", mw.FormDataContentType())
			} else {
				req = httptest.NewRequest(tt.method, "/resize", nil)
			}
			m := httptest.NewRecorder()

			rotateHandler(m, req)

			res := m.Result()
			assert.NilError(t, err)

			assert.Equal(t, res.StatusCode, tt.wantCode)
			if tt.wantCode == http.StatusOK {
				assert.Equal(t, "application/octet-stream", res.Header.Get("Content-Type"))
			}
		})
	}
}
