package main

import "testing"

func TestDivNums(t *testing.T) {
	if divNums(1, 2) != 0.5 {
		t.Fatalf("invalid division: %f", divNums(1, 2))
	}
}
