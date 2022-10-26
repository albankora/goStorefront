package router

import (
	"gostorefront/pkg/response"
	"gostorefront/router/pages"
)

func Route(URI string) (response.Response, error) {

	routes := map[string]interface{}{
		"/":         pages.Index,
		"/list":     pages.List,
		"/category": pages.Category,
		"404":       pages.NoFound,
	}

	for key, v := range routes {
		if key == URI {
			return v.(func() (response.Response, error))()
		}
	}

	return routes["404"].(func() (response.Response, error))()
}
