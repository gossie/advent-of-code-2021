package main

import (
	"testing"
)

func TestDetermineNumberOfLargerMeasurements(t *testing.T) {
	numberOfRaises := determineNumberOfLargerMeasurements()
	if numberOfRaises != 1215 {
		t.Fatalf("numberOfRaises = %v", numberOfRaises)
	}
}

func TestDetermineNumberOfLargerSumedMeasurements(t *testing.T) {
	numberOfRaises := determineNumberOfLargerSumedMeasurements()
	if numberOfRaises != 1150 {
		t.Fatalf("numberOfRaises = %v", numberOfRaises)
	}
}
