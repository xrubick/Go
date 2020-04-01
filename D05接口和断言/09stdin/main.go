package main

import (
	"bufio"
	"fmt"
	"os"
)

//fmt.Scan
func useScan() {
	var s string
	fmt.Println("输入内容：")
	fmt.Scanln(&s) //输入值有空格，则返回值异常
	fmt.Printf("内容是：%v\n", s)
}

//bufio

func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)

	s,_ = reader.ReadString('\n')
	fmt.Printf("内容是：%v",s)
}

func main() {

	//useScan()
	useBufio()
}
