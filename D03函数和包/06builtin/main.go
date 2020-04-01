package main

import "fmt"

//panic 和 recover
//recover必须搭配defer使用
//defer必须定义在可能发生panic的语句之前

func a() {
	fmt.Println("a")
}
func b() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("修复错误发生")
		}
		fmt.Println("释放连接")
	}()
	panic("错误发生")
	//fmt.Println("b")
}
func c() {
	fmt.Println("c")
}

func main() {

	a()
	b()
	c()

}
