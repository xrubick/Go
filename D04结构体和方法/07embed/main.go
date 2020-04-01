package main

import "fmt"

//结构体嵌套
type school struct {
	addr string
	name string
}

type stu1 struct {
	name   string
	age    int
	school //匿名嵌套结构体

}

type stu2 struct {
	name string
	age  int
	sch  school
}

func main() {
	p1 := stu1{
		name: "小明",
		age:  20,
		school: school{
			"北京",
			"师范",
		},
	}
	fmt.Println(p1.name)
	fmt.Println(p1.addr)
	fmt.Println(p1.school.name)
	// p1 := stu2{
	// 	name: "小明",
	// 	age:  20,
	// 	sch: school{
	// 		"北京",
	// 		"师范",
	// 	},
	// }
	// fmt.Println(p1.name)
	// fmt.Println(p1.sch.addr)

}
