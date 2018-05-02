package main

import "fmt"

func my_first_go_func(value int, msg string) (result int, err string) {
	return 0, msg
}

func main() {
	var hello_str string = "Hello World!"
	_, hello_err := my_first_go_func(0, hello_str)
	fmt.Println(hello_err)
	for i, ch := range hello_str {
		fmt.Println(i, ch)
	}
}
