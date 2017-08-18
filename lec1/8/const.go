package main

import "fmt"

func main() {
	const s = "123"
	//s = "456"

	const (
		a = 1
		b = 4815162342
		c = "x"
	)

	const (
		Sunday int = iota // 1 << iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Monday)
}
