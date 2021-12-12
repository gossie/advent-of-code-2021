package day12

import "testing"

func TestNumberOfPaths1(t *testing.T) {
	paths := NumberOfPaths("day12.txt", true)
	if paths != 5457 {
		t.Fatalf("sum = %v", paths)
	}
}

func TestNumberOfPaths2(t *testing.T) {
	paths := NumberOfPaths("day12.txt", false)
	if paths != 128506 {
		t.Fatalf("sum = %v", paths)
	}
}
