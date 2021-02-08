package server

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

// listHandler handles listing a directory
func (s *Server) listHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Debugf("listHandler called with %v", r.URL.RawQuery)

	w.WriteHeader(500)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode([]string{})
}
