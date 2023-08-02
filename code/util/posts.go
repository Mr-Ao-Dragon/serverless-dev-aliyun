package util

import (
	"fmt"
	"net/http"
)

func GetParamValues(req *http.Request) []string {
	paramValues := make([]string, 0)
	for _, paramName := range req.Form {
		paramValue := req.Form.Get(paramName[0])
		if paramValue != "" {
			paramValues = append(paramValues, fmt.Sprintf("%s is %s", paramName, paramValue))
		}
	}
	return paramValues
}
