package pages

import (
	"gostorefront/pkg/response"
	"gostorefront/pkg/ui"
)

type IndexViewData struct {
	Head  ui.Head
	Title string
}

func Index() (response.Response, error) {
	data := IndexViewData{
		Title: "Product Overviews",
		Head: ui.Head{
			Lang:            "en",
			PageTitle:       "Product Overviews",
			MetaDescription: "Product Overviews Meta",
		},
	}
	body, err := ui.Load("index", data)
	if err != nil {
		return response.EmptyResponse(), err
	}

	return response.Response{Body: body, StatusCode: 200}, nil
}
