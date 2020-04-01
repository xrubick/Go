package main

import "fmt"

//map是引用类型，必须初始化(在内存中开辟空间)
func main() {
	var m1 map[string]int
	fmt.Println(m1)

	m1 = make(map[string]int, 8) //估算map的容量
	fmt.Println(m1)
	m1["age"] = 10
	m1["tall"] = 180
	// fmt.Println(m1)
	// fmt.Println(m1["age"])

	// v,ok := m1["height"]
	// if !ok {
	// 	fmt.Println("无指标")
	// } else {
	// 	fmt.Println(v)
	// }

	// //map遍历
	// for k,v := range m1 {
	// 	fmt.Println(k,v)
	// }
	//删除
	delete(m1, "age")
	fmt.Println(m1)

	// //元素类型为map的slice ,注意初始化值
	// var s1 = make([]map[int]int, 2, 10) //slice切片长度为0，s1[0][100] = 100报错
	// s1[0] = make(map[int]int, 2)
	// s1[0][100] = 100
	// fmt.Println(s1)

	//值为切片类型的map
	var m5 = make(map[int][]int, 10)
	m5[1] = []int{1, 2, 3}
	fmt.Println(m5)

}
