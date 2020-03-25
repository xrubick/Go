package main

import (
	"fmt"
	"os"
)

//import "fmt"
//import "os"

func main() {
	for {
		//系统菜单
		fmt.Println("管理系统主页！")
		fmt.Println("功能选择:")

		fmt.Println(`
			 1.查询全体人员
			 2.增加成员
			 3.删除成员
			 4.退出系统
	 `)
		var num int
		fmt.Scanln(&num)

		//系统增删查
		switch num {
		case 1:
			fmt.Println("1.查询全体人员")
		case 2:
			fmt.Println("2.增加成员")
		case 3:
			fmt.Println("3.删除成员")
		case 4:
			os.Exit(1)

		default:
			fmt.Println("输入错误")
		}
	}
}
