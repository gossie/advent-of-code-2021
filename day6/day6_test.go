package day6

import (
	"testing"
)

func TestFishPopulation(t *testing.T) {
	population := FishPopulation("day6.txt", 256)
	if population != 1600306001288 {
		t.Fatalf("population = %v", population)
	}
}
