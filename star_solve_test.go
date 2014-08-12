package main

import (
	"testing"
)

func TestIt(t *testing.T) {
	fillCache(1, 10)
	starSolve()
}

func BenchIt(b *testing.B) {
	starSolve()
}
