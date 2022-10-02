package main

import "net/http"

func main() {
	// 注册一个函数，使它能够相响应 HTTP 请求
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	// 启动 http 服务
	// http.ListenAndServe：调用 http.Serve 对象的 serve.ListenAndServe()
	// 1. 网络地址，如果是 ""，那么就是所有网络接口的 80 端口
	// 2. handler，如果为 nil，那么就是 DefaultServeMux，它是一个 multiplexer，可以看作是路由器
	http.ListenAndServe("localhost:8080", nil) // 使用 DefaultServeMux
}
