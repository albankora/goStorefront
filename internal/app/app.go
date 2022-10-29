package app

import (
	"fmt"

	"gostorefront/pkg/response"
	"gostorefront/routes"
	"net/url"
)

type Request struct {
	Body   string
	Path   string
	Method string
}

func HandleRequest(request Request) response.Response {
	u, err := url.Parse(request.Path)
	if err != nil {
		return response.InternalServerError()
	}

	fmt.Println(u.Path)

	response := routes.Handle(request.Method, u.Path)

	return response
}
