package beacons

import "strconv"

type Point struct {
	X, Y, Z int
}

func CreatePoint(x, y, z int) Point {
	return Point{x, y, z}
}

func (p Point) Eq(other Point) bool {
	return p.X == other.X && p.Y == other.Y && p.Z == other.Z
}

func (p Point) RotateAndTranspose(angleX, angleY, angleZ int, delta Point) Point {
	return p.RotateX(angleX).RotateY(angleY).RotateZ(angleZ).Transpose(delta)
}

func (p Point) Rotate(angleX, angleY, angleZ int) Point {
	return p.RotateX(angleX).RotateY(angleY).RotateZ(angleZ)
}

func (p Point) Transpose(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y, p.Z + other.Z}
}

func (p Point) RotateX(angle int) Point {
	if angle == 0 {
		return Point{p.X, p.Y, p.Z}
	} else if angle == 90 {
		return Point{p.X, p.Z, -p.Y}
	} else if angle == 180 {
		return Point{p.X, -p.Y, -p.Z}
	} else if angle == 270 {
		return Point{p.X, -p.Z, p.Y}
	}
	panic("rotation " + strconv.Itoa(angle) + " not allowed")
}

func (p Point) RotateY(angle int) Point {
	if angle == 0 {
		return Point{p.X, p.Y, p.Z}
	} else if angle == 90 {
		return Point{-p.Z, p.Y, p.X}
	} else if angle == 180 {
		return Point{-p.X, p.Y, -p.Z}
	} else if angle == 270 {
		return Point{p.Z, p.Y, -p.X}
	}
	panic("rotation " + strconv.Itoa(angle) + " not allowed")
}

func (p Point) RotateZ(angle int) Point {
	if angle == 0 {
		return Point{p.X, p.Y, p.Z}
	} else if angle == 90 {
		return Point{-p.Y, p.X, p.Z}
	} else if angle == 180 {
		return Point{-p.X, -p.Y, p.Z}
	} else if angle == 270 {
		return Point{p.Y, -p.X, p.Z}
	}
	panic("rotation " + strconv.Itoa(angle) + " not allowed")
}

func (p Point) AssumeToBe(other Point) Point {
	return Point{other.X - p.X, other.Y - p.Y, other.Z - p.Z}
}

func (p Point) Print() string {
	return "x: " + strconv.Itoa(p.X) + " y: " + strconv.Itoa(p.Y) + " z: " + strconv.Itoa(p.Z)
}
