package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//TCP Client

func main() {
	//发起连接
	conn, err := net.Dial("tcp", "127.0.0.1:12000")
	if err != nil {
		fmt.Println("发起连接失败：", err)
		return
	}
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)

	for {
		//读取用户输入
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")

		//退出条件
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}

		//发送数据
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}

		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("接收值失败：", err)

			return
		}
		fmt.Println(string(buf[:n]))
	}
}
