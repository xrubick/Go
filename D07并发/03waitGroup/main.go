package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup

// math/rand
func f1() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		r1 := rand.Int()    //int64
		r2 := rand.Intn(10) //0<=x<10
		fmt.Printf("%T %d %T %d\n", r1, r1, r2, r2)
	}
}

func f2(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

var wg sync.WaitGroup //相当于计数器

func main() {
	//f1()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f2(i)
	}
	wg.Wait() //等待至wg计数器降至为0
}
