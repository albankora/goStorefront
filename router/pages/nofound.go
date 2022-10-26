package pages

import "gostorefront/pkg/response"

func NoFound() (response.Response, error) {
	return response.Response{Body: "non found", StatusCode: 404}, nil
}
