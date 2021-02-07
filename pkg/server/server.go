package server

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

type Server struct {
	RootPath   string
	ListenPort int
}

type Error struct {
	ErrorMessage string
	Error        error
}

func (e Error) Send(w http.ResponseWriter) {
	logrus.WithError(e.Error).Errorf("sending error: %s", e.ErrorMessage)
	w.WriteHeader(500)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Error{"missing path parameter", nil})
}

// New creates a new Server
func New(listenPort int, rootPath string) Server {
	return Server{
		ListenPort: listenPort,
		RootPath:   rootPath,
	}
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
	json.NewEncoder(w).Encode([]string{})
}

// downloadHandler handles downloading a file
func (s *Server) downloadHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Debugf("downloadHandler called with %v", r.URL.RawQuery)
	path, ok := r.URL.Query()["path"]
	if !ok {
		Error{"missing path parameter", nil}.Send(w)
	}

	file, err := os.Open(path[0])
	if err != nil {
		Error{
			ErrorMessage: "error while opening file",
			Error:        err,
		}.Send(w)
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		Error{
			ErrorMessage: "error while reading file",
			Error:        err,
		}.Send(w)
		return
	}

	contentType, err := GetFileContentType(file)
	if err != nil {
		Error{
			ErrorMessage: "error while determining content type for file",
			Error:        err,
		}.Send(w)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Write(content)
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
