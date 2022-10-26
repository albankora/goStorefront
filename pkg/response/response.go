package response

func EmptyResponse() Response {
	return Response{Body: "", StatusCode: 404}
}

type Response struct {
	Body       string
	StatusCode int
}
