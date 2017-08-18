package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func divNums(a, b int) float64 {
	return float64(a) / float64(b)
}

func main() {
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
		log.Fatalf("file %q scanning error: %s", fileName, err)
	}

	ret := []string{}
	for line := range m {
		ret = append(ret, line)
	}

	return ret
}
