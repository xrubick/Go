package main

import "time"
import "github.com/xrubick/Go/D06/mylogger"

//日志需求
//1.往不同的地方输出日志
//2.日志级别：debug\trace\info\warning\error\fatal
//3.日志支持开关控制
//4.日志记录包含时间、行号、文件名、日志级别、日志信息
//5.日志切割

//log 日志打印

func main() {

	// //往文件里打印
	// logfile, err := os.OpenFile("./test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Printf("wrong:%v", err)
	// 	return
	// }
	// log.SetOutput(logfile)

	// //终端打印日志
	// for {
	// 	log.Println("just for test ")
	// 	time.Sleep(5 * time.Second)
	// }

	//mylogger/mylogger.go
	log := mylogger.NewLog("debug")
	for {
		log.Debug("test log :Debug")
		log.Info("test log :Info")
		log.Error("test log :Error")
		time.Sleep(5 * time.Second)
	}

	// // runtime.Caller()
	// pc, file, lineNum, ok := runtime.Caller(0) //0表示第一层
	// if !ok {
	// 	fmt.Println("runtime.Caller() failed\n")
	// 	return
	// }
	// funcName := runtime.FuncForPC(pc).Name()
	// fmt.Println(funcName)
	// fmt.Println(file)
	// fmt.Println(lineNum)

	//下一步日志写入文件、切割

}
