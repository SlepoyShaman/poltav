package methods

import (
	"lab1/dynamic_models"
	"lab1/vectors"
)

type Euler struct {
	h     float64
	model dynamic_models.DynamicModel
}

func NewEuler(h float64, model dynamic_models.DynamicModel) *Euler {
	return &Euler{
		h:     h,
		model: model,
	}
}

func (e *Euler) OneStep(start vectors.StateVector, t float64) vectors.StateVector {
	delta := vectors.Multiply(e.h, e.model.Funcs(t, start))
	return vectors.Add(start, delta)
}
