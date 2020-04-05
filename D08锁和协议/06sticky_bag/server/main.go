package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/xrubick/Go/D08锁和协议/06sticky_bag/protocol"
)

//粘包问题复现
//Server

func processer(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		// var buf [1024]byte

		// n, err := reader.Read(buf[:])
		// if err != nil {
		// 	fmt.Println("读取错误：",err)
		// 	break
		// }
		// recStr := string(buf[:n])
		// fmt.Println("收到的消息是：",recStr)

		//解码
		msg, err := protocol.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("解码错误:", err)
			return
		}
		fmt.Println("接收到的消息：", msg)
	}
}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:40000")
	if err != nil {
		fmt.Println("监听端口失败：", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接收失败:", err)
			continue
		}
		go processer(conn)
	}
}
