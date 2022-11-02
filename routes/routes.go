package routes

import (
	"gostorefront/pkg/request"
	"gostorefront/pkg/response"
	"gostorefront/pkg/router"
	"gostorefront/routes/actions"
	"log"
)

func Handle(r request.Request) response.Response {

	routes := map[string]interface{}{
		"GET /":                  actions.GetIndex,
		"GET /list/:slug":        actions.GetCategory,
		"GET /categories":        actions.GetCategory,
		"GET /category/:slug":    actions.GetCategory,
		"GET /product/:uuid":     actions.GetIndex,
		"GET /checkout":          actions.GetCategory,
		"GET /checkout/delivery": actions.GetList,
		"GET /checkout/payment":  actions.GetList,
		"GET /checkout/summary":  actions.GetList,
		"GET /checkout/success":  actions.GetList,
		"GET /checkout/error":    actions.GetList,
		"GET /account":           actions.GetCategory,
		"GET /account/login":     actions.GetCategory,
		"GET /account/register":  actions.GetList,
		"GET /account/logout":    actions.GetList,
		"GET /account/wishlist":  actions.GetList,
	}

	path, _ := r.RequestPath()
	res, err := router.Router(path, routes)

	if err != nil {
		log.Panic(err)
	}

	return res
}
