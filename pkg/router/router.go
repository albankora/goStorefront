package router

import (
	"gostorefront/pkg/response"
	"log"
	"regexp"
	"strings"
)

func Router(method string, uri string, routes map[string]interface{}) (response.Response, error) {
	requestRoute := method + " " + uri
	dynamicRoutes := make(map[string]interface{})
	staticRoutes := make(map[string]interface{})
	intRegex := "[0-9]*"
	uuidRegex := "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"
	slugRegex := "[a-zA-Z0-9_-]*"

	for key, v := range routes {
		if strings.Contains(key, "/:") {
			key = strings.ReplaceAll(key, ":int", intRegex)
			key = strings.ReplaceAll(key, ":uuid", uuidRegex)
			key = strings.ReplaceAll(key, ":slug", slugRegex)
			dynamicRoutes[key] = v
		} else {
			staticRoutes[key] = v
		}
	}

	for key, v := range staticRoutes {
		if key == requestRoute {
			return v.(func(...string) (response.Response, error))()
		}
	}

	for key, v := range dynamicRoutes {
		pattern, err := regexp.Compile("^" + key + "$")

		if err != nil {
			log.Panic(err)
		}

		if pattern.MatchString(requestRoute) {
			match := pattern.FindStringSubmatch(intRegex)
			match = append(match, pattern.FindStringSubmatch(intRegex)...)
			match = append(match, pattern.FindStringSubmatch(uuidRegex)...)
			match = append(match, pattern.FindStringSubmatch(slugRegex)...)

			return v.(func(...string) (response.Response, error))(match...)
		}
	}

	return response.NonFoundResponse(), nil
}
