package response

type Response struct {
	Body       string
	StatusCode int
	Headers    map[string]string
}

func JsonResponse(body string, code int) Response {
	return Response{Body: body, StatusCode: code, Headers: map[string]string{"Content-Type": "application/json; charset=utf-8"}}
}

func HttpResponse(body string, code int) Response {
	return Response{Body: body, StatusCode: code, Headers: map[string]string{"Content-Type": "text/html; charset=utf-8"}}
}

func InternalServerError() Response {
	return Response{Body: "Internal Server Error", StatusCode: 500, Headers: map[string]string{"Content-Type": "text/html; charset=utf-8"}}
}

func NonFoundResponse() Response {
	return Response{Body: "Non Found", StatusCode: 404, Headers: map[string]string{"Content-Type": "text/html; charset=utf-8"}}
}
