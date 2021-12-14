package day14

import "testing"

func TestQuantities(t *testing.T) {
	result := Quantities("day14.txt", 10)
	if result != 2509 {
		t.Fatalf("result = %v", result)
	}
}

func TestQuantitiesLarge(t *testing.T) {
	result := Quantities("day14.txt", 40)
	if result != 2827627697643 {
		t.Fatalf("result = %v", result)
	}
}
