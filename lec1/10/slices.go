package main

import "fmt"

func main() {
	//a := [...]int{1, 3, 2, 4}
	//s := a[1:3]

	s1 := []int{7, 8, 9}
	reverse(s1)
	fmt.Println(s1)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {

		s[i], s[j] = s[j], s[i]
	}
}
