package router

func EmptyResponse() Response {
	return Response{Body: "", StatusCode: 404}
}

type Response struct {
	Body       string
	StatusCode int
}

func Route(URI string, routes map[string]interface{}) (Response, error) {
	for key, v := range routes {
		if key == URI {
			return v.(func() (Response, error))()
		}
	}

	return routes["404"].(func() (Response, error))()
}
