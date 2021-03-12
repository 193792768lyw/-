package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s host:port", os.Args[0])
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatal(err)
	}

	//  DialTCP 函数来建立一个 TCP 连接，并返回一个 TCPConn 类型的对象
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	n, err := conn.Write([]byte("HEAD / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n)
	for true {
		by := make([]byte, 999)
		n, _ := conn.Read(by)
		fmt.Println("receive form  server :" + string(by[:n]))
	}
	time.Sleep(8 * time.Second)
}

// go run client.go 127.0.0.1:8000
