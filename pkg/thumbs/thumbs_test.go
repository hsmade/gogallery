package thumbs

import (
	"github.com/hsmade/gogallery/pkg/testfiles/data"
	"os"
	"path"
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    Thumbs
		wantErr bool
	}{
		{
			name:    "happy path without exif",
			args:    args{path: path.Join(cwd, "../testfiles")},
			wantErr: false,
			want: Thumbs{
				Path: path.Join(cwd, "../testfiles"),
				Files: []thumbFile{
					{
						Name:  "blue.gif",
						Exif:  nil,
						Image: data.BlueThumbData(),
					},
					{
						Name:  "gray.png",
						Exif:  nil,
						Image: data.GrayThumbData(),
					},
				},
			},
		},
		{
			name:    "happy path with exif",
			args:    args{path: path.Join(cwd, "../testfiles/subdir")},
			wantErr: false,
			want: Thumbs{
				Path: path.Join(cwd, "../testfiles/subdir"),
				Files: []thumbFile{
					{
						Name:  "green.jpg",
						Exif:  data.GreenExif(),
						Image: data.GreenThumbData(),
					},
				},
			},
		},
		{
			name:    "happy path with invalid file",
			args:    args{path: path.Join(cwd, "../testfiles/broken")},
			wantErr: false,
			want: Thumbs{
				Path:  path.Join(cwd, "../testfiles/broken"),
				Files: []thumbFile{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.path)
			t.Logf("%s", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.want.Equal(got) {
				t.Errorf("Create() got = %v\nwant %v", got, tt.want)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *Thumbs
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{path: path.Join(cwd, "../testfiles/thumbdir")},
			want: &Thumbs{
				Path: path.Join(cwd, "../testfiles/thumbdir"),
				Files: []thumbFile{
					{
						Name:  "green.jpg",
						Exif:  data.GreenExif(),
						Image: data.GreenData(),
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Load(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThumbs_addImage(t1 *testing.T) {
	type fields struct {
		Path  string
		Files []thumbFile
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Thumbs{
				Path:  tt.fields.Path,
				Files: tt.fields.Files,
			}
			if err := t.addImage(tt.args.path); (err != nil) != tt.wantErr {
				t1.Errorf("addImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestThumbs_save(t1 *testing.T) {
	type fields struct {
		Path  string
		Files []thumbFile
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Thumbs{
				Path:  tt.fields.Path,
				Files: tt.fields.Files,
			}
			if err := t.save(); (err != nil) != tt.wantErr {
				t1.Errorf("save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
