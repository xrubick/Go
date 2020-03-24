package main

import "fmt"
//import "math"

func main () {
	//math.MaxFloat32
	f1 := 1.1234
	fmt.Printf("%T",f1)  //小数默认类型为float64类型
	var f2 float32 = 32.3232
	fmt.Printf("%T",f2)
}