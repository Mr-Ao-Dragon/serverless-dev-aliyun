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
	// 在这里写入业务逻辑。
	// Some code here......
	// 返回 HTTP 响应。
	w.WriteHeader(http.StatusOK)
	return nil
}
