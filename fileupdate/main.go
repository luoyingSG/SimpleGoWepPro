package fileupdate

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func process(w http.ResponseWriter, r *http.Request) {
// 	// r.ParseMultipartForm(1024)

// 	// fileHeader := r.MultipartForm.File["uploaded"][0]
// 	// file, err := fileHeader.Open()
// 	file, _, err := r.FormFile("uploaded")

// 	if err == nil {
// 		data, err := ioutil.ReadAll(file)
// 		if err == nil {
// 			fmt.Fprintf(w, string(data))
// 		}
// 	}
// }

// func main() {
// 	// 注册一个函数，使它能够相响应 HTTP 请求
// 	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	w.Write([]byte("Hello, world!"))
// 	// })

// 	server := http.Server{
// 		Addr: "localhost:8080",
// 	}

// 	// 将文件上传的函数注册
// 	http.HandleFunc("/process", process)

// 	// 启动 http 服务
// 	// http.ListenAndServe：调用 http.Serve 对象的 serve.ListenAndServe()
// 	// 1. 网络地址，如果是 ""，那么就是所有网络接口的 80 端口
// 	// 2. handler，如果为 nil，那么就是 DefaultServeMux，它是一个 multiplexer，可以看作是路由器
// 	server.ListenAndServe() // 使用 DefaultServeMux
// }
