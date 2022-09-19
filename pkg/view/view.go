package view

import (
	"bytes"
	"html/template"
	"log"
)

type Head struct {
	Lang            string
	PageTitle       string
	MetaDescription string
}

func Load(view string, data interface{}) (string, error) {
	file := "resources/views/" + view + ".html"
	layout := "resources/layouts/base.html"
	nav := "resources/partials/nav.html"
	footer := "resources/partials/footer.html"
	head := "resources/partials/head.html"

	t, _ := template.ParseFiles(file, layout, nav, footer, head)

	var html bytes.Buffer
	if err := t.ExecuteTemplate(&html, "base", data); err != nil {
		log.Fatal(err)
		return "", err
	}

	return html.String(), nil
}
