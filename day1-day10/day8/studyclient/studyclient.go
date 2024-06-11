package studyclient

import (
	"fmt"
	"net"
)

//CreateClient 创建TCP客户端，尝试连接本机tcp的2375端口

func CreateClient() {
	fmt.Println("客户端 启动！")
	conn, err := net.Dial("tcp", "192.168.0.214:6379")
	if err != nil {
		fmt.Println("客户端连接失败:err", err)
		return
	}
	fmt.Println("连接成功，conn", conn)
}
