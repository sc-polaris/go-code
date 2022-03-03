package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
· TCP协议：
	TCP/IP(Transmission Control Protocol/Internet Protocol) 即传输控制协议/网间
协议，是一种面向连接（连接导向）的、可靠的、基于字节流的传输层（Transport layer）通信协议
，因为是面向连接的协议，数据像水流一样传输，会存在黏包问题。

· TCP服务端
	一个TCP服务端可以同时连接很多个客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝
网。因为Go语言中创建多个goroutine实现并发非常方便和高效，所以我们可以每建立一次链接就创建
一个goroutine去处理。

TCP服务端程序的处理流程：
	1. 监听端口
	2. 接收客户端请求建立链接
	3. 创建goroutine处理链接。
我们使用Go语言的net包实现的TCP服务端代码如下：
*/

// TCP server端

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来端数据：", recvStr)
		conn.Write([]byte("server端：" + recvStr)) // 发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}
