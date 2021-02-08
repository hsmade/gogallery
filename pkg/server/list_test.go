package server

import (
	"net/http"
	"testing"
)

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
