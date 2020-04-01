package main

import "fmt"

//方法

//标识符：变量名，函数名，类型名，方法名
//Go语言中，标识符首字母大写，表示是共有的，能被别的包调用

//结构体
type cat struct {
	name string
	age int
}

//构造函数
func newCat(name string,age int) cat {
	return cat{
		name: name,
		age: age,
	}
}

//方法是用于特定类型的函数
//接收者表示调用该方法的具体类型变量
func (c cat) miao() {
	fmt.Printf("%s: 喵喵喵", c.name)
}

//值接收者和指针接收者
func (c cat ) age1(){
	c.age ++ 
}

func (c *cat) age2(){
	c.age++
}


func main() {

	// c1 := newCat("test")
	// c1.miao()

	c2 := newCat("1",3)
	fmt.Println(c2.age)
	c2.age1()
	fmt.Println(c2.age)
	c2.age2()
	fmt.Println(c2.age)


}
