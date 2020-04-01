package main

import "fmt"

//接口：也是一种类型，定义了一个对象的行为(变量有哪些方法)

//接口的实例
type caller interface {
	call() //实现call方法的变量都是caller类型
}

type dog struct{}

type cat struct{}

type person struct{}

func (d dog) call() {
	fmt.Println("wang")
}
func (c cat) call() {
	fmt.Println("miao")
}

func da(x caller) {
	x.call()
}

func main() {
	var c1 cat
	var d1 dog
	//var p1 person
	da(c1)
	da(d1)
	//da(p1)
}
