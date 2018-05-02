package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var goal int

func primetask(c chan int, id int) {
	p := <-c
	if p > goal {
		os.Exit(0)
	}

	fmt.Printf("p = %d in id [%d]\n", p, id)

	nc := make(chan int)

	go primetask(nc, id+1)

	for {
		i := <-c
		if i%p != 0 {
			fmt.Printf("find i [%d] in func id: %d\n", i, id)
			nc <- i
		}
	}
}

func main() {
	flag.Parse()

	args := flag.Args()
	if args != nil && len(args) > 0 {
		var err error
		goal, err = strconv.Atoi(args[0])
		if err != nil {
			goal = 100
		}
	} else {
		goal = 100
	}

	fmt.Println("goal =", goal)

	c := make(chan int)

	var func_id int = 0

	go primetask(c, func_id)

	for i := 2; ; i += 1 {
		c <- i
	}
}
