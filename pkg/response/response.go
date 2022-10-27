package response

type Response struct {
	Body       string
	StatusCode int
}

func EmptyResponse() Response {
	return Response{Body: "", StatusCode: 404}
}
