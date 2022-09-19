package pages

import (
	"gostorefront/pkg/router"
	"gostorefront/pkg/view"
)

type IndexViewData struct {
	Head  view.Head
	Title string
}

func Index() (router.Response, error) {
	data := IndexViewData{
		Title: "Product Overviews",
		Head: view.Head{
			Lang:            "en",
			PageTitle:       "Product Overviews",
			MetaDescription: "Product Overviews Meta",
		},
	}
	body, err := view.Load("index", data)
	if err != nil {
		return router.EmptyResponse(), nil
	}

	return router.Response{Body: body, StatusCode: 200}, nil
}
