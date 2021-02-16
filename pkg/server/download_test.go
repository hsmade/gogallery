package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func TestServer_downloadHandler_no_path(t *testing.T) {
	s := &Server{}

	req, err := http.NewRequest("GET", "/download", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.downloadHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != 500 {
		t.Errorf("expected code 500, but got: %d", rr.Code)
	}

	errorObject, _ := json.Marshal(Error{Message: "missing path parameter", Code: 500})
	if bytes.Equal(errorObject, bytes.TrimRight(rr.Body.Bytes(), " \n")) {
		t.Errorf("Got unexpected json: %s", rr.Body.String())
	}
}

func TestServer_downloadHandler(t *testing.T) {
	tests := []struct {
		name         string
		responseCode int
		file         string
		contentType  string
		path         string
	}{
		{
			name:         "download png picture",
			responseCode: 200,
			file:         "../testfiles/gray.png",
			path:         "gray.png",
			contentType:  "image/png",
		},
		{
			name:         "download gif picture",
			responseCode: 200,
			file:         "../testfiles/blue.gif",
			path:         "blue.gif",
			contentType:  "image/gif",
		},
		{
			name:         "download gif picture, subdir",
			responseCode: 200,
			file:         "../testfiles/subdir/green.gif",
			path:         "subdir/green.gif",
			contentType:  "image/gif",
		},
		{
			name:         "download outside of path",
			path:         "../pkg/server.go",
			responseCode: 404,
		},
		{
			name:         "download missing file",
			path:         "nosuchfile.png",
			responseCode: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cwd, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}

			s := &Server{
				RootPath: path.Join(cwd, "../testfiles"),
			}

			req, err := http.NewRequest("GET", fmt.Sprintf("/download?path=%s", tt.path), nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(s.downloadHandler)
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.responseCode {
				t.Errorf("responsecode expected: %d, but got: %d", tt.responseCode, rr.Code)
				t.Log(rr.Body.String())
			}

			if tt.contentType != "" && rr.Header().Get("content-type") != tt.contentType {
				t.Errorf("Content-type wrong. Expected: %s but got %s", tt.contentType, rr.Header().Get("content-type"))
			}

			if tt.file != "" {
				expectedFile, err := os.Open(tt.file)
				if err != nil {
					t.Fatalf("Can't open expected file: %v", err)
				}
				defer expectedFile.Close()

				expectedContent, err := ioutil.ReadAll(expectedFile)
				if err != nil {
					t.Fatalf("Can't read expected file: %v", err)
				}

				if bytes.Equal(expectedContent, rr.Body.Bytes()) {
					t.Error("did not receive expected content")
					t.Logf("expected: %s\ngot: %s", expectedContent, rr.Body.Bytes())
				}
			}
		})
	}
}
