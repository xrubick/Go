package main

//事务处理:一个事务对应一个完整的业务,需要执行多次的DML(insert、update、delete)语句共同联合完成.
//ACID：
//原子性（Atomicity）
//一致性（Consistency）
//隔离性（Isolation）
//持久性（Durability）
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	//github.com/jmoiron/sqlx(操作数据库的第三方包)
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

func transaction() {
	//开启事务
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback() //错误回滚
		}
		fmt.Println("开启事务失败:", err)
		return
	}
	sql1 := `update user set age=age-1 where id=?`
	_, err = tx.Exec(sql1, 1)
	if err != nil {
		tx.Rollback()
		fmt.Println("执行sql1失败:", err)
		return
	}
	sql2 := `update user set age=age+1 where id=?`
	_, err = tx.Exec(sql2, 2)
	if err != nil {
		tx.Rollback()
		fmt.Println("执行sql2失败:", err)
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("提交失败:", err)
		return
	}
	fmt.Println("事务执行成功")
}

func main() {
	initDB()
	transaction()
}
