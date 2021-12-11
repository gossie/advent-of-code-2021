package day11

import "testing"

func TestNumberOfFlashes(t *testing.T) {
	sum := NumberOfFlashes("day11.txt")
	if sum != 1640 {
		t.Fatalf("sum = %v", sum)
	}
}

func TestStepWhenAllFlash(t *testing.T) {
	step := StepWhenAllFlash("day11.txt")
	if step != 312 {
		t.Fatalf("step = %v", step)
	}
}
