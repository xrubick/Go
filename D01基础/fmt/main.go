package main


import "fmt"

func main() {
	n := 100

	//查看类型
	fmt.Printf("%T\n",n)
	fmt.Printf("%v\n",n)

	fmt.Printf("%b\n",n)
	fmt.Printf("%x\n",n)
	fmt.Printf("%o\n",n)
	fmt.Printf("%d\n",n)

	var s1 string  = "hello world"

	fmt.Printf("%s\n",s1)
	fmt.Printf("%v\n",s1)
	fmt.Printf("%#v\n",s1)



}
