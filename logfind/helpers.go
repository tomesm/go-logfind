package logfind

import "strings"

func fullPath(path string) string {
	if !strings.HasSuffix(path, "/") {
		return path + "/"
	}
	return path
}

func isMatch(str string, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}
