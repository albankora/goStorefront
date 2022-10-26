package pages

import (
	"gostorefront/bridge"
	"gostorefront/pkg/response"
	"gostorefront/pkg/ui"
)

func Index() (response.Response, error) {
	b := bridge.Setup()
	data := b.Index()
	body, err := ui.Load("index", data)
	if err != nil {
		return response.EmptyResponse(), err
	}

	return response.Response{Body: body, StatusCode: 200}, nil
}
