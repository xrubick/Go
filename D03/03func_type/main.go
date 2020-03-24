package main

import "fmt"

func f1() {
	fmt.Println("hello")
}

func f2() int {
	return 5
}

//函数也可以作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

//函数作为返回值
func f4(x func() int) func(int, int) int {
	return func(a, b int) int {
		return a + b
	}
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)

	f3(f2)
	f5 := f4(f2)
	fmt.Printf("%T\n", f5)
}
