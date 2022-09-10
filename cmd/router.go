package main

import "fmt"

func index(p string) string {
	fmt.Println("function f parameter:", p)

	return p
}

func list(p string) string {
	fmt.Println("function g parameters:", p)

	return p
}

func nonFound() string {
	return "non found"
}

func router(path string) string {
	routes := map[string]interface{}{
		"/":     index,
		"/list": list,
	}

	for key, v := range routes {
		if key == path {
			return v.(func(string) string)("astringe")
		}
	}

	return nonFound()
}
