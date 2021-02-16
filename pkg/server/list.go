package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hsmade/gogallery/pkg/helpers"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// listHandler handles listing a directory
func (s *Server) listHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Debugf("listHandler called with %v", r.URL.RawQuery)
	filePath, ok := r.URL.Query()["path"]
	if !ok {
		Error{Message: "missing path parameter"}.Send(w)
		return
	}

	// make sure path is always inside the root path
	// This was copied from https://github.com/golang/go/blob/7211694a1e3f9eaebff7074944feead968e00e72/src/net/http/fs.go#L79
	finalPath := filepath.Join(s.RootPath, filepath.FromSlash(path.Clean("/"+filePath[0])))

	info, err := os.Stat(finalPath)
	if err != nil {
		Error{Message: fmt.Sprintf("Could not stat '%s'", finalPath), Error: err}.Send(w)
		return
	}
	if !info.IsDir() {
		Error{Message: fmt.Sprintf("Is not a directory: '%s'", finalPath)}.Send(w)
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithCancel(context.Background())
	go helpers.Keepalive(ctx, w)
	err, files := getOrCreateIndex(finalPath)
	cancel()
	if err != nil {
		Error{Message: "Could not get index", Error: err}.Send(w)
		return
	}
	_ = json.NewEncoder(w).Encode(files)
}

// getOrCreateIndex will check for the thumbs.db file to be there and create one if it's older than the directory or
// nonexistent. getOrCreateIndex will return a list of file objects
func getOrCreateIndex(path string) (error, []*interface{}) {
	//files, err := ioutil.ReadDir(finalPath)
	//if err != nil {
	//	Error{Message: "Failed to read directory", Error: err}.Send(w)
	//	return
	//}

	// serialize list of struct
	//   filename
	//   date
	//   location?
	//   thumb

	return nil, nil
}
