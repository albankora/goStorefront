package pages

import (
	"gostorefront/pkg/response"
	"gostorefront/pkg/ui"
)

type CategoryViewData struct {
	Title string
}

func Category() (response.Response, error) {
	data := CategoryViewData{Title: "Category"}
	body, err := ui.Load("category", data)

	if err != nil {
		return response.EmptyResponse(), err
	}

	return response.Response{Body: body, StatusCode: 200}, err
}
