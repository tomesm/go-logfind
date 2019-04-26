package main

import (
	"flag"

	"github.com/tomesm/go-logfind/logfind"
)

func main() {

	dirname := flag.String("dirname", "/var/log", "Name of dir to search")
	text := flag.String("text", "ERROR", "Searched text")

	logfind := logfind.New(*dirname, *text)
	logfind.Search()
}
