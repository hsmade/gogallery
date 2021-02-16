package thumbs

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hsmade/gogallery/pkg/thumbs/protobuf"
	"github.com/pkg/errors"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

func init() {
	exif.RegisterParsers(mknote.All...)
}

type Thumbs struct {
	Path  string
	Files []*thumbFile
}

type thumbFile struct {
	Name  string
	Exif  *exif.Exif
	Image *[]byte
}

func (t *thumbFile) String() string {
	if t.Exif != nil {
		return fmt.Sprintf("Name: '%s', Exif: '%s', image: %x.", t.Name, t.Exif.String(), *t.Image)
	}
	return fmt.Sprintf("Name: '%s', Exif: empty, image: %x.", t.Name, *t.Image)
}

func (t *thumbFile) Equal(b *thumbFile) bool {
	part1 := t.Name == b.Name && bytes.Equal(*t.Image, *b.Image)
	if !part1 {
		return false
	}

	if t.Exif == nil && b.Exif == nil {
		return part1
	}

	if t.Exif == nil || b.Exif == nil {
		return false
	}

	return part1 && bytes.Equal(t.Exif.Raw, b.Exif.Raw)
}

func (t *Thumbs) Equal(b *Thumbs) bool {
	if t.Path != b.Path {
		return false
	}
	for index, thumb := range t.Files {
		if !thumb.Equal(b.Files[index]) {
			return false
		}
	}
	return true
}

// Load creates a new Thumbs object from a path
func Load(path string) (*Thumbs, error) {
	logrus.Debugf("loading thumbs file from path '%s'", path)
	file, err := os.Open(filepath.Join(path, "thumbs.bin"))
	if err != nil {
		return nil, errors.Wrap(err, "opening thumb file")
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "reading thumb file")
	}

	thumbsBuffer := protobuf.Thumbs{}
	err = proto.Unmarshal(data, &thumbsBuffer)
	if err != nil {
		return nil, errors.Wrap(err, "decoding thumb file")
	}

	thumbs := Thumbs{
		Path: path,
	}

	for _, thumbBuffer := range thumbsBuffer.Thumb {
		exifData, err := exif.Decode(file) // FIXME: somehow reproduce the exif data from the raw
		if err != nil {
			return &thumbs, errors.Wrap(err, "decoding file")
		}
		thumb := thumbFile{
			Name:  thumbBuffer.Name,
			Exif:  exifData,
			Image: &thumbBuffer.Thumbnail,
		}
		thumbs.Files = append(thumbs.Files, &thumb)
	}
	return &thumbs, nil
}

// Create scans a path for images and creates a new Thumbs object
func Create(path string) (*Thumbs, error) {
	logrus.Debugf("creating new thumbs object for path '%s'", path)
	thumbs := Thumbs{Path: path}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.Wrap(err, "reading directory")
	}

	for _, fileInfo := range files {
		if fileInfo.IsDir() || fileInfo.Name() == "thumbs.bin" {
			continue
		}
		err = thumbs.addImage(fileInfo.Name())
		if err != nil {
			logrus.Errorf("could not add file: %v", err)
		}
	}

	return &thumbs, thumbs.save()
}

// save stores the Thumbs file
func (t *Thumbs) save() error {
	logrus.Infof("storing thumbs file '%s'", t.Path)
	thumbsBuffer := protobuf.Thumbs{}
	for _, thumb := range t.Files {
		thumbBuffer := protobuf.Thumb{
			Name:      thumb.Name,
			Thumbnail: *thumb.Image,
		}
		if thumb.Exif != nil {
			thumbBuffer.Exifdata = thumb.Exif.Raw
		}

		thumbsBuffer.Thumb = append(thumbsBuffer.Thumb, &thumbBuffer)
	}

	data, err := proto.Marshal(&thumbsBuffer)
	if err != nil {
		return errors.Wrap(err, "encoding thumbs file")
	}

	err = ioutil.WriteFile(filepath.Join(t.Path, "thumbs.bin"), data, 0644)
	if err != nil {
		return errors.Wrap(err, "writing thumbs file")
	}
	return nil
}

func (t *Thumbs) addImage(path string) error {
	logrus.Debugf("adding image with path '%s' to thumbs object for path '%s'", path, t.Path)
	file, err := os.Open(filepath.Join(t.Path, path))
	if err != nil {
		return errors.Wrap(err, "opening file")
	}

	exifData, err := exif.Decode(file)
	if err != nil {
		logrus.Debugf("file has no or invalid exif data: '%s'", filepath.Join(t.Path, path))
	}
	_ = file.Close()

	thumb := thumbFile{
		Name: path,
		Exif: exifData,
	}

	img, err := CreateThumbImage(filepath.Join(t.Path, path))
	if err != nil {
		return errors.Wrap(err, "creating thumbnail image")
	}
	thumb.Image = img

	t.Files = append(t.Files, &thumb)
	return nil
}
