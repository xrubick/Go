package main

import "fmt"

//模拟实现继承
type animal struct {
	name string
}

func (a animal) move() {
	fmt.Println("会走")
}

type cat struct {
	name string
	animal
}

func (c cat) miao() {
	fmt.Printf("%s会走，喵", c.animal.name)
}
func main() {

	c1 := cat{
		name:   "小花",
		animal: animal{name: "重名"},
	}
	c1.miao()
	c1.move()

}

//使用变量name，结果显示不够简明
