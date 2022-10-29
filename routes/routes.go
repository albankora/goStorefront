package routes

import (
	"gostorefront/pkg/response"
	"gostorefront/pkg/router"
	"gostorefront/routes/pages"
	"log"
)

func Handle(method string, URI string) response.Response {

	routes := map[string]interface{}{
		"GET /":                  pages.Index,
		"GET /list/:slug":        pages.List,
		"GET /categories":        pages.Category,
		"GET /category/:slug":    pages.Category,
		"GET /product/:uuid":     pages.Index,
		"GET /checkout":          pages.Category,
		"GET /checkout/delivery": pages.List,
		"GET /checkout/payment":  pages.List,
		"GET /checkout/summary":  pages.List,
		"GET /checkout/success":  pages.List,
		"GET /checkout/error":    pages.List,
		"GET /account":           pages.Category,
		"GET /account/login":     pages.Category,
		"GET /account/register":  pages.List,
		"GET /account/logout":    pages.List,
		"GET /account/wishlist":  pages.List,
	}

	res, err := router.Router(method, URI, routes)

	if err != nil {
		log.Panic(err)
	}

	return res
}
