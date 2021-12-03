package main

import (
	"testing"
)

func TestPowerConsumption(t *testing.T) {
	powerConsumption := powerConsumption()
	if powerConsumption != 3882564 {
		t.Fatalf("product = %v", powerConsumption)
	}
}

func TestLifeSupport(t *testing.T) {
	lifeSupport := lifeSupport()
	if lifeSupport != 3385170 {
		t.Fatalf("lifeSupport = %v", lifeSupport)
	}
}
