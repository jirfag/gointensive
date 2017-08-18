package main

import (
	"os"
	"strconv"

	"github.com/jirfag/gointensive/lec3/3_production/2_apig/db"
	"github.com/jirfag/gointensive/lec3/3_production/2_apig/server"
)

// main ...
func main() {
	database := db.Connect()
	s := server.Setup(database)
	port := "8080"

	if p := os.Getenv("PORT"); p != "" {
		if _, err := strconv.Atoi(p); err == nil {
			port = p
		}
	}

	s.Run(":" + port)
}
