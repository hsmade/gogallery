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
	if !bytes.Equal(errorObject, bytes.TrimRight(rr.Body.Bytes(), " \n")) {
		t.Errorf("Got unexpected json: %s", rr.Body.String())
	}
}

func TestServer_listHandler(t *testing.T) {
	tests := []struct {
		name         string
		responseCode int
		path         string
		response     []byte
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
			name:         "should return list of files and directories",
			path:         "/",
			responseCode: 200,
			response:     []byte("{\"Path\":\"/Users/wim/git/gogallery/pkg/testfiles\",\"Files\":[{\"Name\":\"blue.gif\",\"Exif\":null,\"Image\":\"iVBORw0KGgoAAAANSUhEUgAAAIAAAACAEAIAAAAczCrfAAABdElEQVR4nOzTMRHAMADEsL9e+VNOYGSwhMCL/207Z5D0vQ6AlwxAmgFIMwBpBiDNAKQZgDQDkGYA0gxAmgFIMwBpBiDNAKQZgDQDkGYA0gxAmgFIMwBpBiDNAKQZgDQDkGYA0gxAmgFIMwBpBiDNAKQZgDQDkGYA0gxAmgFIMwBpBiDNAKQZgDQDkGYA0gxAmgFIMwBpBiDNAKQZgDQDkGYA0gxAmgFIMwBpBiDNAKQZgDQDkGYA0gxAmgFIMwBpBiDNAKQZgDQDkGYA0gxAmgFIMwBpBiDNAKQZgDQDkGYA0gxAmgFIMwBpBiDNAKQZgDQDkGYA0gxAmgFIMwBpBiDNAKQZgDQDkGYA0gxAmgFIMwBpBiDNAKQZgDQDkGYA0gxAmgFIMwBpBiDNAKQZgDQDkGYA0gxAmgFIMwBpBiDNAKQZgDQDkGYA0gxAmgFIMwBpBiDNAKQZgDQDkGYA0gxAmgFIMwBpBiDNAKQZgLQbAAD//26HAwHo5H+2AAAAAElFTkSuQmCC\"},{\"Name\":\"gray.png\",\"Exif\":null,\"Image\":\"iVBORw0KGgoAAAANSUhEUgAAAIAAAACACAIAAABMXPacAAABM0lEQVR4nOzRUQnAMADFwDEq/EmvjPtoTkEgZ9sX59cBr2sA1gCsAVgDsAZgDcAagDUAawDWAKwBWAOwBmANwBqANQBrANYArAFYA7AGYA3AGoA1AGsA1gCsAVgDsAZgDcAagDUAawDWAKwBWAOwBmANwBqANQBrANYArAFYA7AGYA3AGoA1AGsA1gCsAVgDsAZgDcAagDUAawDWAKwBWAOwBmANwBqANQBrANYArAFYA7AGYA3AGoA1AGsA1gCsAVgDsAZgDcAagDUAawDWAKwBWAOwBmANwBqANQBrANYArAFYA7AGYA3AGoA1AGsA1gCsAVgDsAZgDcAagDUAawDWAKwBWAOwBmANwBqANQBrANYArAFYA7AGYA3AGoA1AGsA1gCsAVgDsAZgDcBuAAAA//8VvAKDziUC7wAAAABJRU5ErkJggg==\"}],\"Directories\":[\"broken\",\"data\",\"subdir\",\"thumbdir\"]}\n"),
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

			if tt.response != nil && !bytes.Equal(tt.response, rr.Body.Bytes()) {
				t.Errorf("response expected:\n'%s'\nbut got:\n'%s'\n", tt.response, rr.Body.String())
			}
		})
	}
}
