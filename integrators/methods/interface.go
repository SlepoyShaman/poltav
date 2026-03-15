package methods

import "lab1/vectors"

type Method interface {
	OneStep(start vectors.StateVector, t float64) vectors.StateVector
}
