package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		panic("invalid usage")
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("invalid usage: arg 1 is not a number: %s", err))
	}

	nums := genRandomNums(n)

	fmt.Println(nums)
}

func genRandomNums(n int) []int {
	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, 1)
	}

	return nums
}

func genRandomNumsOptimized(n int) []int {
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, 1)
	}

	return nums
}
