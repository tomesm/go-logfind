package finder

import (
	"testing"
)

func TestFullPath(t *testing.T) {
	tests := []struct {
		name    string
		dirName string
		want    string
	}{
		{
			name:    "Relative path",
			dirName: "log",
			want:    "log/",
		},
		{
			name:    "Absolute path",
			dirName: "/var/log",
			want:    "/var/log/",
		},
		{
			name:    "Full absolute path",
			dirName: "/var/log/",
			want:    "/var/log/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fullPath(tt.dirName)
			if path != tt.want {
				t.Errorf("want %q; got %q", tt.want, path)
			}
		})
	}
}
