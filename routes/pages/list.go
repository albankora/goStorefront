package pages

import (
	"gostorefront/internal/ui"
	"gostorefront/pkg/response"
)

type ListViewData struct {
	Title string
}

func List() (response.Response, error) {
	data := ListViewData{Title: "Product list"}
	body, err := ui.Load("list", data)

	if err != nil {
		return response.InternalServerError(), err
	}

	return response.Response{Body: body, StatusCode: 200}, nil
}
