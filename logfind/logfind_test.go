package logfind

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

func TestMatch(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		substrs []string
		want    int
	}{
		{
			name:    "Match all",
			str:     "buildInstallPlanReturningError:]: file://",
			substrs: []string{"Error", "plan", "FILE"},
			want:    3,
		},
		{
			name:    "Match some",
			str:     "buildInstallPlanReturningError:]: file://",
			substrs: []string{"ERROR", "File"},
			want:    2,
		},
		{
			name:    "No Match",
			str:     "buildInstallPlanReturningError:]: file://",
			substrs: []string{"System", "WIN"},
			want:    0,
		},
		{
			name:    "Empty",
			str:     "buildInstallPlanReturningError:]: file://",
			substrs: []string{"", "", ""},
			want:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match := findMatch(tt.str, tt.substrs)
			if match != tt.want {
				t.Errorf("want %d; got %d", tt.want, match)
			}
		})
	}
}
