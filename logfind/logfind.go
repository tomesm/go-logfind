package logfind

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

// Finder  is a struct holding all necessary information about searched Finderectory
type Finder struct {
	dirName string
	texts   []string
	all     bool
}

// New creates a new Finder struct to perform search
func New(dirName string, texts []string, all bool) *Finder {
	return &Finder{
		dirName: fullPath(dirName),
		texts:   texts,
		all:     all,
	}
}

func (f *Finder) searchFile(fname string, wg *sync.WaitGroup) {
	file, err := os.Open(f.dirName + fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer wg.Done()

	s := bufio.NewScanner(file)
	lnum := 1

	for s.Scan() {
		matches := findMatch(s.Text(), f.texts)

		if f.all && matches == len(f.texts) {
			printLine(fname, lnum, s.Text())
			break
		} else if !f.all && matches > 0 {
			printLine(fname, lnum, s.Text())
		}
		lnum++
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

// Search traverses all .log files in a Finderectory and searches for match
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
