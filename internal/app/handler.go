package app

import (
	"fmt"
	"lambdacommerce/pkg/router"
	"net/url"
)

type Request struct {
	Body   string
	Path   string
	Method string
}

func HandleRequest(request Request) (router.Response, error) {
	fmt.Printf("Body size = %d.\n", len(request.Body))

	fmt.Println("Headers:")

	u, err := url.Parse(request.Path)
	if err != nil {
		return router.EmptyResponse(), err
	}

	fmt.Println(u.Path)

	return router.Route(u.Path), nil
}
