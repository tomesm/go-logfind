package logfind

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Finder  is a struct holding matchAll necessary information about searched Finderectory
type Finder struct {
	dirName  string
	texts    []string
	fileType string
	matchAll bool
	format   bool
}

// Result represents formatted JSON output of the search
type Result struct {
	Directory string
	Files     []LogFile
}

// LogFile holds log records from particular file
type LogFile struct {
	File  string `json:"file"`
	Lines []Line `json:"lines"`
}

// Line represents a found line in JSON
type Line struct {
	Line int    `json:"line"`
	Log  string `json:"log"`
}

// New creates a new Finder struct to perform search
func New(dirName string, texts []string, matchAll bool, fileType string, format bool) *Finder {
	return &Finder{
		dirName:  fullPath(dirName),
		texts:    texts,
		fileType: fileType,
		matchAll: matchAll,
		format:   format,
	}
}

// Search traverses .log files in a directory and searches for match
func (f *Finder) Search() {
	files, err := ioutil.ReadDir(f.dirName)
	if err != nil {
		log.Fatal(err)
	}
	var logs []LogFile
	for _, file := range files {
		if strings.HasSuffix(file.Name(), f.fileType) {
			logFile := *f.searchFile(file.Name())
			if logFile.Lines != nil {
				logs = append(logs, logFile)
			}
		}
	}
	if len(logs) > 0 && f.format {
		f.printJSON(logs)
	}
}

// Searches given file
func (f *Finder) searchFile(fname string) *LogFile {
	file, err := os.Open(f.dirName + fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []Line
	f.scanFile(scanner, &lines, fname)
	return &LogFile{
		File:  file.Name(),
		Lines: lines,
	}
}

func (f *Finder) scanFile(scanner *bufio.Scanner, lines *[]Line, fname string) {
	lnum := 1
	for scanner.Scan() {
		matches := findMatch(scanner.Text(), f.texts)
		if f.matchAll && matches == len(f.texts) {
			// Output in JSON format else print line by line
			if f.format {
				*lines = append(*lines, *line(lnum, scanner.Text()))
			} else {
				printLine(fname, lnum, scanner.Text())
			}
			break
		} else if !f.matchAll && matches > 0 {
			// Output in JSON format else print line by line
			if f.format {
				*lines = append(*lines, *line(lnum, scanner.Text()))
			} else {
				printLine(fname, lnum, scanner.Text())
			}
		}
		lnum++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// print directory in JSON format
func (f *Finder) printJSON(logs []LogFile) {
	result := Result{
		Directory: f.dirName,
		Files:     logs,
	}

	var jsonData []byte

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return
	}
	fmt.Printf("%s\n", string(jsonData))
}

func line(lnum int, log string) *Line {
	return &Line{
		Line: lnum,
		Log:  log,
	}
}

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
	fmt.Printf("%s:%d:%s\n", fname, lnum, line)
}
