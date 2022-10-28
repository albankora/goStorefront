package ui

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
	file := "ui/pages/" + view + ".html"
	layout := "ui/layouts/base.html"
	nav := "ui/partials/nav.html"
	footer := "ui/partials/footer.html"
	head := "ui/partials/head.html"

	t, _ := template.ParseFiles(file, layout, nav, footer, head)

	var html bytes.Buffer
	if err := t.ExecuteTemplate(&html, "base", data); err != nil {
		log.Panic(err)
		return "", err
	}

	return html.String(), nil
}
