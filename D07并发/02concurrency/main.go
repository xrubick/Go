package main

import (
	"fmt"
	"time"
)

//goroutine

func hello(i int) {
	fmt.Println("hello man", i)
}

//程序启动之后会创建一个主goroutine执行main()
func main() {
	for i := 0; i < 2000; i++ {
		//go hello(i) //开启一个单独的goroutine去执行hello函数
		go func(i int) { //匿名函数
			fmt.Println(i) //i代表函数本身在外层复制过来的i
		}(i)
	}

	fmt.Println("main") //main函数结束后，由main启动的goroutine也会结束

	time.Sleep(time.Second * 5)

}
