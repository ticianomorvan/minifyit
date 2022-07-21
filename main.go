package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var m minify
	directory := flag.String("d", ".", "Set the directory where are the files to minify")
	output := flag.String("o", "dist", "Set the output directory")
	flag.Parse()

	files, err := os.ReadDir(*directory)
	if err != nil {
		log.Fatal(err)
	}

	m = minify{output: *output, directory: *directory}
	m.CreateOutputDir()

	for _, file := range files {
		m = minify{path: file.Name(), output: *output, directory: *directory}

		switch m.GetExtension() {
		case "css":
			m.CSS()
		case "html":
			m.HTML()
		default:
			return
		}
	}
}
