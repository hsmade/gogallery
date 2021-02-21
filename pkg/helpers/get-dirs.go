package helpers

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
)

func GetDirs(path string) ([]string, error) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("scanning directory %s", path))
	}

	directories := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			directories = append(directories, entry.Name())
		}
	}
	return directories, nil
}
