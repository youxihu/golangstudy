package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
func main() {
	fmt.Println("服务端已启动,等待连接")
	listener, err := net.Listen("tcp", "127.0.0.1:2377")
	if err != nil {
		fmt.Println("服务端创建失败,", err)
		return
	}
	//等待客户端连接
	for {
		Conn, err := listener.Accept()
		if err != nil {
			fmt.Println("等待连接失败,err=", err)
		} else {
			fmt.Printf("等待连接成功,conn=%+v，具体连接到的客户端信息是%v\n", &Conn, Conn.RemoteAddr().String())
		}

		//使用协程处理客户端请求
		go process(Conn)
	}
}
