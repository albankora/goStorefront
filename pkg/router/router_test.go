package router

import (
	"gostorefront/pkg/response"
	"testing"
)

func foo() (response.Response, error) {
	return response.Response{Body: "foo", StatusCode: 200}, nil
}

func baz() (response.Response, error) {
	return response.Response{Body: "baz", StatusCode: 200}, nil
}

func postBaz() (response.Response, error) {
	return response.Response{Body: "postBaz", StatusCode: 200}, nil
}

func TestStaticRoutes(t *testing.T) {
	routes := map[string]interface{}{
		"GET /":         foo,
		"GET /contact":  baz,
		"POST /contact": postBaz,
	}

	res, _ := Router("GET", "/", routes)
	if res.Body != "foo" {
		t.Fatalf(`Route want match for %q`, "GET /")
	}

	res, _ = Router("GET", "/contact", routes)
	if res.Body != "baz" {
		t.Fatalf(`Route want match for %q`, "GET /contact")
	}

	res, _ = Router("POST", "/contact", routes)
	if res.Body != "postBaz" {
		t.Fatalf(`Route want match for %q`, "POST /contact")
	}
}

func TestDynamicIntegerRoutes(t *testing.T) {
	routes := map[string]interface{}{
		"GET /contact/:int":  baz,
		"POST /contact/:int": postBaz,
	}

	res, _ := Router("GET", "/contact/123", routes)
	if res.Body != "baz" {
		t.Fatalf(`Route want match for %q`, "GET /contact/:int")
	}

	res, _ = Router("POST", "/contact/123", routes)
	if res.Body != "postBaz" {
		t.Fatalf(`Route want match for %q`, "POST /contact/:int")
	}
}

func TestDynamicUUIDRoutes(t *testing.T) {
	routes := map[string]interface{}{
		"GET /contact/:uuid":  baz,
		"POST /contact/:uuid": postBaz,
	}

	res, _ := Router("GET", "/contact/51e5fe10-5325-4a32-bce8-7ebe9708c453", routes)
	if res.Body != "baz" {
		t.Fatalf(`Route want match for %q`, "GET /contact/:uuid")
	}

	res, _ = Router("POST", "/contact/51e5fe10-5325-4a32-bce8-7ebe9708c453", routes)
	if res.Body != "postBaz" {
		t.Fatalf(`Route want match for %q`, "POST /contact/:uuid")
	}
}

func TestDynamicStringRoutes(t *testing.T) {
	routes := map[string]interface{}{
		"GET /:slug":          foo,
		"GET /contact/:slug":  baz,
		"POST /contact/:slug": postBaz,
	}

	res, _ := Router("GET", "/this-is-random-route", routes)
	if res.Body != "foo" {
		t.Fatalf(`Route want match for %q`, "GET /:slug")
	}

	res, _ = Router("GET", "/contact/this-is-random-route", routes)
	if res.Body != "baz" {
		t.Fatalf(`Route want match for %q`, "GET /contact/:slug")
	}

	res, _ = Router("POST", "/contact/this-is-random-route", routes)
	if res.Body != "postBaz" {
		t.Fatalf(`Route want match for %q`, "POST /contact/:slug")
	}
}
