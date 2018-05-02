package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var g_num int = 0

func Count(ch chan int, lock *sync.Mutex) {
	for i := 0; i < 1000; i += 1 {
		lock.Lock()
		g_num += 1
		lock.Unlock()
	}
	ch <- 1
}

func main() {
	fmt.Printf("max cpu core: %d\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	lock := &sync.Mutex{}

	chs := make([]chan int, 5000)
	curr := time.Now().UnixNano()
	for i := 0; i < 5000; i += 1 {
		chs[i] = make(chan int)
		go Count(chs[i], lock)
	}

	for _, ch := range chs {
		//fmt.Println("get channel:", i)
		<-ch
	}
	curr = time.Now().UnixNano() - curr
	fmt.Println(curr / int64(time.Millisecond))
	//fmt.Println("Now gobal num:", g_num)

}
