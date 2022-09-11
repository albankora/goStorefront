package router

import "fmt"

func index(p string) Response {
	fmt.Println("function f parameter:", p)

	return Response{Body: "function f parameter:", StatusCode: 200}
}

func list(p string) Response {
	fmt.Println("function g parameters:", p)

	return Response{Body: "function g parameters:", StatusCode: 200}
}

func nonFound() Response {
	return Response{Body: "non found", StatusCode: 404}
}

func EmptyResponse() Response {
	return Response{Body: "", StatusCode: 0}
}

type Response struct {
	Body       string
	StatusCode int
}

func Route(URI string) Response {
	routes := map[string]interface{}{
		"/":     index,
		"/list": list,
	}

	for key, v := range routes {
		if key == URI {
			return v.(func(string) Response)("astringe")
		}
	}

	return nonFound()
}
