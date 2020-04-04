package main

import (
	"fmt"
	"sync"
)

//互斥锁
//控制共享资源访问的方法，能够保证同时只有一个goroutine可以访问共享资源

var x int

var wg sync.WaitGroup
var lock sync.Mutex

func sum() {
	for i := 0; i < 100000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go sum()
	go sum()
	wg.Wait()
	fmt.Println(x)

}
