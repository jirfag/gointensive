package main

import (
	"fmt"
	"log"
	"strconv"
)

type Value struct {
	i int
}

func main() {
	voidFunc()
	r := retIntFunc()
	r1, err := transformInt(r)
	if err != nil {
		panic(err)
	}
	log.Println(r1)

	v := Value{i: 7}
	passByValueOrPointer(v)
	fmt.Println(v)

	// what if add pointer to struct?
}

func voidFunc() {
	fmt.Println("hello")
}

func retIntFunc() int {
	return 1
}

func transformInt(v int) (int, error) {
	if v < 0 {
		return 0, fmt.Errorf("invalid value %d: below zero", v)
	}

	ret := v * 2
	return ret, nil
}

func passByValueOrPointer(v Value) {
	v.i++
}

func namedRetValues(c, d string) (a int, b int) {
	a, _ = strconv.Atoi(c)
	b, _ = strconv.Atoi(d)
	return
}
