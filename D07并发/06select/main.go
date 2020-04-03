package main

import "fmt"

//select多路复用
/*
select{
    case <-ch1:
        ...
    case data := <-ch2:
        ...
    case ch3<-data:
        ...
    default:
        默认操作
}
*/

//处理一个或多个channel的发送/接收操作。
//多个case同时满足，select会随机选择一个。
//对于没有case的select{}会一直等待，可用于阻塞main函数。

func main() {
	ch1 := make(chan int, 1)
	for i := 0; i < 20; i++ {
		select {
		case j := <-ch1:
			fmt.Println(j)
		case ch1 <- i:
		}
	}
}
