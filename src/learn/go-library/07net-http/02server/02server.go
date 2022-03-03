package main

import (
	"fmt"
	"net/http"
)

/*
· 默认的Server：
ListenAndServe使用指定的监听地址和处理器启动一个HTTP服务端。处理器参数通常是nil，
这表示采用包变量DefaultServeMux作为处理器。
Handle和HandleFunc函数可以向DefaultServeMux添加处理器。
http.Handle("/foo", fooHandler)
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})
log.Fatal(http.ListenAndServe(":8080", nil))

默认的Server示例

使用Go语言中的net/http包来编写一个简单的接收HTTP请求的Server端示例，
net/http包是对net包的进一步封装，专门用来处理HTTP协议的数据。具体的代码如下：

自定义Server：
要管理服务端的行为，可以创建一个自定义的Server：

s := &http.Server{
	Addr:           ":8080",
	Handler:        myHandler,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe())
*/

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello 沙河！")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
