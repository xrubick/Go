package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//file.Read()
func osRead() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("wrong:", err)
		return
	}
	//关闭文件
	defer fileObj.Close()
	for {
		//读取文件
		var tmp [128]byte //设置读取长度
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("wrong:", err)
			return
		}
		//fmt.Printf("read %d byte", n)
		fmt.Print(string(tmp[:n]))
	}
}

//bufio
func bufioRead() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("wrong:", err)
		return
	}
	//关闭文件
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("wrong:", err)
			return
		}
		fmt.Print(line)
	}
}

//ioutil
func ioutilRead() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("wrong:", err)
		return
	}
	//关闭文件
	defer fileObj.Close()

	ret, err := ioutil.ReadFile("./main.go")
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Println("wrong:", err)
		return
	}
	fmt.Print(string(ret))
}

func main() {
	//osRead()
	//bufioRead()
	ioutilRead()
}
