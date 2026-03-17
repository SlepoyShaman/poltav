package methods

import (
	"lab1/dynamic_models"
	"lab1/vectors"
)

type RungeKutte struct {
	h     float64
	model dynamic_models.DynamicModel
}

func NewRungeKutte(h float64, model dynamic_models.DynamicModel) *RungeKutte {
	return &RungeKutte{
		h:     h,
		model: model,
	}
}

func (r *RungeKutte) OneStep(start vectors.StateVector, t float64) vectors.StateVector {
	k1 := r.k1(t, start)
	k2 := r.k2(t, start, r.k2(t, start, k1))
	k3 := r.k3(t, start, k2)
	k4 := r.k4(t, start, k3)
	ksum := vectors.Add(k1, vectors.Multiply(2, k2))
	ksum = vectors.Add(ksum, vectors.Multiply(2, k3))
	ksum = vectors.Add(ksum, k4)
	return vectors.Add(start, vectors.Multiply(r.h/6, ksum))
}

func (r *RungeKutte) k1(t float64, vector vectors.StateVector) vectors.StateVector {
	return r.model.Funcs(t, vector)
}

func (r *RungeKutte) k2(t float64, vector vectors.StateVector, k1 vectors.StateVector) vectors.StateVector {
	deltaVector := vectors.Add(vector, vectors.Multiply(r.h/2, k1))
	return r.model.Funcs(t+r.h/2, deltaVector)
}

func (r *RungeKutte) k3(t float64, vector vectors.StateVector, k2 vectors.StateVector) vectors.StateVector {
	deltaVector := vectors.Add(vector, vectors.Multiply(r.h/2, k2))
	return r.model.Funcs(t+r.h/2, deltaVector)
}

func (r *RungeKutte) k4(t float64, vector vectors.StateVector, k3 vectors.StateVector) vectors.StateVector {
	deltaVector := vectors.Add(vector, vectors.Multiply(r.h, k3))
	return r.model.Funcs(t+r.h, deltaVector)
}
