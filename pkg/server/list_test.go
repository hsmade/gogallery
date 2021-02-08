package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func TestServer_listHandler_no_path(t *testing.T) {
	s := &Server{}

	req, err := http.NewRequest("GET", "/list", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.listHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != 500 {
		t.Errorf("expected code 500, but got: %d", rr.Code)
	}

	errorObject, _ := json.Marshal(Error{Message: "missing path parameter", Code: 500})
	if bytes.Compare(errorObject, bytes.TrimRight(rr.Body.Bytes(), " \n")) != 0 {
		t.Errorf("Got unexpected json: %s", rr.Body.String())
	}
}

func TestServer_listHandler(t *testing.T) {
	tests := []struct {
		name         string
		responseCode int
		path         string
	}{
		{
			name:         "should error on nonexistent path",
			path:         "/a/b/c",
			responseCode: 500,
		},
		{
			name:         "should error on file",
			path:         "blue.gif",
			responseCode: 500,
		},
		{
			name: "should return list of files",
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

			req, err := http.NewRequest("GET", fmt.Sprintf("/list?path=%s", tt.path), nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(s.listHandler)
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.responseCode {
				t.Errorf("responsecode expected: %d, but got: %d", tt.responseCode, rr.Code)
				t.Log(rr.Body.String())
			}

			// TODO: test response body
		})
	}
}
