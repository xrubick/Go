package main

import "fmt"

//同一个结构体可实现多个接口
//接口可以嵌套

//空接口
// interface{} //空接口类型  //所有类型都实现了空接口

type person interface {
	speaker
	runner
}

type speaker interface {
	speak(string)
}

type runner interface {
	run()
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

func show(a interface{}) {
	fmt.Printf("类型：%T 值：%v\n", a, a)
}

func main() {
	// var p2 person
	// p2 = teacher{
	// 	name: "teacher",
	// 	age:  38,
	// }
	// fmt.Println(p2)
	// p2.speak("english")
	// p2.run()

	//空接口
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 5)
	m1["name"] = "test"
	m1["age"] = 20
	m1["merried"] = false
	m1["hobby"] = [...]string{"sing", "dancing"}
	fmt.Println(m1)
	fmt.Printf("类型：%T\n", m1["hobby"])

	show(false)
	show(6)
	show(m1)
}
