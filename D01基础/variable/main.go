package main

import "fmt"

//声明变量,推荐使用驼峰式命名

//var name string
//var age int
//var ok bool

var (
	name string // ""
	age  int    // 0
	ok   bool   // false
)

//非全局变量声明必须使用

func main() {
	name = "xcf"
	age = 18
	ok = true
	fmt.Println(name)
	fmt.Printf("%d", age)

	//声明变量并赋值
	var s1 string = "testString"
	//类型推导
	var s2 = "testString"
	//简短变量声明,只能在函数中声明
	s3 := "testString"

	fmt.Println(s1, s2, s3)

	//匿名变量用 "_" 表示
	//同一个作用域{}不能声明同名变量

}
