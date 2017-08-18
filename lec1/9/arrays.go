package main

import "fmt"

func main() {
	a := [3]int{1, 3, 2}
	a = [...]int{1, 3, 2}
	const n = 3
	a = [n]int{1, 3, 2}
	fmt.Println(a)
}
