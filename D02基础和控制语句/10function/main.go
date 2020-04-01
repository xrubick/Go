package main

import "fmt"

//函数,一段代码的封装，将逻辑抽象到一个函数，需要时调用函数即可

//函数的定义

func sum(x int, y int) (ret int) {
	return x + y
}

//有参数无返回值
//无参数无返回值的函数
//无参数有返回值的函数
//多个返回值

//返回值可以命名也可不命名
//命名的返回值相当于在函数中声明了变量，且return后可省略

//参数类型简写
func f1(x, y int, i, j string, f, t bool)  int {
	return x + y
}
//可变长参数
//可变长参数的必须放在函数参数最后
func f2 (s string, t ...int) {
	fmt.Println(s)
	fmt.Println(t)
}

func main() {
	r := sum(1, 3)
	fmt.Println(r)

	f2("hello")
	f2("hello",1,2,3,4,5)
}
