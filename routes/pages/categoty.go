package pages

import (
	"gostorefront/bridges"
	"gostorefront/internal/ui"
	"gostorefront/pkg/response"
)

type CategoryPageData struct {
	Title string
}

func Category() (response.Response, error) {

	b := bridges.CommerceSetup()
	b.Category()
	data := CategoryPageData{Title: "Category"}
	body, err := ui.Load("category", data)

	if err != nil {
		return response.InternalServerError(), err
	}

	return response.Response{Body: body, StatusCode: 200}, err
}
