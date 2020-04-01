package main

import "fmt"

//结构体  (值类型)

//new初始化返回的是指针，make初始化返回的是值

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func f1(x person) {
	x.gender = "woman" //修改的副本gender
}

func f2(x *person) {
	//(*x).gender = "woman" //根据内存地址找到原变量，修改原来位置的值
	x.gender = "woman" //根据内存地址找到原变量，修改原来位置的值
}

//构造函数：返回一个结构体变量的函数,约定用new开头
//(返回的是结构体还是结构体指针需看结构体本身的大小，减小内存)

type people struct {
	name string
	age  int
}

func newPeople(name string, age int) *people {
	return &people{
		name: name,
		age:  age,
	}
}

func main() {
	// var xiaoming person
	// xiaoming.name = "小明"
	// xiaoming.age = 19
	// xiaoming.gender = "man"
	// xiaoming.hobby = []string{"唱唱歌", "跳跳舞"}

	// fmt.Println(xiaoming)
	// fmt.Printf("%T\n",xiaoming)
	// fmt.Println(xiaoming.name)

	// //匿名结构体
	// var s struct {
	// 	x string
	// 	y int
	// }
	// s.x = "hello"
	// s.y = 1024
	// fmt.Printf("%T,%v",s,s)

	// var p person
	// p.name = "小明"
	// p.gender = "man"
	// f1(p)
	// fmt.Println(p.gender)
	// f2(&p)
	// fmt.Println(p.gender)

	// var p2 = new(person) //new初始化返回的是指针
	// (*p2).name = "小芳"
	// p2.gender = "woman"
	// fmt.Printf("%T\n", p2)
	// fmt.Printf("%p\n", p2)
	// fmt.Printf("%p\n", &p2)

	//构造函数
	p1 := newPeople("test", 20)
	fmt.Println(*p1)
}
