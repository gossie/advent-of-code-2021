package day16

import "testing"

func TestVersions(t *testing.T) {
	versions := Versions("day16.txt")
	if versions != 1014 {
		t.Fatalf("versions = %v", versions)
	}
}
