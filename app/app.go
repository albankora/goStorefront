package app

import (
	"fmt"

	"gostorefront/pkg/response"
	"gostorefront/router"
	"net/url"
)

type Request struct {
	Body   string
	Path   string
	Method string
}

func HandleRequest(request Request) (response.Response, error) {
	u, err := url.Parse(request.Path)
	if err != nil {
		return response.EmptyResponse(), err
	}

	fmt.Println(u.Path)

	response, _ := router.Route(u.Path)

	return response, nil
}
