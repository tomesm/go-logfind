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

func TestMatchStrings(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		substr string
		want   bool
	}{
		{
			name:   "Is a match: Camel Case",
			str:    "buildInstallPlanReturningError:]: file://",
			substr: "Error",
			want:   true,
		},
		{
			name:   "Is a match: Upper Case",
			str:    "buildInstallPlanReturningError:]: file://",
			substr: "ERROR",
			want:   true,
		},
		{
			name:   "Is a match: Lower Case",
			str:    "buildInstallPlanReturningError:]: file://",
			substr: "error",
			want:   true,
		},
		{
			name:   "Not a match: Lower Case",
			str:    "buildInstallPlanReturningError:]: file://",
			substr: "login",
			want:   false,
		},
		{
			name:   "Not a match: Upper Case",
			str:    "buildInstallPlanReturningError:]: file://",
			substr: "LOGIN",
			want:   false,
		},
		{
			name:   "Not a match: Camel Case",
			str:    "buildInstallPlanReturningError:]: file://",
			substr: "Login",
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match := isMatch(tt.str, tt.substr)
			if match != tt.want {
				t.Errorf("want %t; got %t", tt.want, match)
			}
		})
	}
}
