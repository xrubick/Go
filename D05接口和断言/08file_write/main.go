package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//file write
func writeFile() {
	file, err := os.OpenFile("./test.md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("wrong:%v", err)
		return
	}
	defer file.Close()

	//write
	file.Write([]byte("this is a test string\n"))

	//writeString
	file.WriteString("zzl")
}

//bufio.NewWrite
func bufioWrite() {
	file, err := os.OpenFile("./test.md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("wrong:%v", err)
		return
	}
	defer file.Close()

	wr := bufio.NewWriter(file)
	wr.WriteString("hello") //写入缓存
	wr.Flush()              //将缓存的内容写入文件
}

//ioutil.WriteFile
func ioutilWrite() {
	str := "test ioutil"
	err := ioutil.WriteFile("./test.md",[]byte(str),0644)
	if err != nil {
		fmt.Printf("wrong:%v",err)
		return
	}
}

func main() {
	//writeFile()
	//bufioWrite()
	ioutilWrite()
}
