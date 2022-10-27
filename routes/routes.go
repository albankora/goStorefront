package routes

import (
	"gostorefront/pkg/response"
	"gostorefront/pkg/router"
	"gostorefront/routes/pages"
	"log"
)

func Handle(URI string) response.Response {

	routes := map[string]interface{}{
		"/":         pages.Index,
		"/list":     pages.List,
		"/category": pages.Category,
		"404":       pages.NoFound,
	}

	res, err := router.Router(URI, routes)

	if err != nil {
		log.Fatal(err)
	}

	return res
}
