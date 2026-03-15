package vectors

type StateVector struct {
	X, Y, Z, Vx, Vy, Vz float64
}

func Add(first, second StateVector) StateVector {
	return StateVector{
		X:  first.X + second.X,
		Y:  first.Y + second.Y,
		Z:  first.Z + second.Z,
		Vx: first.Vx + second.Vx,
		Vy: first.Vy + second.Vy,
		Vz: first.Vz + second.Vz,
	}
}

func Multiply(value float64, vector StateVector) StateVector {
	return StateVector{
		X:  vector.X * value,
		Y:  vector.Y * value,
		Z:  vector.Z * value,
		Vx: vector.Vx * value,
		Vy: vector.Vy * value,
		Vz: vector.Vz * value,
	}
}
