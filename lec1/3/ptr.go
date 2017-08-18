package main

import "fmt"

func main() {
	i := 1
	iPtr := &i
	fmt.Printf("i=%d, *iPtr=%d\n", i, *iPtr)

	*iPtr = 2
	fmt.Printf("i=%d, *iPtr=%d\n", i, *iPtr)
}
