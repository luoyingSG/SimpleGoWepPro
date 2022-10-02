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

## 内置的 Handler

```go
// go 语言内置的5个handler  NOTFoundHandler  redirectHandler stripPrefixHandler  TimeoutHandle fileserver

func NOTFoundHandler() Handler
// 返回一个handler， 它给的每个请求的响应都是"404 page not found"

func RedirectHandler(url string ,code int) Handler {}
// 返回一个 handler 它的每个请求使用给定的状态码跳转到指定的 URL
 // URL 要跳转的 URL code 跳转的状态码（3xx）常见的 statusMovedPermanentrly statusFound statusseeother

 func StripPrefixHandler(prefix string, h handler) handler
 // 返回一个 handler 它从请求 URL 中去掉指定的前缀 然后再调用另一个 handler
 // 如果请求 URL 与提供的前缀不符合 那么 404
 // 有点像中间件 prefix URL 将要被移除的字符串前缀
 // h 是一个 handler 在移除字符串前缀之后 handler 将会接受到请求

func TimeoutHandler(h handler, dt time.Duration, msg string){
// 返回一个 handler 他用来在指定时间内运行传入的 h
// h 将要被修饰的handler
// dt 第一个 handler 允许的处理时间
// msg 如果超时那么就把 msg 返回给请求表示响应时间过长

func fileserver(root FileSystem) Handler
// 返回一个 handler 基于root的文件系统来响应请求
// type FileSystem struct{
// 	Open (name string)(File,error)
// }
// 使用时需要到操作系统的文件系统，所以还需要委托给 type dir string
// func(d dir)Open(name string)(File,error)

```