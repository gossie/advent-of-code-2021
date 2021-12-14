package day14

import "testing"

func TestQuantities(t *testing.T) {
	result := Quantities("day14.txt", 10)
	if result != 2509 {
		t.Fatalf("result = %v", result)
	}
}
