package pages

import "lambdastorefront/pkg/router"

func NoFound() (router.Response, error) {
	return router.Response{Body: "non found", StatusCode: 404}, nil
}
