package main

import (
	"fmt"
	"net"
)

/*
在第一次看到这个UDP 服务器的编程模型的时候， 我就好奇这怎么支持并发，
从代码里我们可以看到 服务器只有一个socket。所有的客户端都是通过同一个socket进行通信。
对于TCP 服务器来说，有一个新的客户端连接的时候，会产生一个新的socket用于和新客户端通信。
*/
func main() {
	// 建立 udp 服务器
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})
	if err != nil {
		fmt.Printf("listen failed error:%v\n", err)
		return
	}
	defer listen.Close() // 使用完关闭服务

	for {
		// 接收数据
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read data error:%v\n", err)
			return
		}
		fmt.Printf("addr:%v\t count:%v\t data:%v\n", addr, n, string(data[:n]))
		// 发送数据
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Printf("send data error:%v\n", err)
			return
		}
	}
}
