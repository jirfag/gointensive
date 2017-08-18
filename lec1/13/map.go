package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	m := map[string]int{}

	m["a"] = 0

	m["b"]++
	m["b"] += 5

	//c, ok := m["a"]

	fmt.Printf("m is %v\n", m)

	lines := getUniqLinesFromFile(os.Args[1])
	fmt.Println(lines)
}

func getUniqLinesFromFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can't open file %q: %s", fileName, err)
	}
	defer file.Close()

	m := map[string]bool{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		//return err
	}

	ret := []string{}
	for line := range m {
		ret = append(ret, line)
	}

	return ret
}
