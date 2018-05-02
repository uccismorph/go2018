package main

import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println("a less than 2")
	} else {
		fmt.Println("a larger than 2")
	}
}
