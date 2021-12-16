package day16

import "testing"

func TestVersions(t *testing.T) {
	versions := Versions("day16.txt")
	if versions != 1014 {
		t.Fatalf("versions = %v", versions)
	}
}

func TestCalculate(t *testing.T) {
	calculationResult := Calculate("day16.txt")
	if calculationResult != 1922490999789 {
		t.Fatalf("calculationResult = %v", calculationResult)
	}
}
