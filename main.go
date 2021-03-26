package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/roman-wb/go_sensor/generator"
)

func main() {
	filename := "templates.xml"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	generator := generator.NewGenerator(file)

	json, err := json.MarshalIndent(generator.Tree, "  ", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON: %s\n", json)
}
