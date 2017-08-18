package main

import "os"

func init() {
	if os.Getenv("SOME_KEY") == "" {
		panic("no SOME_KEY in ENV")
	}
}
