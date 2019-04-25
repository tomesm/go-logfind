package main

import (
	"flag"

	"github.com/tomesm/logfind/cmd"
)

func main() {

	dirname := flag.String("dirname", "/var/log", "Name of dir to search")
	text := flag.String("text", "ERROR", "Searched text")

	dir := cmd.NewDir(*dirname, *text)
	dir.Search()
}
