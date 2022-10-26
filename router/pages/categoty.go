package pages

import (
	"gostorefront/bridge"
	"gostorefront/pkg/response"
	"gostorefront/pkg/ui"
)

func Category() (response.Response, error) {

	b := bridge.Setup()
	data := b.Category()
	body, err := ui.Load("category", data)

	if err != nil {
		return response.EmptyResponse(), err
	}

	return response.Response{Body: body, StatusCode: 200}, err
}
