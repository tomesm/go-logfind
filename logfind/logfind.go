package logfind

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

// Finder  is a struct holding matchAll necessary information about searched Finderectory
type Finder struct {
	dirName  string
	texts    []string
	matchAll bool
}

// Result represents formatted JSON output of the search
type Result struct {
	File  string `json:"file"`
	Lines []Line `json:"lines"`
}

// Line represents a found line in JSON
type Line struct {
	Line int    `json:"line"`
	Log  string `json:"log"`
}

// New creates a new Finder struct to perform search
func New(dirName string, texts []string, matchAll bool) *Finder {
	return &Finder{
		dirName:  fullPath(dirName),
		texts:    texts,
		matchAll: matchAll,
	}
}

// Search traverses .log files in a directory and searches for match
func (f *Finder) Search() {
	files, err := ioutil.ReadDir(f.dirName)
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".log") {
			wg.Add(1)
			go f.searchFile(file.Name(), &wg)
		}
	}
	wg.Wait()
}

// Searches given file
func (f *Finder) searchFile(fname string, wg *sync.WaitGroup) {
	file, err := os.Open(f.dirName + fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer wg.Done()

	s := bufio.NewScanner(file)
	lnum := 1

	var lines []Line

	for s.Scan() {
		matches := findMatch(s.Text(), f.texts)

		if f.matchAll && matches == len(f.texts) {
			//printLine(fname, lnum, s.Text())
			lines = append(lines, *toJSON(lnum, s.Text()))
			break
		} else if !f.matchAll && matches > 0 {
			printLine(fname, lnum, s.Text())
		}
		lnum++
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	// TODO: create a channel to send Lines to one result??
	result := Result{
		File:  fname,
		Lines: lines,
	}
	var jsonData []byte

	jsonData, err = json.Marshal(result)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", string(jsonData))
}

func toJSON(lnum int, log string) *Line {
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
	//fmt.Printf("%s:%d:%s\n", fname, lnum, line)

}
