package day8

import "testing"

func TestNumberOf1And4And7And8(t *testing.T) {
	count := NumberOf1And4And7And8("day8.txt")
	if count != 239 {
		t.Fatalf("count = %v", count)
	}
}

func TestOutputSum(t *testing.T) {
	count := OutputSum("day8.txt")
	if count != 946346 {
		t.Fatalf("count = %v", count)
	}
}
