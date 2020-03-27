package main

import (
	//	"encoding/json"
	"encoding/json"
	"fmt"
)

//结构体和json格式转换

//序列化：Go语言结构体变量--->json格式的字符串
//反序列化：Go语言结构体变量<---json格式的字符串
type person struct {
	// name string
	// age  int
	// Name string
	// Age  int
	Name string `json:"name" db:"name"`
	Age  int `json:"age"`
}

func main() {
	p1 := person{
		// name: "小岛",
		// age:  18,
		// Name: "小岛",
		// Age:  18,
		Name: "小岛",
		Age:  18,
	}
	//序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	fmt.Printf("%#v\n", string(b))

	//反序列化
	str := `{"name":"小津","age":20}`
	var p2 person
	json.Unmarshal([]byte(str),&p2)
	fmt.Printf("%#v",p2)
}
