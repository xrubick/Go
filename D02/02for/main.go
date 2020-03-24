package main

import "fmt"

//for循环

func main() {
	//基本格式
	// for i := 1; i < 7; i++ {
	// 	fmt.Println(i)
	// }
	// i := 1
	// for ; i < 5; i++ {
	// 	fmt.Println(i)

	// }
	// b := 0
	// for b < 4 {
	// 	fmt.Println(b)
	// 	b++
	// }

	//for无限循环
	// for {
	// 	fmt.Println("123")

	// }

	//for range循环

	// s := "hello"
	// for k, v := range s {
	// 	fmt.Printf("%d,%c\n", k, v)
	// }

	// for i := 1; i < 10; i ++{
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	// fmt.Println("over")
	for i := 1; i < 10; i ++{
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}
