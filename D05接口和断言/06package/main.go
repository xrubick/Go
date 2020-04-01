package main

//包 package
//导入路径是在$GOPATH/src后面的路径开始
import "fmt"
//import "github.com/xrubick/Go/D05/06demo"

//每个包需要被调用的标识符，首字母需要大写
//导入时可以指定别名，支持单行和多行导入
//import  test "github.com/xrubick/Go/D05/06demo"

//导入时不需要使用包内的标识符，可以使用匿名导入:“_  packagePath”
import  _ "github.com/xrubick/Go/D05/06demo"
//包导入时会自动执行包内名为:init()的函数，无参数无返回值且不能被调用

func  init() {
	fmt.Println("自身init")
}

func main() {
	//每个包需要被调用的标识符，首字母需要大写
	//ret := add.Add(1,2)
	// ret := test.Add(1,2)
    // fmt.Println(ret)
}