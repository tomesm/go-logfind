package logfind

import (
	"fmt"
	"strings"
)

func fullPath(path string) string {
	if !strings.HasSuffix(path, "/") {
		return path + "/"
	}
	return path
}

// Counts a match for every searched string.
func findMatch(str string, substrs []string) int {
	matches := 0
	for _, sub := range substrs {
		// strings.Contains() returns true for empty substring. We don't want that.
		if sub == "" {
			break
		}
		if strings.Contains(strings.ToLower(str), strings.ToLower(sub)) {
			matches++
		}
	}
	return matches
}

func printLine(fname string, lnum int, line string) {
	fmt.Printf("%s : %d : %s\n", fname, lnum, line)
}
