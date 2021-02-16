package data

import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GreenThumbData() *[]byte {
	return &[]byte{0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x80, 0x10, 0x02, 0x00, 0x00, 0x00, 0x1c, 0xcc, 0x2a, 0xdf, 0x00, 0x00, 0x01, 0x78, 0x49, 0x44, 0x41, 0x54, 0x78, 0x9c, 0xec, 0xd3, 0x31, 0x0d, 0x80, 0x50, 0x00, 0xc5, 0xc0, 0x17, 0x82, 0x18, 0x9c, 0xe1, 0xdf, 0x00, 0x41, 0xc6, 0x1f, 0x7a, 0xa7, 0xa0, 0x4b, 0xef, 0x6d, 0xef, 0xf7, 0x0c, 0x92, 0xae, 0xd3, 0x01, 0x70, 0x92, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0xf6, 0x07, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x84, 0x02, 0x90, 0xb9, 0xab, 0xbd, 0x4f, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82}
	//cwd, _ := os.Getwd()
	//file, _ := os.Open(filepath.Join(cwd,"../testfiles/subdir/green.jpg"))
	//img, _ := jpeg.Decode(file)
	//resized := resize.Thumbnail(128, 128, img, resize.Lanczos3)
	//buffer := bytes.Buffer{}
	//_ = png.Encode(&buffer, resized)
	//data := buffer.Bytes()
	//return &data
}

func GreenData() *[]byte {
	cwd, _ := os.Getwd()
	data, _ := ioutil.ReadFile(filepath.Join(cwd, "../testfiles/subdir/green.jpg"))
	return &data
}

func GreenExif() *exif.Exif {
	exif.RegisterParsers(mknote.All...)
	cwd, _ := os.Getwd()
	file, _ := os.Open(filepath.Join(cwd, "../testfiles/subdir/green.jpg"))
	exifData, _ := exif.Decode(file)
	fmt.Printf("HIER: %v", exifData.String())
	return exifData
}

func BlueThumbData() *[]byte {
	return &[]byte{0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x80, 0x10, 0x02, 0x00, 0x00, 0x00, 0x1c, 0xcc, 0x2a, 0xdf, 0x00, 0x00, 0x01, 0x74, 0x49, 0x44, 0x41, 0x54, 0x78, 0x9c, 0xec, 0xd3, 0x31, 0x11, 0xc0, 0x30, 0x00, 0xc4, 0xb0, 0xbf, 0x5e, 0xf9, 0x53, 0x4e, 0x60, 0x64, 0xb0, 0x84, 0xc0, 0x8b, 0xff, 0x6d, 0x3b, 0x67, 0x90, 0xf4, 0xbd, 0x0e, 0x80, 0x97, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0x34, 0x03, 0x90, 0x66, 0x00, 0xd2, 0x0c, 0x40, 0x9a, 0x01, 0x48, 0x33, 0x00, 0x69, 0x06, 0x20, 0xcd, 0x00, 0xa4, 0x19, 0x80, 0xb4, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x87, 0x03, 0x01, 0xe8, 0xe4, 0x7f, 0xb6, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82}
	//cwd, _ := os.Getwd()
	//file, _ := os.Open(filepath.Join(cwd,"../testfiles/blue.gif"))
	//img, _ := gif.Decode(file)
	//resized := resize.Thumbnail(128, 128, img, resize.Lanczos3)
	//buffer := bytes.Buffer{}
	//_ = png.Encode(&buffer, resized)
	//data := buffer.Bytes()
	//return &data
}

func GrayThumbData() *[]byte {
	return &[]byte{0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x80, 0x08, 0x02, 0x00, 0x00, 0x00, 0x4c, 0x5c, 0xf6, 0x9c, 0x00, 0x00, 0x01, 0x33, 0x49, 0x44, 0x41, 0x54, 0x78, 0x9c, 0xec, 0xd1, 0x51, 0x09, 0xc0, 0x30, 0x00, 0xc5, 0xc0, 0x31, 0x2a, 0xfc, 0x49, 0xaf, 0x8c, 0xfb, 0x68, 0x4e, 0x41, 0x20, 0x67, 0xdb, 0x17, 0xe7, 0xd7, 0x01, 0xaf, 0x6b, 0x00, 0xd6, 0x00, 0xac, 0x01, 0x58, 0x03, 0xb0, 0x06, 0x60, 0x0d, 0xc0, 0x1a, 0x80, 0x35, 0x00, 0x6b, 0x00, 0xd6, 0x00, 0xac, 0x01, 0x58, 0x03, 0xb0, 0x06, 0x60, 0x0d, 0xc0, 0x1a, 0x80, 0x35, 0x00, 0x6b, 0x00, 0xd6, 0x00, 0xac, 0x01, 0x58, 0x03, 0xb0, 0x06, 0x60, 0x0d, 0xc0, 0x1a, 0x80, 0x35, 0x00, 0x6b, 0x00, 0xd6, 0x00, 0xac, 0x01, 0x58, 0x03, 0xb0, 0x06, 0x60, 0x0d, 0xc0, 0x1a, 0x80, 0x35, 0x00, 0x6b, 0x00, 0xd6, 0x00, 0xac, 0x01, 0x58, 0x03, 0xb0, 0x06, 0x60, 0x0d, 0xc0, 0x1a, 0x80, 0x35, 0x00, 0x6b, 0x00, 0xd6, 0x00, 0xac, 0x01, 0x58, 0x03, 0xb0, 0x06, 0x60, 0x0d, 0xc0, 0x1a, 0x80, 0x35, 0x00, 0x6b, 0x00, 0xd6, 0x00, 0xac, 0x01, 0x58, 0x03, 0xb0, 0x06, 0x60, 0x0d, 0xc0, 0x1a, 0x80, 0x35, 0x00, 0x6b, 0x00, 0xd6, 0x00, 0xac, 0x01, 0x58, 0x03, 0xb0, 0x06, 0x60, 0x0d, 0xc0, 0x1a, 0x80, 0x35, 0x00, 0x6b, 0x00, 0xd6, 0x00, 0xac, 0x01, 0x58, 0x03, 0xb0, 0x06, 0x60, 0x0d, 0xc0, 0x1a, 0x80, 0x35, 0x00, 0x6b, 0x00, 0xd6, 0x00, 0xac, 0x01, 0x58, 0x03, 0xb0, 0x06, 0x60, 0x0d, 0xc0, 0x1a, 0x80, 0x35, 0x00, 0x6b, 0x00, 0xd6, 0x00, 0xac, 0x01, 0x58, 0x03, 0xb0, 0x06, 0x60, 0x0d, 0xc0, 0x1a, 0x80, 0x35, 0x00, 0x6b, 0x00, 0xd6, 0x00, 0xac, 0x01, 0x58, 0x03, 0xb0, 0x06, 0x60, 0x0d, 0xc0, 0x1a, 0x80, 0x35, 0x00, 0x6b, 0x00, 0xd6, 0x00, 0xac, 0x01, 0x58, 0x03, 0xb0, 0x06, 0x60, 0x0d, 0xc0, 0x1a, 0x80, 0x35, 0x00, 0x6b, 0x00, 0xd6, 0x00, 0xac, 0x01, 0x58, 0x03, 0xb0, 0x06, 0x60, 0x0d, 0xc0, 0x1a, 0x80, 0x35, 0x00, 0x6b, 0x00, 0xd6, 0x00, 0xac, 0x01, 0x58, 0x03, 0xb0, 0x06, 0x60, 0x0d, 0xc0, 0x1a, 0x80, 0x35, 0x00, 0x6b, 0x00, 0xd6, 0x00, 0xac, 0x01, 0x58, 0x03, 0xb0, 0x06, 0x60, 0x0d, 0xc0, 0x6e, 0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0xbc, 0x02, 0x83, 0xce, 0x25, 0x02, 0xef, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82}
	//cwd, _ := os.Getwd()
	//file, _ := os.Open(filepath.Join(cwd,"../testfiles/gray.png"))
	//img, _ := png.Decode(file)
	//resized := resize.Thumbnail(128, 128, img, resize.Lanczos3)
	//buffer := bytes.Buffer{}
	//_ = png.Encode(&buffer, resized)
	//data := buffer.Bytes()
	//return &data
}
