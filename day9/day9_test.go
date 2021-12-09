package day9

import "testing"

func TestRiskLevel(t *testing.T) {
	riskLevel := RiskLevel("day9.txt")
	if riskLevel != 439 {
		t.Fatalf("riskLevel = %v", riskLevel)
	}
}

func TestBasins(t *testing.T) {
	basin := Basin("day9.txt")
	if basin != 900900 {
		t.Fatalf("basin = %v", basin)
	}
}
