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
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	logfind := logfind.New(*dirname, flag.Args(), *matchAll)
	logfind.Search()
}

func myUsage() {
	fmt.Println("Usage: logfind [OPTIONS] [TEXT]")
	flag.PrintDefaults()
}
