package main

import "fmt"

func main() {

	// //append()为切片追加元素
	// s1 := []string{"a", "b", "c", "d"}
	// fmt.Println(s1[0])
	// //s1[4] = "e"   //index out of range [4] with length 4
	// fmt.Println(len(s1),cap(s1))

	// s1 = append(s1,"e")  //append追加元素，原来的数组放不下时，go会把底层数组换一个
	// fmt.Println(len(s1),cap(s1))

	// s2 :=[]string{"上海","呜","en","a"}

	// s1 = append(s1,s2...)
	// fmt.Println(s1,len(s1),cap(s1))

	//copy
	a1 := []int{1, 2, 3,5,7,8}

	a2 := a1
	var a3 = make([]int, 3, 3)
	copy(a3,a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	//切片删除
	a1 = append(a1[:1],a1[2:]...)
	fmt.Println(a1)


}
