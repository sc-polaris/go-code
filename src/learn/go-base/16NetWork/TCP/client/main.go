package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/*
一个TCP客户端进行TCP通信的流程如下：
	1. 建立与服务端的链接
	2. 进行数据收发
	3. 关闭链接
使用Go语言的net包实现的TCP客户端代码如下：

	将下面的代码编译成client或client.exe可执行文件，先启动server端再启动client端，
在client端输入任意内容回车之后就能够在server端看到client端发送的数据，从而实现TCP通信。
*/

// 客户端
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err :", err)
		return
	}

	defer conn.Close() // 关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err := conn.Write([]byte("client端" + inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
