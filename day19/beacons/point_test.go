package beacons

import "testing"

func TestRotateX90(t *testing.T) {
	point := CreatePoint(2, 1, 3)
	rotated := point.RotateX(90)
	if !rotated.Eq(CreatePoint(2, 3, -1)) {
		t.Fatalf("rotated = %v", rotated)
	}
}

func TestRotateX180(t *testing.T) {
	point := CreatePoint(2, 1, 3)
	rotated := point.RotateX(180)
	if !rotated.Eq(CreatePoint(2, -1, -3)) {
		t.Fatalf("rotated = %v", rotated)
	}
}

func TestRotateX270(t *testing.T) {
	point := CreatePoint(2, 1, 3)
	rotated := point.RotateX(270)
	if !rotated.Eq(CreatePoint(2, -3, 1)) {
		t.Fatalf("rotated = %v", rotated)
	}
}

func TestRotateY90(t *testing.T) {
	point := CreatePoint(2, 1, 3)
	rotated := point.RotateY(90)
	if !rotated.Eq(CreatePoint(-3, 1, 2)) {
		t.Fatalf("rotated = %v", rotated)
	}
}

func TestRotateY180(t *testing.T) {
	point := CreatePoint(2, 1, 3)
	rotated := point.RotateY(180)
	if !rotated.Eq(CreatePoint(-2, 1, -3)) {
		t.Fatalf("rotated = %v", rotated)
	}
}

func TestRotateY270(t *testing.T) {
	point := CreatePoint(2, 1, 3)
	rotated := point.RotateY(270)
	if !rotated.Eq(CreatePoint(3, 1, -2)) {
		t.Fatalf("rotated = %v", rotated)
	}
}

func TestRotateZ90(t *testing.T) {
	point := CreatePoint(2, 1, 3)
	rotated := point.RotateZ(90)
	if !rotated.Eq(CreatePoint(-1, 2, 3)) {
		t.Fatalf("rotated = %v", rotated)
	}
}

func TestRotateZ180(t *testing.T) {
	point := CreatePoint(2, 1, 3)
	rotated := point.RotateZ(180)
	if !rotated.Eq(CreatePoint(-2, -1, 3)) {
		t.Fatalf("rotated = %v", rotated)
	}
}

func TestRotateZ270(t *testing.T) {
	point := CreatePoint(2, 1, 3)
	rotated := point.RotateZ(270)
	if !rotated.Eq(CreatePoint(1, -2, 3)) {
		t.Fatalf("rotated = %v", rotated)
	}
}

func TestAssumeToBe1(t *testing.T) {
	point1 := CreatePoint(-618, -824, -621)
	point2 := CreatePoint(-686, 422, -578)
	expected := CreatePoint(68, -1246, -43)
	actual := point2.AssumeToBe(point1)
	if !actual.Eq(expected) {
		t.Fatalf("delta = %v", actual)
	}
}

func TestAssumeToBe2(t *testing.T) {
	point1 := CreatePoint(-537, -823, -458)
	point2 := CreatePoint(-605, 423, -415)
	expected := CreatePoint(68, -1246, -43)
	actual := point2.AssumeToBe(point1)
	if !actual.Eq(expected) {
		t.Fatalf("delta = %v", actual)
	}
}

func TestTranspose(t *testing.T) {
	expected := CreatePoint(-537, -823, -458)
	point1 := CreatePoint(-605, 423, -415)
	delta := CreatePoint(68, -1246, -43)
	actual := point1.Transpose(delta)
	if !actual.Eq(expected) {
		t.Fatalf("new point = %v", actual)
	}
}
