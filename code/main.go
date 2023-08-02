// 包 main 是应用程序的入口点。
package main

// 导入必要的库。
import (
	"context"
	"main/util"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

// 主函数启动 HTTP 服务器。
func main() {
	fc.StartHttp(HandleHttpRequest)
}

// HandleHttpRequest 是 HTTP 处理程序函数。
func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	_, err := w.Write([]byte("Hello World!"))
	//err = fmt.Errorf("test error")
	return util.EndErrorProcess(err, w)
}
