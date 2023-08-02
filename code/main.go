// 包 main 是应用程序的入口点。
package main

// 导入必要的库。
import (
	"context"
	"fmt"
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

	// 在这里接收A方法的请求，并依次存入值名对应的变量中
	// 例如：A方法的请求为：http://localhost:9000/2016-08-15/proxy/FCGoDemo/A?name=fc&age=18
	// 则req.URL.Path为：/2016-08-15/proxy/FCGoDemo/A
	err := req.ParseForm()
	if err != nil {
		return err
	}
	paramValues := make([]string, 0)
	for _, paramName := range req.Form {
		paramValue := req.Form.Get(paramName[0])
		if paramValue != "" {
			paramValues = append(paramValues, fmt.Sprintf("%s is %s", paramName, paramValue))
		}
	}
	// 设置 Content-Type 头部为 text/plain。
	w.Header().Add("Content-Type", "text/plain")

	// 写入响应主体。
	for _, paramValue := range paramValues {
		_, err := w.Write([]byte(fmt.Sprintf("%s\n", paramValue)))
		return err
	}
	// 如果没有错误，则返回 nil。
	return util.EndErrorProcess(err, w)
}
