package main

import (
	"log"

	"github.com/jirfag/gointensive/lec3/2_http_pool/server"
)

func main() {
	log.Fatal(server.RunHTTPServer(":8000"))
}
