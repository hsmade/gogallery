package server

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

type Server struct {
	RootPath   string
	ListenPort int
}

type Error struct {
	Message string
	Error   error
	Code    int
}

func (e Error) Send(w http.ResponseWriter) {
	if e.Code == 0 {
		e.Code = 500
	}
	logrus.WithError(e.Error).Errorf("sending error: %v", e.Message)
	w.WriteHeader(e.Code)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(e)
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

// listHandler handles listing a directory
func (s *Server) listHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Debugf("listHandler called with %v", r.URL.RawQuery)

	w.WriteHeader(500)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode([]string{})
}

// downloadHandler handles downloading a file
func (s *Server) downloadHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Debugf("downloadHandler called with %v", r.URL.RawQuery)
	filePath, ok := r.URL.Query()["path"]
	if !ok {
		Error{Message: "missing path parameter"}.Send(w)
	}

	// make sure path is valid and inside the root path
	finalPath := filepath.Join(s.RootPath, filepath.FromSlash(path.Clean("/"+filePath[0])))

	file, err := os.Open(finalPath)
	if err != nil {
		Error{
			Message: "error while opening file",
			Error:   err,
			Code:    404,
		}.Send(w)
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		Error{
			Message: "error while reading file",
			Error:   err,
		}.Send(w)
		return
	}

	contentType, err := GetFileContentType(file)
	if err != nil {
		Error{
			Message: "error while determining content type for file",
			Error:   err,
		}.Send(w)
		return
	}

	w.Header().Set("Content-Type", contentType)
	_, _ = w.Write(content)
}

// GetFileContentType finds the content type for a file
func GetFileContentType(file *os.File) (string, error) {
	logrus.Debugf("determining content type for file %v", file.Name())
	buffer := make([]byte, 512)

	_, _ = file.Seek(0, 0)
	_, err := file.Read(buffer)
	if err != nil {
		return "", errors.Wrap(err, "reading file")
	}

	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
