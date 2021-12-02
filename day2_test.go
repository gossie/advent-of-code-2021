package main

import (
	"testing"
)

func TestPosition1(t *testing.T) {
	position := position1()
	if position != 1654760 {
		t.Fatalf("position = %v", position)
	}
}

func TestPosition2(t *testing.T) {
	position := position2()
	if position != 1956047400 {
		t.Fatalf("position = %v", position)
	}
}
