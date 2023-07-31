package main

import (
	"context"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func main() {
	fc.StartHttp(HandleHttpRequest)
}

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	_, err := w.Write([]byte("hello there\n"))
	if err != nil {
		return err
	}
	return nil
}
