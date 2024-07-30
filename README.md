# Text Template

This project is about using go text/template to automate creation of text files using template.

## Parameters

```
--title       Document title
--version     Document version
--template    Path to template file
--loremipsum  Include lorem ipsum text
--pangram     Include pangram
--output      Output filename
```

## Usage

Run: ```go run main.go --title "Text Template with Go" --version 1.0 --template template.md --loremipsum true```

It should generate: template-yyyy-mm-dd.md