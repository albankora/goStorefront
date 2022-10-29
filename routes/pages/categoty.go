package pages

import (
	"gostorefront/connectors"
	"gostorefront/internal/ui"
	"gostorefront/pkg/response"
)

type CategoryPageData struct {
	Head  ui.Head
	Title string
}

func Category(args ...string) (response.Response, error) {

	b := connectors.Connector{Page: "category"}
	b.Category()
	data := CategoryPageData{Title: "Category"}
	body, err := ui.Load("category", data)

	if err != nil {
		return response.InternalServerError(), err
	}

	return response.Response{Body: body, StatusCode: 200}, err
}
