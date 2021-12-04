package day3

import (
	"testing"
)

func TestPowerConsumption(t *testing.T) {
	powerConsumption := PowerConsumption("day3.txt")
	if powerConsumption != 3882564 {
		t.Fatalf("product = %v", powerConsumption)
	}
}

func TestLifeSupport(t *testing.T) {
	lifeSupport := LifeSupport("day3.txt")
	if lifeSupport != 3385170 {
		t.Fatalf("lifeSupport = %v", lifeSupport)
	}
}
