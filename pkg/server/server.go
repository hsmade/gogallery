package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"path/filepath"
)

type Server struct {
	RootPath   string
	ListenPort int
}

// New creates a new Server
func New(listenPort int, rootPath string) (Server, error) {
	absPath, err := filepath.Abs(rootPath)
	return Server{
		ListenPort: listenPort,
		RootPath:   absPath,
	}, err
}

// Run starts a new Server
func (s *Server) Run() error {
	http.HandleFunc("/list", s.listHandler)
	http.HandleFunc("/download", s.downloadHandler)
	logrus.Infof("starting webserver on port %d", s.ListenPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.ListenPort), nil)
}
