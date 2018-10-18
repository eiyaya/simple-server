package main

import (
	"fmt"
	"net"
)

func main() {
	ip := "127.0.0.1"
	port := "6666"
	address := ip + ":" + port
	l, err := net.Listen("tcp4", address) // Listen announces on the local network address.
	if err != nil {
		fmt.Println("监听端口出错!")
	}
	for {
		conn, err := l.Accept() // 
		if err != nil {
			fmt.Println("客户端连接失败")
			conn.Close()
		}

		handleMessage(conn)
	}
}

func handleMessage(conn net.Conn){
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		nByte, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("read %d byte from client!\n", nByte)
		s := string(buf[:nByte])
		fmt.Println(s)
	}

}
