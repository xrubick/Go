package main

import "fmt"

//数组array
//必须指定存放元素的类型和容量

func main() {
	/*
		var a1 [3]bool
		var a2 [4]bool
		fmt.Printf("%T,%T\n",a1 ,a2)

		//数组初始化
		fmt.Println(a1,a2)

		a1 = [3]bool{true,true,false}
		fmt.Println(a1,a2)

		a3 := [...]int{1,2,3,4}
		fmt.Println(a3)

		a4 := [6]int{1,2}
		fmt.Println(a4)

		a5 := [6]int{1:1,4:5}
		fmt.Println(a5)
	*/

	//数组遍历
	citys := [...]string{"b", "s", "g"}

	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	for k, v := range citys {
		fmt.Println(k, v)
	}

	//多维数组
	//[[1 2] [3 4] [5 6]]
	var all [3][2]int
	all = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(all)

	for _, v1 := range all {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//数组是值类型，赋值和传参会复制整个数组，改变副本，不会改变数组本身
	b1 := [3]int{1, 2, 3} //[1 2 3]
	b2 := b1              //[1 2 3]  ctrl+c ctrl+v
	b2[0] = 10            //[10 2 3]
	fmt.Println(b1, b2)
}
