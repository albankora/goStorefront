package pages

import (
	"gostorefront/bridge"
	"gostorefront/pkg/response"
	"gostorefront/pkg/ui"
)

type CategoryPageData struct {
	Title string
}

func Category() (response.Response, error) {

	b := bridge.CommerceSetup()
	b.Category()
	data := CategoryPageData{Title: "Category"}
	body, err := ui.Load("category", data)

	if err != nil {
		return response.EmptyResponse(), err
	}

	return response.Response{Body: body, StatusCode: 200}, err
}
