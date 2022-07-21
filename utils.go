package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type minify struct {
	path      string
	directory string
	output    string
}

func (m minify) CreateOutputDir() {
	err := os.Mkdir(m.output, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}

func (m minify) GetFilePath() string {
	filePath := fmt.Sprintf("%s/%s", m.directory, m.path)
	return filePath
}

func (m minify) GetExtension() string {
	filename := strings.Split(m.path, ".")
	return filename[len(filename)-1]
}

func (m minify) ReadFile() string {
	file, err := os.ReadFile(m.GetFilePath())
	if err != nil {
		log.Fatal(err)
	}

	return string(file)
}

func (m minify) FileName() string {
	var filename string

	pathStructure := strings.Split(m.path, ".")
	name := pathStructure[0]
	extension := pathStructure[len(pathStructure)-1]

	if extension == "html" {
		filename = fmt.Sprintf("%s/%s.%s", m.output, name, extension)
	} else {
		filename = fmt.Sprintf("%s/%s.min.%s", m.output, name, extension)
	}

	return filename
}

func (m minify) CSS() {
	var content, filename string
	replacements := map[string]string{
		"\n": "",
		" ":  "",
	}

	content = m.ReadFile()
	filename = m.FileName()

	for index, replacement := range replacements {
		content = strings.ReplaceAll(content, index, replacement)
	}

	os.WriteFile(filename, []byte(content), 0666)
}

func (m minify) HTML() {
	var content, filename string

	content = m.ReadFile()
	filename = m.FileName()

	space := regexp.MustCompile(`(\s{2,}|\n)`)
	content = space.ReplaceAllString(content, "")
	os.WriteFile(filename, []byte(content), 0666)
}
