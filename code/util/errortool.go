package util

import "net/http"

func EndErrorProcess(err error, w http.ResponseWriter) error {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	} else { // 返回 HTTP 响应。
		w.WriteHeader(http.StatusOK)
		return nil
	}
}
