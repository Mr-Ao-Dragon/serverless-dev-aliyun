// 包 main 是应用程序的入口点。
package main

// 导入必要的库。
import (
	"context"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

// 主函数启动 HTTP 服务器。
func main() {
	fc.StartHttp(HandleHttpRequest)
}

// HandleHttpRequest 是 HTTP 处理程序函数。
func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	// 设置 HTTP 响应状态码为 200 OK。
	w.WriteHeader(http.StatusOK)

	// 设置 Content-Type 头部为 text/plain。
	w.Header().Add("Content-Type", "text/plain")

	// 写入响应主体。
	_, err := w.Write([]byte("hello there\n"))
	if err != nil {
		return err
	}

	// 如果没有错误，则返回 nil。
	return nil
}
