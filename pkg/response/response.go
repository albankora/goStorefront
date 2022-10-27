package response

type Response struct {
	Body       string
	StatusCode int
}

func InternalServerError() Response {
	return Response{Body: "Internal Server Error", StatusCode: 500}
}

func NonFoundResponse() Response {
	return Response{Body: "Non Found", StatusCode: 404}
}
