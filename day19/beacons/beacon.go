package beacons

type Beacon struct {
	Position Point
}

func CreateBeacon(position Point) Beacon {
	return Beacon{position}
}

func (b Beacon) Eq(other Beacon) bool {
	return b.Position.Eq(other.Position)
}

func (b Beacon) RotateAndTranspose(angleX, angleY, angleZ int, p Point) Beacon {
	return Beacon{b.Position.RotateAndTranspose(angleX, angleY, angleZ, p)}
}

func (b Beacon) Rotate(angleX, angleY, angleZ int) Beacon {
	return Beacon{b.Position.Rotate(angleX, angleY, angleZ)}
}

func (b Beacon) Transpose(p Point) Beacon {
	return Beacon{b.Position.Transpose(p)}
}

func (b Beacon) RotateX(angle int) Beacon {
	return Beacon{b.Position.RotateX(angle)}
}

func (b Beacon) RotateY(angle int) Beacon {
	return Beacon{b.Position.RotateY(angle)}
}

func (b Beacon) RotateZ(angle int) Beacon {
	return Beacon{b.Position.RotateZ(angle)}
}

func (b Beacon) AssumeToBe(other Beacon) Point {
	return b.Position.AssumeToBe(other.Position)
}
