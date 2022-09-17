package view

import (
	"bytes"
	"html/template"
	"log"
	"os"
)

func Load(data interface{}, view string, layout string) (string, error) {
	file := "resources/views/" + view + ".html"
	if _, err := os.Stat(file); err != nil {
		log.Fatal(err)
		return "", err
	}

	layoutFile := "resources/layouts/" + layout + ".html"
	if _, err := os.Stat(layoutFile); err != nil {
		log.Fatal(err)
		return "", err
	}

	t, _ := template.ParseFiles(file, layoutFile)

	var html bytes.Buffer
	if err := t.ExecuteTemplate(&html, layout, data); err != nil {
		log.Fatal(err)
		return "", err
	}

	return html.String(), nil
}
