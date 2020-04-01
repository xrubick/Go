package main

import "fmt"

//自定义类型和类型别名
type myInt int     //自定义类型
type yourInt = int //类型别名

func main() {

	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T", n)

	var m yourInt
	m = 10
	fmt.Println(m)
	fmt.Printf("%T", m)

}
