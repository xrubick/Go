package main

import (
	"bufio"
	"fmt"
	"net"
)

//TCP Server

//处理连接函数

func processer(conn net.Conn) {
	//关闭连接
	defer conn.Close()
	for {

		//读取数据
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n,err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("读取失败:",err)
			break
		}
		resString := string(buf[:n])
		fmt.Println("接收到客户端数据：",resString)

		//发送数据
		conn.Write([]byte(resString))
	}
}

func main() {
	//监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:12000")

	if err != nil {
		fmt.Println("监听端口失败:", err)
		return
	}

	for {
		//建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("连接失败：", err)
			continue
		}

		//启动goroutine处理连接
		go processer(conn)
	}
}
