package main

import (
	"flag"

	"github.com/tomesm/logfind/finder"
)

func main() {

	dirname := flag.String("dirname", "/var/log", "Name of dir to search")
	text := flag.String("text", "ERROR", "Searched text")

	finder := finder.New(*dirname, *text)
	finder.Search()
}
