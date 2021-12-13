package day13

import "testing"

func TestNumberOfPaths1(t *testing.T) {
	hashtags := AfterOneFold("day13.txt")
	if hashtags != 751 {
		t.Fatalf("hashtags = %v", hashtags)
	}
}
