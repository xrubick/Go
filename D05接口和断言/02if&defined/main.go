package main

import "fmt"

//接口的定义：用来给参数、变量、返回值设置类型
/*
type 接口名字 interface {
	方法名字1(参数1,参数2...) (返回值1,返回值2...)
	方法名字2(参数1,参数2...) (返回值1,返回值2...)
	...
}
*/

//接口的实现：如果一个变量实现了接口中的所有方法，也称变量实现了这个接口

type person interface {
	speak(string)
	run()
}

type student struct {
	name string
	age  int
}

func (s student) speak() {
	fmt.Printf("说中文")
}
func (s student) run() {
	fmt.Println("能跑")
}

type teacher struct {
	name string
	age  int
}

func (s teacher) speak(language string) {
	fmt.Printf("说%s\n", language)
}
func (s teacher) run() {
	fmt.Println("能跑")
}

func all(p person) {
	p.speak("english")
	p.run()
}

func main() {
	//第一种
	var p1 person
	var t1 = teacher{
		name: "王老师",
		age:  34,
	}
	p1 = t1 //接口的两部分：(动态类型和动态值) //p1类型：main.teacher;p1值：{王老师 34}
	fmt.Println(p1)
	p1.run()
	p1.speak("English")
	fmt.Printf("p1类型：%T\n", p1)

	//第二种
	all(t1)

}
