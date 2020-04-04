package main

import (
	"fmt"
	"sync"
	"time"
)

//读写互斥锁

var (
	x    int
	wg   sync.WaitGroup
	lock sync.Mutex   //互斥锁
	rwlock sync.RWMutex //读写互斥锁
)

//读
func read() {
	defer wg.Done()
	//lock.Lock()
	rwlock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwlock.RUnlock()

}

//写

func write() {
	defer wg.Done()
	//lock.Lock()
	rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 10)
	//lock.Unlock()
	rwlock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go read()

	}

	wg.Wait()
	fmt.Println(time.Now().Sub(start))

}
