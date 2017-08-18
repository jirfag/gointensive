package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"./subpackage"
)

type lineCounter struct {
	line  string
	count int
}

func newLineCounter(line string, count int) *lineCounter {
	return &lineCounter{
		line:  line,
		count: count,
	}
}

func main() {
	//s := lineCounter{"a", 1}
	/*s := lineCounter{
		line:  "a",
		count: 1,
	}*/
	//s := newLineCounter("a", 1)

	//s := subpackage.testStruct{}

	lines := getUniqLinesFromFile(os.Args[1])
	fmt.Println(lines)
}

func getUniqLinesFromFile(fileName string) []lineCounter {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can't open file %q: %s", fileName, err)
	}
	defer file.Close()

	m := map[string]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("file %q scanning error: %s", fileName, err)
	}

	ret := []lineCounter{}
	for line := range m {
		lc := lineCounter{
			line:  line,
			count: m[line],
		}
		ret = append(ret, lc)
	}

	return ret
}
