package app

import (
	"fmt"

	"gostorefront/internal/pages"
	"gostorefront/pkg/router"
	"net/url"
)

type Request struct {
	Body   string
	Path   string
	Method string
}

func HandleRequest(request Request) (router.Response, error) {
	u, err := url.Parse(request.Path)
	if err != nil {
		return router.EmptyResponse(), err
	}

	fmt.Println(u.Path)

	routes := map[string]interface{}{
		"/":         pages.Index,
		"/list":     pages.List,
		"/category": pages.Category,
		"404":       pages.NoFound,
	}

	response, _ := router.Route(u.Path, routes)

	return response, nil
}
