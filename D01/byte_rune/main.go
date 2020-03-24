package main

import "fmt"

//byte类型和rune类型
//byte类型 : A 、B、C
//rune类型 : 好、天、ど、う、も

func main() {
	// s := "hello 你好 どうも"
	// //len()求取的是byte字节的数量
	// n := len(s)
	// fmt.Println(n)

	// for _,c := range s {
	// 	fmt.Printf("%c",c)
	// }

	//字符串修改，字符串本身不支持修改
	s1 := "hello"
	c1 := []byte(s1)
	c1[0] = 'g'
	fmt.Println(string(c1))

	//类型转换

}
