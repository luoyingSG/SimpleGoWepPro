package main

import "net/http"

func main() {
	// 注册一个函数，使它能够相应 HTTP 请求
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	// 启动 http 服务
	http.ListenAndServe("localhost:8080", nil) // 使用 DefaultServeMux
}
