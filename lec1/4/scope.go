package main

import "fmt"

func test() *int {
	var v int = 2
	return &v
}

func main() {
	a := 0
	if true {
		a := 1
		fmt.Printf("inner a = %d\n", a)
	}
	v := *test()
	fmt.Printf("a = %d, v = %d\n", a, v)
}
