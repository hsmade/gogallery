package thumbs

import (
	"fmt"
	"github.com/barasher/go-exiftool"
	"github.com/pkg/errors"
)

// ExtractExif gets the exif data from a file
func ExtractExif(path string) (map[string]string, error) {
	et, err := exiftool.NewExiftool()
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize exiftool")
	}

	fileInfo := et.ExtractMetadata(path)[0]
	if fileInfo.Err != nil {
		return nil, errors.Wrap(fileInfo.Err, fmt.Sprintf("extracting exif data from '%s", path))
	}

	result := make(map[string]string, len(fileInfo.Fields))
	for k, _ := range fileInfo.Fields {
		result[k], err = fileInfo.GetString(k)
		if err != nil {
			result[k] = "unknown"
		}
	}
	return result, nil
}
