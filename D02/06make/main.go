package main

import "fmt"

//make()函数创建切片，并初始化

func main() {

	s1 := make([]int, 3, 5)
	fmt.Println(s1, len(s1), cap(s1))

	//判断切片是否为空，应该用len(s1) == 0, 而不是 s1 == nil

	//切片的遍历
	s1 = []int {2,4,6,8}
	for i:=0; i< len(s1);i ++ {
		fmt.Println(i,s1[i])
	}
	for k,v := range s1 {
		fmt.Println(k,v)
	}
}
