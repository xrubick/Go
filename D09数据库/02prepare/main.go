package main

//预处理

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //一个连接池对象

//初始化数据库函数
func initDB() (err error) {
	//定义数据源名称
	dsn := "root:root@tcp(3.1.16.61:3306)/go" //密码为空
	db, err = sql.Open("mysql", dsn)          //判断数据源格式
	if err != nil {
		fmt.Println("数据源格式错误:", err)
		return
	}
	//defer db.Close()
	fmt.Println("数据源格式正确")

	err = db.Ping() //通过账号测试数据库连接
	if err != nil {
		fmt.Println("连接数据库失败:", err)
		return
	}
	fmt.Println("连接成功")
	//设置连接池最大连接数
	db.SetMaxOpenConns(5)
	//设置最大空闲连接数
	db.SetMaxIdleConns(6)
	return
}

type user struct {
	id   int
	name string
	age  int
}

// 预处理插入
// SQL语句拆分成两部分：命令部分与数据部分
// 先把命令部分发送给MySQL服务端，进行SQL预处理
// 再把数据部分发送给MySQL服务端，对SQL语句进行占位符替换
// 服务端执行完整SQL语句后，再将结果返回给客户端

func preInsert() {
	sqlStr := "insert into user(name, age) values (?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed:", err)
		return
	}
	defer stmt.Close()
	var m1 = map[string]int{
		//"柯基": 1,
		//"藏獒": 6,
		"test": 0,
	}
	for k, v := range m1 {
		_, err = stmt.Exec(k, v)
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
			return
		}
	}
	fmt.Println("insert success.")
}

func main() {
	initDB()
	preInsert()
}
