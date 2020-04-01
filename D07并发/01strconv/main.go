package main

import (
	"fmt"
	"strconv"
)

//strconv

func main() {

	//数字转换成字符串类型
	i := int(24)
	//s1 := string(i)  //此时会把i当作字符编码处理，返回对应字符值
	st2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v\n", st2)

	//从字符串中解析出对应的数据
	st3 := "200"

	v1, err := strconv.ParseInt(st3, 10, 64)
	if err != nil {
		fmt.Println("wrong:", err)
		return
	}

	fmt.Printf("%#v %T\n", v1, v1)

	i1, _ := strconv.Atoi(st3)
	fmt.Printf("%#v  %T", i1, i1)
}
