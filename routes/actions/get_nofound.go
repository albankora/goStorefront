package actions

import "gostorefront/pkg/response"

func GetNoFound() (response.Response, error) {
	return response.Response{Body: "non found", StatusCode: 404}, nil
}
