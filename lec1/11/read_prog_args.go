package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		panic("invalid usage")
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("can't read arg1: %s", err)
	}

	fmt.Println(n)
}
