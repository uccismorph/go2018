package main

import "fmt"

func foo(a []int) {
	a[0], a[1] = a[1], a[0]
}

func main() {
	my_array := []int{1, 2, 3, 4, 5}
	foo(my_array)
	fmt.Println(my_array)
}
