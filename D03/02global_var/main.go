package main

import "fmt"

//变量的作用域

var x = 10 //定义一个全局变量

//定义函数
//变量查找范围：先在函数内部查找，找不到就往函数外查找，直至查到全局变量
//函数内部定义的变量只能在函数内部使用
func f1() {
	fmt.Println(x)
}
func main() {
	f1()
}
