package main

import (
	"testing"
)

func TestDetermineNumberOfLargerMeasurements(t *testing.T) {
	numberOfRaises := numberOfLargerMeasurements()
	if numberOfRaises != 1215 {
		t.Fatalf("numberOfRaises = %v", numberOfRaises)
	}
}

func TestDetermineNumberOfLargerSumedMeasurements(t *testing.T) {
	numberOfRaises := numberOfLargerSumedMeasurements()
	if numberOfRaises != 1150 {
		t.Fatalf("numberOfRaises = %v", numberOfRaises)
	}
}
