package main

import (
	"flag"
	"fmt"
)

//命令行参数解析flag
//支持的命令行参数类型有bool、int、int64、uint、uint64、float float64、string、duration

func main() {

	// //命令行无输入
	// name := flag.String("name", "二哈", "input your wants name")
	// age := flag.Int("age",2,"input your wants age")
	// fmt.Printf("wants name: %v wants age: %v", *name,*age)

	//支持命令行参数
	/*
		-flag xxx
		--flag xxx
		-flag=xxx
		--flag=xxx
	*/
	var name string
	var age int
	flag.StringVar(&name, "name", "边牧", "input your wants name")
	flag.IntVar(&age, "age", 2, "input your wants name")
	flag.Parse()
	fmt.Println(name, age)

	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
