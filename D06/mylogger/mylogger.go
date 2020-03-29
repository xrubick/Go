package mylogger

import (
	"fmt"
	"time"
)

//结构体
type Logger struct {
	Level LogLevel
}

//构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

//日志开关
func (l Logger) enable(loglevel LogLevel) bool {
	return l.Level <= loglevel
}

func logString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	}
	return "DEBUG"
}

func log(lv LogLevel, msg string) {
	now := time.Now()
	funcName, fileName, lineNum := getLaw(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-1-2 15:04:05"), logString(lv), funcName, fileName, lineNum, msg)

}

//方法
func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		//log(DEBUG,msg)
		now := time.Now()
		funcName, fileName, lineNum := getLaw(2)
		fmt.Printf("[%s] [Debug] [%s:%s:%d] %s\n", now.Format("2006-1-2 15:04:05"), funcName, fileName, lineNum, msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		//fmt.Printf("[%s] [Info] %s\n", now.Format("2006-1-2 15:04:05"), msg)
		funcName, fileName, lineNum := getLaw(2)
		fmt.Printf("[%s] [Info] [%s:%s:%d] %s\n", now.Format("2006-1-2 15:04:05"), funcName, fileName, lineNum, msg)

	}
}
func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		now := time.Now()
		//fmt.Printf("[%s] [Warning] %s\n", now.Format("2006-1-2 15:04:05"), msg)
		funcName, fileName, lineNum := getLaw(2)
		fmt.Printf("[%s] [Warning] [%s:%s:%d] %s\n", now.Format("2006-1-2 15:04:05"), funcName, fileName, lineNum, msg)

	}
}
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		//fmt.Printf("[%s] [Error] %s\n", now.Format("2006-1-2 15:04:05"), msg)
		funcName, fileName, lineNum := getLaw(2)
		fmt.Printf("[%s] [Error] [%s:%s:%d] %s\n", now.Format("2006-1-2 15:04:05"), funcName, fileName, lineNum, msg)

	}
}
