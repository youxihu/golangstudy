package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("客户端启动")
	//调用Dial函数：参数需要指定tcp协议,需要指定服务器的IP+PORT
	client, err := net.Dial("tcp", "127.0.0.1:2377")
	if err != nil {
		fmt.Println("客户端连接失败:", err)
		return
	}
	fmt.Printf("客户端连接成功,client=%+v\n", &client)
	//通过客户端发送单行数据
	reader := bufio.NewReader(os.Stdin) // os.Stout终端标准输出 os.Stderr终端标准错误
	//从终端读取一行用户输入的信息
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("终端输入错误", err)
	}
	//将读取到的信息发送给服务器
	fa, err := client.Write([]byte(str))
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Printf("发送成功,发送了%d字节的数据", fa)
}
