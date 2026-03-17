package integrators

import "lab1/vectors"

type Integrator interface {
	MoveTo(t float64) vectors.StateVector
}
