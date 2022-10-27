package pages

import (
	"gostorefront/bridge"
	"gostorefront/pkg/response"
	"gostorefront/pkg/ui"
)

type IndexPageData struct {
	Head  ui.Head
	Title string
}

func Index() (response.Response, error) {
	b := bridge.ContentSetup()
	b.HomePage() // todo implement homepage
	data := IndexPageData{
		Title: "Product Overviews",
		Head: ui.Head{
			Lang:            "en",
			PageTitle:       "Product Overviews",
			MetaDescription: "Product Overviews Meta",
		},
	}

	body, err := ui.Load("index", data)

	if err != nil {
		return response.InternalServerError(), err
	}

	return response.Response{Body: body, StatusCode: 200}, nil
}
