package main

import (
	"fmt"
	"sync"
)

//channel声明
//var 变量名 chan 变量类型

var a chan int //需要指定chan中元素类型，chan为引用类型
var b chan int

//make函数初始化slice\map\chan

//chan接收值   ch <-
//chan取值     <-ch  或者 x := <-ch
//关闭close(ch)

var wg sync.WaitGroup

//无缓冲
func nobuff() {
	fmt.Println(a)
	a = make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-a
		fmt.Println(x)

	}()
	a <- 10
	wg.Wait()
}

//有缓冲
func hasbuff() {
	fmt.Println(b)
	b = make(chan int, 2)
	b <- 8
	b <- 10
	//b <- 12
	y := <-b
	z := <-b
	fmt.Println(b, y, z)
	close(b)
}

//1.启动goroutine生成100个数发送到ch1；
//2.启动一个goroutine,从ch1中取值，计算其平方数放入ch2中；
//3.打印ch2中的值

//func f1(ch1 chan int) {
func f1(ch1 chan<- int) {  //单项通道，只能接收
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

var once sync.Once

func f2(ch1 <-chan int,  ch2 chan<-  int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	//close(ch2)
	once.Do(func() { close(ch2) })
}
func main() {
	// fmt.Println(a)

	// a = make(chan int) //通道初始化

	// b = make(chan int, 10) //带缓冲区的通道初始化
	// fmt.Println(a)
	// fmt.Println(b)

	//nobuff()
	//hasbuff()

	a = make(chan int, 100)
	b = make(chan int, 100)
	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}

}
