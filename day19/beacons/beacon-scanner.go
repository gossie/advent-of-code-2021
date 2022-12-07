package beacons

type BeaconScanner struct {
	Position Point
	Beacons  []Beacon
}

func CreateBeaconScanner(beacons []Beacon) BeaconScanner {
	return BeaconScanner{Beacons: beacons}
}

func (bs BeaconScanner) RotateAndTranspose(angleX, angleY, angleZ int, p Point) BeaconScanner {
	beacons := make([]Beacon, 0, len(bs.Beacons))
	for _, b := range bs.Beacons {
		beacons = append(beacons, b.RotateAndTranspose(angleX, angleY, angleZ, p))
	}
	return BeaconScanner{Beacons: beacons}
}

func (bs BeaconScanner) Rotate(angleX, angleY, angleZ int) BeaconScanner {
	beacons := make([]Beacon, 0, len(bs.Beacons))
	for _, b := range bs.Beacons {
		beacons = append(beacons, b.Rotate(angleX, angleY, angleZ))
	}
	return BeaconScanner{Beacons: beacons}
}

func (bs BeaconScanner) Transpose(p Point) BeaconScanner {
	beacons := make([]Beacon, 0, len(bs.Beacons))
	for _, b := range bs.Beacons {
		beacons = append(beacons, b.Transpose(p))
	}
	return BeaconScanner{Beacons: beacons}
}

func (bs BeaconScanner) RotateX(angle int) BeaconScanner {
	beacons := make([]Beacon, 0, len(bs.Beacons))
	for _, b := range bs.Beacons {
		beacons = append(beacons, b.RotateX(angle))
	}
	return BeaconScanner{Beacons: beacons}
}

func (bs BeaconScanner) RotateY(angle int) BeaconScanner {
	beacons := make([]Beacon, 0, len(bs.Beacons))
	for _, b := range bs.Beacons {
		beacons = append(beacons, b.RotateY(angle))
	}
	return BeaconScanner{Beacons: beacons}
}

func (bs BeaconScanner) RotateZ(angle int) BeaconScanner {
	beacons := make([]Beacon, 0, len(bs.Beacons))
	for _, b := range bs.Beacons {
		beacons = append(beacons, b.RotateZ(angle))
	}
	return BeaconScanner{Beacons: beacons}
}

func (bs BeaconScanner) ForEachBeacon(consumer func(b Beacon)) {
	for _, b := range bs.Beacons {
		consumer(b)
	}
}

func (bs BeaconScanner) Overlapping(other BeaconScanner) bool {
	sum := 0
	for _, b1 := range bs.Beacons {
		for _, b2 := range other.Beacons {
			if b1.Eq(b2) {
				sum++
			}
		}
	}
	// println("found ", sum, " overlapping points")
	return sum >= 12
}

func (bs *BeaconScanner) AddMissingBeacons(newBeacons []Beacon) {
	for _, newBeacon := range newBeacons {
		if !bs.contains(newBeacon) {
			bs.Beacons = append(bs.Beacons, newBeacon)
		}
	}
}

func (bs BeaconScanner) contains(newBeacon Beacon) bool {
	for _, b := range bs.Beacons {
		if b.Eq(newBeacon) {
			return true
		}
	}
	return false
}
