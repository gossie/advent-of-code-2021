package day7

import (
	"testing"
)

func TestFuelConstant(t *testing.T) {
	fuel := FuelConstantConsumptions("day7.txt")
	if fuel != 342641 {
		t.Fatalf("fuel = %v", fuel)
	}
}

func TestFuelLinear(t *testing.T) {
	fuel := FuelLinearConsumption("day7.txt")
	if fuel != 93006301 {
		t.Fatalf("fuel = %v", fuel)
	}
}
