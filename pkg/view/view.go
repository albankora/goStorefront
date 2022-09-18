package view

import (
	"bytes"
	"html/template"
	"log"
)

func Load(view string, data interface{}) (string, error) {
	file := "resources/views/" + view + ".html"
	layout := "resources/layouts/base.html"
	nav := "resources/partials/nav.html"

	t, _ := template.ParseFiles(file, layout, nav)

	var html bytes.Buffer
	if err := t.ExecuteTemplate(&html, "base", data); err != nil {
		log.Fatal(err)
		return "", err
	}

	return html.String(), nil
}
