package main

//粘包问题复现
//Client

import (
	"fmt"
	"net"

	"github.com/xrubick/Go/D08锁和协议/06sticky_bag/protocol"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:40000")
	if err != nil {
		fmt.Println("建立连接失败：", err)
		return
	}

	defer conn.Close()
	for i := 0; i < 5; i++ {
		msg := `How old are you? 88888888888888888888888`
		//conn.Write([]byte(msg))

		//封包
		data, err := protocol.Encode(msg)
		if err != nil {
			fmt.Println("封包错误:", err)
			return
		}
		conn.Write(data)
	}
}
