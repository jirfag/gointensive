package main

// Lazy and Very Random Server
import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", LazyServer)
	http.ListenAndServe(":1111", nil)
}

// sometimes really fast server, sometimes really slow server
func LazyServer(w http.ResponseWriter, req *http.Request) {
	headOrTails := rand.Intn(2)

	if headOrTails == 0 {
		time.Sleep(1 * time.Second)
		fmt.Fprintf(w, "Go! slow %v\n", headOrTails)
		fmt.Printf("Go! slow %v\n", headOrTails)
		return
	}

	fmt.Fprintf(w, "Go! quick %v\n", headOrTails)
	fmt.Printf("Go! quick %v\n", headOrTails)
	return
}
