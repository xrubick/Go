package main

import "fmt"

//pointer
func main() {
	// //&:取地址
	// n := 8
	// fmt.Println(&n)
	// fmt.Printf("%T\n", n)
	// //*:根据地址取值
	// p := &n
	// fmt.Printf("%T\n", p)

	// v := *p

	// fmt.Println(v)
	// fmt.Printf("%T", v)
	var a1 *int // nil pointer
	fmt.Println(a1)
	//new用于基本数据类型sring、int的内存申请，返回的是对应类型的指针pointer；make用于slice、map、chan的内存创建
	var a2 = new(int)
	fmt.Println(a2)
	fmt.Println(*a2)

	*a2 = 100
	fmt.Println(*a2)

}
