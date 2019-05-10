package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tomesm/go-logfind/logfind"
)

func main() {
	flag.Usage = myUsage
	dirname := flag.String("dirname", "/var/log", "Name of dir to search")
	matchAll := flag.Bool("match-all", false, "Determine if search should match all searched strings")
	fileType := flag.String("file-type", ".log", "Type/suffix of files to be searched")
	format := flag.Bool("format", false, "Print output in JSON format")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	logfind := logfind.New(*dirname, flag.Args(), *matchAll, *fileType, *format)
	logfind.Search()
}

func myUsage() {
	fmt.Println("Usage: logfind [OPTIONS] [STRINGS]")
	flag.PrintDefaults()
}
