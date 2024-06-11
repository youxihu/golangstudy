package studyclient

import (
	"fmt"
	"net"
)

func CreateServer() {
	fmt.Println("studyclient server start")
	ser, err := net.Listen("tcp", "192.168.0.214:2379")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	defer func(ser net.Listener) {
		err := ser.Close()
		if err != nil {
			return
		}
	}(ser)
	fmt.Println("监听成功，ser=%v,接收到的客户端信息=%v\n", ser, ser.Addr().(*net.TCPAddr).Port)
}
