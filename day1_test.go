package main

import (
	"testing"
)

func testDetermineNumberOfLargerMeasurements(t *testing.T) {
	numberOfRaises := determineNumberOfLargerMeasurements()
	if numberOfRaises != 1215 {
		t.Fatalf("numberOfRaises = %v", numberOfRaises)
	}
}

func testDetermineNumberOfLargerSumedMeasurements(t *testing.T) {
	numberOfRaises := determineNumberOfLargerSumedMeasurements()
	if numberOfRaises != 1150 {
		t.Fatalf("numberOfRaises = %v", numberOfRaises)
	}
}
