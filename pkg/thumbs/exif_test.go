package thumbs

import (
	"os"
	"path"
	"testing"
)

func TestExtractExif(t *testing.T) {
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
		want    map[string]string
		wantErr bool
	}{
		{
			name:    "happy path - jpg",
			args:    args{path: path.Join(cwd, "../testfiles/subdir/green.jpg")},
			wantErr: false,
			want: map[string]string{
				"GPSLatitude":  "29 deg 0' 0.00\" N",
				"GPSLongitude": "47 deg 0' 0.00\" E",
			},
		},
		{
			name:    "error path - invalid file",
			args:    args{path: path.Join(cwd, "../testfiles/subdir/nosuch.file")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractExif(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractExif() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for k, v := range tt.want {
				gotV, ok := got[k]
				if !ok {
					t.Errorf("missing exif key: '%s'", k)
				}
				if gotV != v {
					t.Errorf("wrong value '%s' for key '%s', expected value: '%s'", gotV, k, v)
				}
			}
		})
	}
}
