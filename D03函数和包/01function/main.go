package main

import "fmt"

//一个命名的函数中不能再声明一个命名函数
//defer用于函数结束之前释放资源（数据库连接，文件句柄，socket连接）
func deferDemo() {
	//defer语句
	fmt.Println("start")
	defer fmt.Println("1") //defer后面的语句延迟到函数即将结束并返回值时，再倒序执行
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("end")
}

//GO语言中函数的return不是原子操作，在底层分为两步执行：第一步：返回值赋值，第二步：真正RET返回值
//defer执行时机在第一步和第二步之间
func f1() int {
	x := 3
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 3
}

func f3() (y int) {
	x := 3
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++ //改变的是参数x的副本
	}(x)
	return 3
}

func main() {
	//deferDemo()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

}
