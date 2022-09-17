package pages

import (
	"fmt"
	"lambdacommerce/pkg/router"
)

func List() (router.Response, error) {
	fmt.Println("function g parameters:")

	return router.Response{Body: "function g parameters:", StatusCode: 200}, nil
}
