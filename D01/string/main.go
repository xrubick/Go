package main

import (
	"fmt"
	"strings"
)

//字符串使用双引号"" 包裹的
//单引号''包裹的是字符
//字节 1字节=8bit(8个二进制位)，一个字符'A'=1个字节，一个utf8编码汉字'好'= 一般占3个字节

//转义符号
// \r  \n \\ \' \"

func main() {

	path := "\"F:\\go\""
	fmt.Println(path)
	//定义多行字符串

	s2 := `
窗前名誉光

屋里没有床

`
	fmt.Println(s2)

	//字符串相关操作
	fmt.Println(len(s2))

	//字符串拼接

	s3 := "hello"
	s4 := "world"

	ss1 := s3 + s4
	fmt.Println(ss1)

	ss2 := fmt.Sprintf("%s%s", s3, s4)
	fmt.Println(ss2)

	//分隔
	sSplit := strings.Split(path, "\\")

	fmt.Println(sSplit)

	//包含 strings.Contains
	//前后缀 strings.HasPrefix strings.HasSuffix

}
