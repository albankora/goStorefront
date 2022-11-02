package actions

import (
	"gostorefront/connectors"
	"gostorefront/pkg/response"
	"gostorefront/template"
)

type GetCategoryPageData struct {
	Head  template.Head
	Title string
}

func GetCategory(args ...string) (response.Response, error) {

	b := connectors.Connector{Page: "category"}
	b.Category()
	data := GetCategoryPageData{Title: "Category"}
	body, err := template.Load("category", data)

	if err != nil {
		return response.InternalServerError(), err
	}

	return response.Response{Body: body, StatusCode: 200}, err
}
