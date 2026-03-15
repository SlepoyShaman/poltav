package dynamic_models

import "lab1/vectors"

type DynamicModel interface {
	Funcs(t float64, vector vectors.StateVector) vectors.StateVector
}
