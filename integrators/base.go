package integrators

import (
	"lab1/integrators/methods"
)

type BaseIntegrator struct {
	tk, h  float64
	method methods.Method
}

func NewIntegrator(tk, h float64, method methods.Method) *BaseIntegrator {

	return &BaseIntegrator{
		tk:     tk,
		h:      h,
		method: method,
	}
}

func (i *BaseIntegrator) MoveTo(t float64) {
	panic("aaaa")
}
