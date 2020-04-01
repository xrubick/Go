package main

import (
	"fmt"
	"strings"
)

//匿名函数
// var f1 = func(x, y int) {
// 	fmt.Println(x + y)
// }

//闭包也是一个函数，其函数包含了他外部作用域的变量
//原理：函数可以作返回值；函数内部查找变量时先在自己内部找，再向外部找。
func f1(f func()) {
	f()
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {

	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//f1(f2)

func f3(f func(int, int), i, j int) func() {
	tmp := func() {
		//fmt.Println(x + y)
		//f2
		f(i, j)
	}
	return tmp
}

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}

}

func makeSuffix(suffix string) func(string) string {
	tmp := func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
	return tmp
}

func main() {
	// //函数内部没法声明命名的函数
	// f1 := func(x, y int) {
	// 	fmt.Println(x + y)
	// }
	// f1(10, 2)

	// //若只调用一次，可写成立即执行函数
	// func (x,y int) {
	// 	fmt.Println(x + y)
	// 	fmt.Println("hello")
	// }(100,200)
	// ret := adder(100)
	// fmt.Printf("%T", ret)
	// retnum := ret(300)
	// fmt.Println(retnum)

	// ret := f3(f2, 1, 2)
	// f1(ret)
	jpg := makeSuffix(".jpg")
	txt := makeSuffix(".txt")
	fmt.Println(jpg("test"))
	fmt.Println(jpg("call.jpg"))

	fmt.Println(txt("test"))
	fmt.Println(txt("call.txt"))



}
