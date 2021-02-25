package main

import (
	"log"
	"os"
)

func main() {
	filename := "templates.xml"
	xmlfile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer xmlfile.Close()
}
