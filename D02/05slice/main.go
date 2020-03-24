package main

//切片slice
import "fmt"

func main() {

	//定义切片

	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)

	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"a", "b", "c"}
	fmt.Println(s1, s2)

	//长度和容量
	fmt.Printf("len(s1): %d,cap(s1):%d", len(s1), cap(s1))

	fmt.Printf("len(s2): %d,cap(s2):%d", len(s2), cap(s2))

	//由数组得到切片
	a1 := [...]int{1, 3, 5, 6, 9}

	s3 := a1[0:4] //左包含右不包含(左闭右开)
	fmt.Println(s3)

	s4 := a1[:4]
	s5 := a1[2:]
	s6 := a1[:]
	fmt.Println(s4, s5, s6)

	//切片容量指从切片的第一个元素开始到底层数组的最后一个元素
	//切片的长度就是切片本身元素的个数
	fmt.Println("len(s4): %d ,cap(s4): %d", len(s4), cap(s4))
	fmt.Println("len(s5): %d ,cap(s5): %d", len(s5), cap(s5))

	//切片是引用类型，都指向了底层数组
	s6[3] = 6888
	fmt.Println(a1,s6)


}
