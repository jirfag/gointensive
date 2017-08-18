package main

import "fmt"

func main() {
	//var w string = "world"
	w := "world"

	var i, j int
	i = 0
	j = 0

	b := true

	//i := 1 // `=` vs `:=`

	fmt.Printf("hello, %s; i=%d, j=%d, b=%t!\n", w, i, j, b)
}
