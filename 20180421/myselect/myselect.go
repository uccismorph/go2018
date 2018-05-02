package main

import "fmt"

var g_num int = 0

func main() {
	ch := make(chan int, 1024)
	for idx := 0; idx < 500; idx += 1 {
		select {
		case ch <- idx:
		case ch <- 1:
		}
	}

	for idx := 0; idx < 500; idx += 1 {
		res, ok := <-ch

		if ok {
			fmt.Println("We get:", res, "at:", idx)
		} else {
			fmt.Println("channel closed")
		}
	}

}
