package request

import (
	"net/url"
)

type Request struct {
	Body   string
	Path   string
	Method string
}

func (r Request) RequestPath() (string, error) {
	u, err := url.Parse(r.Path)

	if err != nil {
		return "", err
	}

	return r.Method + " " + u.Path, nil
}
