package server

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func TestMain(m *testing.M) {
	logrus.SetLevel(logrus.DebugLevel)
	code := m.Run()
	os.Exit(code)
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
			file:         "testfiles/gray.png",
			path:         "gray.png",
			contentType:  "image/png",
		},
		{
			name:         "download gif picture",
			responseCode: 200,
			file:         "testfiles/blue.gif",
			path:         "blue.gif",
			contentType:  "image/gif",
		},
		{
			name:         "download gif picture, subdir",
			responseCode: 200,
			file:         "testfiles/subdir/green.gif",
			path:         "subdir/green.gif",
			contentType:  "image/gif",
		},
		{
			name:         "download outside of path",
			path:         "../server.go",
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
				RootPath: path.Join(cwd, "testfiles"),
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

				expectedContent, err := ioutil.ReadAll(expectedFile)
				if err != nil {
					t.Fatalf("Can't read expected file: %v", err)
				}

				if bytes.Compare(expectedContent, rr.Body.Bytes()) > 0 {
					t.Error("did not receive expected content")
					t.Logf("expected: %s\ngot: %s", expectedContent, rr.Body.Bytes())
				}
			}
		})
	}
}

func TestServer_listHandler(t *testing.T) {
	type fields struct {
		RootPath   string
		ListenPort int
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &Server{
				RootPath:   tt.fields.RootPath,
				ListenPort: tt.fields.ListenPort,
			}
		})
	}
}

func TestServer_GetFileContentType(t *testing.T) {
	tests := []struct {
		name                string
		file                string
		expectedContentType string
		wantError           bool
	}{
		{
			name:                "valid gif",
			file:                "testfiles/green.gif",
			expectedContentType: "image/gif",
			wantError:           false,
		},
		{
			name:                "valid png",
			file:                "testfiles/gray.png",
			expectedContentType: "image/png",
			wantError:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Open(tt.file)
			if err != nil {
				t.Fatalf("can't open requested file: %v", err)
			}
			contentType, err := GetFileContentType(file)
			if err != nil && !tt.wantError {
				t.Errorf("got unexpected error: %v", err)
			}

			if contentType != tt.expectedContentType {
				t.Errorf("expected content type: %s but got: %s", tt.expectedContentType, contentType)
			}
		})
	}
}
