package main

import "fmt"

//匿名字段
type person struct {
	int
	string
}

func main() {

	p1 := person{
		20,
		"test",
	}
	fmt.Println(p1)
	fmt.Println(p1.string)

}
