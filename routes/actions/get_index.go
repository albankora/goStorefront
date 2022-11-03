package actions

import (
	"gostorefront/connectors"
	"gostorefront/pkg/response"
	"gostorefront/template"
)

type GetIndexPageData struct {
	Head  template.Head
	Title string
}

func GetIndex() (response.Response, error) {
	b := connectors.Connector{Page: "category"}
	b.HomePage() // todo implement homepage
	data := GetIndexPageData{
		Title: "Product Overviews",
		Head: template.Head{
			Lang:            "en",
			PageTitle:       "Product Overviews",
			MetaDescription: "Product Overviews Meta",
		},
	}

	body, err := template.Load("homepage", data)

	if err != nil {
		return response.InternalServerError(), err
	}

	return response.HttpResponse(body, 200), nil
}
