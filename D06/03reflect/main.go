package main

import (
	"fmt"
	"reflect"
)

type myInt int32

func reflectType(j interface{}) {
	// t := reflect.TypeOf(j)
	// fmt.Printf("类型:%v\n", t)
	t := reflect.TypeOf(j)
	fmt.Printf("type: %v kind: %v\n", t.Name(), t.Kind())
}

type cat struct {
	name string
	age  int
}

type dog struct {
	name string
	age  int
}

func main() {
	var a int32 = 38
	var c myInt = 40
	reflectType(a)
	reflectType(c)
	var b = "hello"
	var d *string
	reflectType(b)
	reflectType(d)

	var miao = cat{
		name: "小花",
		age:  2,
	}

	var wang = dog{
		name: "二哈",
		age:  3,
	}

	reflectType(miao)
	reflectType(wang)
	//反射中像数组、切片、Map、指针等类型的变的.Name()都是返回空
}
