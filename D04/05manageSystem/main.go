package main

import (
	"fmt"
	"os"
)

//import "fmt"
//import "os"

//结构体
type stu struct {
	id   int
	name string
}

//构造函数
func newStu(id int, name string) *stu {
	return &stu{
		id:   id,
		name: name,
	}
}

var all = make(map[int]*stu, 20)

func listAll() {
	for k, v := range all {
		fmt.Printf("id:%d name:%s\n", k, v.name)
	}
}

func addSomeone() {
	var (
		id   int
		name string
	)
	fmt.Println("id:")
	fmt.Scanln(&id)
	fmt.Println("name:")
	fmt.Scanln(&name)
	newStu := newStu(id, name)
	all[id] = newStu

}

func delSomeone() {
	var delID int
	fmt.Println("delId:")
	fmt.Scanln(&delID)

	delete(all, delID)

}

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
			listAll()
		case 2:
			fmt.Println("2.增加成员")
			addSomeone()
		case 3:
			fmt.Println("3.删除成员")
			delSomeone()
		case 4:
			os.Exit(1)

		default:
			fmt.Println("输入错误")
		}
	}
}
