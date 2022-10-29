package pages

import (
	"gostorefront/connectors"
	"gostorefront/internal/ui"
	"gostorefront/pkg/response"
)

type IndexPageData struct {
	Head  ui.Head
	Title string
}

func Index(args ...string) (response.Response, error) {
	b := connectors.Connector{Page: "category"}
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
