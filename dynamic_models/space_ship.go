package dynamic_models

import "lab1/vectors"

type SpaceShip struct{}

func (s *SpaceShip) Funcs(t float64, vector vectors.StateVector) vectors.StateVector {
	acceleration := vectors.AccelerationFromCoords(vector.X, vector.Y, vector.Z)
	return vectors.StateVector{
		X:  t * vector.Vx,
		Y:  t * vector.Vy,
		Z:  t * vector.Vz,
		Vx: t * acceleration.A1x,
		Vy: t * acceleration.A1y,
		Vz: t * acceleration.A1z,
	}
}
