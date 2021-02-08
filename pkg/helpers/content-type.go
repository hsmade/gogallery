package helpers

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

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
