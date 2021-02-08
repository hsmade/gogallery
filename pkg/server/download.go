package server

import (
	"github.com/hsmade/gogallery/pkg/helpers"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

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

	contentType, err := helpers.GetFileContentType(file)
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
