package router

import (
	"gostorefront/pkg/response"
)

func Router(uri string, routes map[string]interface{}) (response.Response, error) {

	for key, v := range routes {
		if key == uri {
			return v.(func() (response.Response, error))()
		}
	}

	return response.NonFoundResponse(), nil
}
