package day1

import (
	"testing"
)

func TestDetermineNumberOfLargerMeasurements(t *testing.T) {
	numberOfRaises := NumberOfLargerMeasurements("day1.txt")
	if numberOfRaises != 1215 {
		t.Fatalf("numberOfRaises = %v", numberOfRaises)
	}
}

func TestDetermineNumberOfLargerSumedMeasurements(t *testing.T) {
	numberOfRaises := NumberOfLargerSumedMeasurements("day1.txt")
	if numberOfRaises != 1150 {
		t.Fatalf("numberOfRaises = %v", numberOfRaises)
	}
}
