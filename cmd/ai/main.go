package main

import (
	"flag"
	"log"
)

func main() {
	numAliens := flag.Int("N", 0, "Number of aliens")
	flag.Parse()
	if numAliens == nil {
		log.Panicln("Number of aliens not provided")
	}
}
