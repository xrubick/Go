package main

import "fmt"

type person interface {
	call()
}

type student struct {
	name string
	age  int
}

//值接收者和指针接收者的区别

// //值接收者
// func (s student) call() {
// 	fmt.Println("chinese")
// }

//指针接收者
func (s *student) call() {
	fmt.Println("chinese")
}

func main() {

	// //值接收者
	// var p1 person
	// var s1 = student{name: "xiaoming", age: 20} //student
	// var s2 = &student{name: "xiaohua", age: 18} //*student
	// p1 = s1
	// fmt.Println(p1) //{xiaoming 20}
	// p1 = s2
	// fmt.Println(p1) //&{xiaoming 20}

	//指针接收者
	var p1 person
	var s1 = student{name: "xiaoming", age: 20} //student
	var s2 = &student{name: "xiaohua", age: 18} //*student
	p1 = &s1
	fmt.Println(p1) //{xiaoming 20}
	p1 = s2
	fmt.Println(p1) //&{xiaoming 20}
}
