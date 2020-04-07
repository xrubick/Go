package main

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

//查询数据库--单行查询
func queryOne(n int) {
	sqlStr := `select id,name,age from  user where id=?;`
	var user1 user

	// //不释放连接，等待查询(最大连接数是5)
	// for i := 0; i < 6; i++ {
	// 	db.QueryRow(sqlStr, n) //在连接池中拿出一个连接去查询数据库
	// 	fmt.Printf("第%d次查询\n", i)
	// }

	// //执行查询
	// obj := db.QueryRow(sqlStr, n) //在连接池中拿出一个连接去查询数据库
	// //对obj必须使用Scan方法，以释放连接
	// obj.Scan(&user1.id, &user1.name, &user1.age) //获取查询结果

	//执行查询并获取结果
	err := db.QueryRow(sqlStr, n).Scan(&user1.id, &user1.name, &user1.age)
	if err != nil {
		fmt.Println("获取失败:", err)
		return
	}
	fmt.Println(user1)
}

//查询数据库--多行查询
func queryMany(n int) {
	sqlStr := `select id,name,age from  user where id>?;`
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}
	defer rows.Close() //必须释放连接
	for rows.Next() {
		var user2 user
		err := rows.Scan(&user2.id, &user2.name, &user2.age)
		if err != nil {
			fmt.Println("取值失败:", err)
			return
		}
		fmt.Println(user2)
	}
}

//插入数据
func insertData() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "灰短", 1)
	if err != nil {
		fmt.Println("插入失败:", err)
		return
	}
	id, err := ret.LastInsertId() // 获取新插入数据的id
	if err != nil {
		fmt.Println("获取id失败:", err)
		return
	}
	fmt.Println(id)
}

//更新(或者删除)数据
func updateData() {
	sqlStr := "update user set age=? where id = ?"
	//sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 5, 3)
	if err != nil {
		fmt.Println("更新失败:", err)
		return
	}
	n, err := ret.RowsAffected() // 获取操作影响的行数
	if err != nil {
		fmt.Println("获取失败", err)
		return
	}
	fmt.Println("更新成功:", n)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("初始化失败:", err)
	}
	fmt.Println("初始化成功")
	//queryOne(1)
	//insertData()
	updateData()
	queryMany(0)
}
