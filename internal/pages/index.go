package pages

import (
	"lambdastorefront/pkg/router"
	"lambdastorefront/pkg/view"
)

type IndexViewData struct {
	Name     string
	Greeting string
}

func Index() (router.Response, error) {
	data := IndexViewData{Name: "Astaxie", Greeting: "Hello"}
	body, err := view.Load(data, "index", "base")
	if err != nil {
		return router.EmptyResponse(), nil
	}

	return router.Response{Body: body, StatusCode: 200}, nil
}
