package server

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jirfag/gointensive/lec3/2_http_pool/workers"
)

type IPool interface {
	Size() int
	Run()
	AddTaskSyncTimed(f workers.Func, timeout time.Duration) (interface{}, error)
}

var wp IPool = workers.NewPool(5)

func init() {
	wp.Run()
}

const requestWaitInQueueTimeout = time.Millisecond * 100

// ab -c10 -n20 localhost:8000/hello

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, err := wp.AddTaskSyncTimed(func() interface{} {
		time.Sleep(time.Millisecond * 500)
		io.WriteString(w, "Hello, world!")
		return nil
	}, requestWaitInQueueTimeout)

	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s!\n", err), 500)
	}
}

func RunHTTPServer(addr string) error {
	http.HandleFunc("/hello", rootHandler)
	return http.ListenAndServe(addr, nil)
}
