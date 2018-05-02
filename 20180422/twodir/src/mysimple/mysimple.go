package main

import (
	"fmt"
	"libadd"
	"libsub"
)

func main() {
	num1 := 10
	num2 := 20
	fmt.Println(libadd.Add(num1, num2))
	fmt.Println(libsub.Sub(num1, num2))

}
