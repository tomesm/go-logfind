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
	all := flag.Bool("all", false, "Determine if search should match all strings")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	logfind := logfind.New(*dirname, flag.Args(), *all)
	logfind.Search()
}

func myUsage() {
	fmt.Println("Usage: logfind [OPTIONS] [TEXT]")
	flag.PrintDefaults()
}
