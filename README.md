# SimpleGoWepPro

A simple go web project

## Handler

handler 是一个接口（interface）
handler 定义了一个方法 ServeHTTP()，参数：

- HTTPResponseWriter
- 指向 Request 这个 struct 的指针（包含 HTTP 请求相关的一些信息）

自行实现一个 Handler：

```go
type myHandler struct {}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, world!"))
}
```

## DefaultServeMux

![简介](./DefaultServeMux%E7%AE%80%E4%BB%8B.png)

- 它是一个 Multiplexer（多路复用器）
- 它**也是一个 Handler**

## 多个 Handler

要使用多个 Handler 处理多个路径的请求，需要的步骤为：

1. 不指定 Server struct 里面的 Handler 字段值
2. 可以使用 http.Handle 将某个 Handler 附加到 DefaultServeMux
   
   http 包有一个Handle 函数
   
   ServeMux struct 也有一个 Handle 方法
3. 如果你调用 http.Handle，实际上调用的是 DefaultServeMux 上的 Handle 方法
   
   DefaultServeMux 就是 ServeMux 的指针变量

例如：

```go
mh := helloHandler{}
server := http.Server{
    Addr: "localhost:8080",
    Handler: nil // 不指定 Handler
}
// 将 Handler 附加到 DefaultServeMux
http.Handle("/hello", &mh) // 注意，http.Handle 需要的是 Handler 指针
```