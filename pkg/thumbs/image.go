package thumbs

import (
	"bytes"
	"fmt"
	"github.com/nfnt/resize"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

func CreateThumbImage(path string) (*[]byte, error) {
	logrus.Debugf("creating thumbnail image from '%s", path)
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "opening image file")
	}
	defer file.Close()

	logrus.Debug("trying jpeg decoding")
	img, err := jpeg.Decode(file)
	if err == nil {
		logrus.Debug("jpeg decoding succeeded")
		return resizeImage(img)
	}
	logrus.Debugf("jpeg decoding failed with %v", err)
	file.Seek(0, 0)

	logrus.Debug("trying gif decoding")
	img, err = gif.Decode(file)
	if err == nil {
		logrus.Debug("gif decoding succeeded")
		return resizeImage(img)
	}
	logrus.Debugf("gif decoding failed with %v", err)
	file.Seek(0, 0)

	logrus.Debug("trying png decoding")
	img, err = png.Decode(file)
	if err == nil {
		logrus.Debug("png decoding succeeded")
		return resizeImage(img)
	}
	logrus.Debugf("png decoding failed with %v", err)

	return nil, errors.New(fmt.Sprintf("Could not decode image with path '%s'", path))
}

func resizeImage(img image.Image) (*[]byte, error) {
	logrus.Debugf("creating thumbnail image")
	resized := resize.Thumbnail(128, 128, img, resize.Lanczos3)
	buffer := bytes.Buffer{}
	err := png.Encode(&buffer, resized)
	if err != nil {
		return nil, errors.Wrap(err, "encoding thumbnail")
	}
	b := buffer.Bytes()
	return &b, nil
}
