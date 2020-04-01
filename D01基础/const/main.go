package main

import "fmt"

const pi = 3.14

//批量声明常量
const (
	statusOK = 200
	notFound = 404
)

//批量声明未赋值，默认和上一行一致

const (
	n1 = 100
	n2
	n3
)

//iota 在const关键字出现时，将被重置为0.const中没新增一行常量声明将使iota计数加1

const (
	a1 = iota //0
	a2        //1
	a3        //2
)

const (
	b1 = iota //0
	b2 = 100  //100
	b3 = iota //2
)

const (
	d1, d2 = iota + 1, iota + 2 // 1,2
	d3, d4 = iota + 1, iota + 2 //2,3
)

const (
	_  = iota
	//KB unit
	KB = 1 << (10 * iota) //2^10=1024 (向左移动10位)
	MB = 1 << (10 * iota) //2^20
	GB = 1 << (10 * iota) //2^30
	TB = 1 << (10 * iota) //2^40

)

func main() {
	//pi = 2.1
	fmt.Println(pi)
}
