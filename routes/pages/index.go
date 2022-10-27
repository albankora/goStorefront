package pages

import (
	"gostorefront/bridges"
	"gostorefront/internal/ui"
	"gostorefront/pkg/response"
)

type IndexPageData struct {
	Head  ui.Head
	Title string
}

func Index() (response.Response, error) {
	b := bridges.ContentSetup()
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
