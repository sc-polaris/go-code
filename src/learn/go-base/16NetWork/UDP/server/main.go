package main

import (
	"fmt"
	"net"
)

/*
· UDP协议：
	UDP协议（User Datagram Protocol）中文名称是用户数据报协议，是OSI（Open System
Interconnection，开放式系统互联）参考模型中一种无连接的传输层协议，不需要建立连接
就能直接进行数据发送和接收，属于不可靠的、没有时序的通信，但是UDP协议的实时性比较好，
通常用于视频直播相关领域。

使用Go语言的net包实现的UDP服务端代码如下：
*/

// UDP server端
func main() {
	listenUDP, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listenUDP failed, err:", err)
		return
	}
	defer listenUDP.Close()
	for {
		var data [1024]byte
		n, addr, err := listenUDP.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, err = listenUDP.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}
