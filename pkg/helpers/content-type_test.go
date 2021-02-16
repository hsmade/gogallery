package helpers

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	logrus.SetLevel(logrus.DebugLevel)
	code := m.Run()
	os.Exit(code)
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
			file:                "../testfiles/blue.gif",
			expectedContentType: "image/gif",
			wantError:           false,
		},
		{
			name:                "valid png",
			file:                "../testfiles/gray.png",
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
			defer file.Close()
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
