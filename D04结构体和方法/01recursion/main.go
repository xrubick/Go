package main

import "fmt"

//递归函数，自己调用自己
//要有明确的退出条件

//计算n的阶乘

func f(n uint) uint {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

func main() {
	fmt.Println(f(6))
}
