package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//创建客户端连接 调用Dial函数
	client, err := net.Dial("tcp", "localhost:2377")
	if err != nil {
		fmt.Println("客户端连接失败:", err)
		return
	}
	fmt.Printf("连接成功,连接地址是%v\n", client)
	//通过os.Stdin发送数据
	reader := bufio.NewReader(os.Stdin)
	//读取终端的数据
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("输入错误", err)
	}
	//向服务器发送数据
	fa, err := client.Write([]byte(str))
	if err != nil {
		fmt.Println("发送失败", err)
	}
	fmt.Printf("发送成功%v\n", fa)
}
