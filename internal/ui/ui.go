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
	theme := "base"
	file := "ui/" + theme + "/pages/" + view + ".html"
	layout := "ui/" + theme + "/layouts/base.html"
	header := "ui/" + theme + "/partials/header.html"
	footer := "ui/" + theme + "//partials/footer.html"
	head := "ui/" + theme + "/partials/head.html"

	t, _ := template.ParseFiles(file, layout, header, footer, head)

	var html bytes.Buffer
	if err := t.ExecuteTemplate(&html, "base", data); err != nil {
		log.Panic(err)
		return "", err
	}

	return html.String(), nil
}
