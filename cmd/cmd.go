package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

// Dir  is a struct holding all necessary information about searched directory
type Dir struct {
	dirName string
	text    string
}

// NewDir creates a new Dir struct to perform search
func NewDir(name string, text string) *Dir {
	if !bytes.HasSuffix([]byte(name), []byte("/")) {
		name = name + "/"
	}
	return &Dir{
		dirName: name,
		text:    text,
	}
}

func (d *Dir) searchFile(fname string, wg *sync.WaitGroup) {
	f, err := os.Open(d.dirName + fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	defer wg.Done()

	s := bufio.NewScanner(f)
	lnum := 1

	for s.Scan() {
		if strings.Contains(s.Text(), d.text) {
			fmt.Printf("%s : %d : %s\n", fname, lnum, s.Text())
		}
		lnum++
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

// Search traverses all .log files in a directory and searches for match
func (d *Dir) Search() {
	files, err := ioutil.ReadDir(d.dirName)
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".log") {
			wg.Add(1)
			go d.searchFile(file.Name(), &wg)
		}
	}
	wg.Wait()
}
