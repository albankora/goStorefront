package actions

import (
	"gostorefront/pkg/response"
	"gostorefront/template"
)

type GetListData struct {
	Title string
}

func GetList() (response.Response, error) {
	data := GetListData{Title: "Product list"}
	body, err := template.Load("list", data)

	if err != nil {
		return response.InternalServerError(), err
	}

	return response.Response{Body: body, StatusCode: 200}, nil
}
