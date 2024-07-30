package main

import (
	"flag"
	"os"
	"strings"
	"text/template"
	"time"
)

type TemplateValue struct {
	Title      string
	UpdateDate string
	Version    string
	LoremIpsum bool
	Pangram    bool
}

func main() {
	var titlePtr = flag.String("title", "{title}", "Document title")
	var versionPtr = flag.String("version", "{default_version}", "Document version")
	var templatePtr = flag.String("template", "{default_template}", "Path to template file")
	var loremIpsumPtr = flag.Bool("loremipsum", false, "Include lorem ipsum text")
	var pangramPtr = flag.Bool("pangram", false, "Include pangram")
	var outputPtr = flag.String("output", "", "Output filename")
	flag.Parse()

	now := time.Now()
	tValue := TemplateValue{
		Title:      *titlePtr,
		UpdateDate: now.Format("2006/01/02"),
		Version:    *versionPtr,
		LoremIpsum: *loremIpsumPtr,
		Pangram:    *pangramPtr,
	}

	var tFile = *templatePtr
	var tFileArray = strings.Split(tFile, ".")
	var outputFile = tFileArray[0] + "-" + now.Format("2006-01-02") + "." + tFileArray[1]
	if len(*outputPtr) > 0 {
		outputFile = *outputPtr
	}
	output, err := os.Create(outputFile)

	textTmpl, err := template.New(tFile).ParseFiles(tFile)
	if err != nil {
		panic(err)
	}
	err = textTmpl.Execute(output, tValue)
	if err != nil {
		panic(err)
	}
}
