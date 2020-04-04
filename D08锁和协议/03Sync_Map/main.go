package main

import (
	"fmt"
	"strconv"
	"sync"
)

//内置的map不是并发安全

var m1 = make(map[string]int)

//取值
func get(key string) int {
	return m1[key]
}

//添加值
func set(key string, value int) {
	m1[key] = value
}

// func main() {
// 	wg := sync.WaitGroup{}

// 	for i := 0; i < 19; i++ {   //fatal error: concurrent map writes
// 		wg.Add(1)
// 		go func(n int) {
// 			key := strconv.Itoa(n)   //int to  array
// 			set(key, n)
// 			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

var wg sync.WaitGroup

var m2 sync.Map //sync包中提供了并发安全版map---sync.Map

func main() {

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {

			key := strconv.Itoa(n)

			m2.Store(key, n)         //存值
			value, _ := m2.Load(key) //取值
			fmt.Printf("key:%s  value:%d\n", key, value)
			wg.Done()

		}(i)
	}
	wg.Wait()

}
