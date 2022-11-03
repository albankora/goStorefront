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
		"GET /":               actions.GetIndex,
		"GET /list/:slug":     actions.GetCategory,
		"GET /categories":     actions.GetCategory,
		"GET /category/:slug": actions.GetCategory,
	}

	path, _ := r.RequestPath()
	res, err := router.Router(path, routes)

	if err != nil {
		log.Panic(err)
	}

	return res
}
