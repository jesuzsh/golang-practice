package lengthconv

// FToM converts feet to meters
func FToM(f Foot) Meter { return Meter(f * 0.3048) }

// MToF converts meters to feet
func MToF(m Meter) Foot { return Foot(m * 3.28084) }
