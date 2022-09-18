package pages

import (
	"gostorefront/pkg/router"
	"gostorefront/pkg/view"
)

type CategoryViewData struct {
	Title string
}

func Category() (router.Response, error) {
	data := CategoryViewData{Title: "Category"}
	body, err := view.Load("category", data)
	if err != nil {
		return router.EmptyResponse(), nil
	}

	return router.Response{Body: body, StatusCode: 200}, nil
}
