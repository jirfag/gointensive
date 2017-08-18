package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/context"
)

var (
	wg sync.WaitGroup
)

// main is not changed
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	fmt.Println("Hey, I'm going to do some work")

	wg.Add(1)
	go work(ctx)
	wg.Wait()

	fmt.Println("Finished. I'm going home")

}

func work(ctx context.Context) {
	defer wg.Done()

	req, _ := http.NewRequest("GET", "http://localhost:1111", nil)
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.Printf("can't execute http request: %s", err)
		return
	}
	fmt.Println("Doing http request is a hard job")

	defer resp.Body.Close()
	out, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Server Response: %s\n", out)
}
