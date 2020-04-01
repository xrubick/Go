package main

import "fmt"

//类型断言

func assert(a interface{}) {
	fmt.Printf("类型：%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("string", t)
	case int:
		fmt.Println("int", t)
	case bool:
		fmt.Println("bool", t)
	case int32:
		fmt.Println("int32", t)
	}
}

func main() {

	assert("test")
	assert(4)
	assert(true)
}
