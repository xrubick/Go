package main

import "fmt"
import "time"

func testTime() {

	//当前时间
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	//time.Unix
	ret := time.Unix(1585362290, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())

	//时间间隔
	fmt.Println(time.Second)

	//now + 10min
	fmt.Println(now.Add(10 * time.Minute))

	//时间格式化:转换成字符串类型 以go语言诞生时间2006年 01月 02日 15时 04分为准
	fmt.Println(now.Format("2006-01-02 03:04:05PM"))

	//按照对应格式，解析字符串类型时间

	timeObj, err := time.Parse("2006-01-02", "2020-03-30")
	if err != nil {
		fmt.Printf("wrong:%v", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
	fmt.Println("----------------------------")

	//sub
	d := now.Sub(timeObj)
	fmt.Println(d)

	//sleep 
	var n = 6
	fmt.Println("sleep")
	//time.Sleep(6 * time.Second)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("sleep 6s")

}

func main() {
	testTime()

}
