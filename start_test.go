package main

import (
	"testing"
)

func TestsumGamePoint (t *testing.T) {
	playerTotal := 10
	computerTotal := 4
	total := playerTotal + computerTotal
	t.Log(total)
}