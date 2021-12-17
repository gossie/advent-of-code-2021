package day17

import "testing"

func TestHeighest(t *testing.T) {
	heighest := Heighest("day17.txt")
	if heighest != 4095 {
		t.Fatalf("heighest = %v", heighest)
	}
}
func TestAmount(t *testing.T) {
	amount := HowMany("day17.txt")
	if amount != 3773 {
		t.Fatalf("amount = %v", amount)
	}
}
