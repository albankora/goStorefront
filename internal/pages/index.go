package pages

import (
	"gostorefront/pkg/router"
	"gostorefront/pkg/view"
)

type IndexViewData struct {
	Title string
}

func Index() (router.Response, error) {
	data := IndexViewData{Title: "Product Overviews"}
	body, err := view.Load("index", data)
	if err != nil {
		return router.EmptyResponse(), nil
	}

	return router.Response{Body: body, StatusCode: 200}, nil
}
