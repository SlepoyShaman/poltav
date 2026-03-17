package integrators

import (
	"image/color"
	"lab1/integrators/methods"
	"lab1/vectors"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const filename = "chart.png"

type StepIntegrator struct {
	h      float64
	start  vectors.StateVector
	method methods.Method

	vxPoints plotter.XYs
	vyPoints plotter.XYs
	vzPoints plotter.XYs
}

func NewIntegrator(h float64, method methods.Method, start vectors.StateVector) *StepIntegrator {
	return &StepIntegrator{
		h:      h,
		start:  start,
		method: method,

		vxPoints: make(plotter.XYs, 0),
		vyPoints: make(plotter.XYs, 0),
		vzPoints: make(plotter.XYs, 0),
	}
}

func (s *StepIntegrator) MoveTo(tk float64) vectors.StateVector {
	var t float64 = 0
	currenVector := s.start

	for t < tk {
		t += s.h
		if t >= tk {
			t = tk
		}

		currenVector = s.method.OneStep(currenVector, t)
		s.appendVectorOnSpeedChart(currenVector, t)
	}

	s.printSpeedChart(filename)
	return currenVector
}

func (s *StepIntegrator) appendVectorOnSpeedChart(vector vectors.StateVector, t float64) {
	s.vxPoints = append(s.vxPoints, plotter.XY{X: t, Y: vector.Vx})
	s.vyPoints = append(s.vyPoints, plotter.XY{X: t, Y: vector.Vy})
	s.vzPoints = append(s.vzPoints, plotter.XY{X: t, Y: vector.Vz})
}

func (s *StepIntegrator) printSpeedChart(filename string) error {
	p := plot.New()
	p.Title.Text = "Графики скоростей"
	p.X.Label.Text = "t"
	p.Y.Label.Text = "V"

	lineVx, err := plotter.NewLine(s.vxPoints)
	if err != nil {
		return err
	}
	lineVx.Color = color.RGBA{R: 255, A: 255} // Красный
	lineVx.LineStyle.Width = vg.Points(1)

	lineVy, err := plotter.NewLine(s.vyPoints)
	if err != nil {
		return err
	}
	lineVy.Color = color.RGBA{G: 255, A: 255} // Зеленый
	lineVy.LineStyle.Width = vg.Points(1)

	lineVz, err := plotter.NewLine(s.vzPoints)
	if err != nil {
		return err
	}
	lineVz.Color = color.RGBA{B: 255, A: 255} // Синий
	lineVz.LineStyle.Width = vg.Points(1)

	p.Add(lineVx, lineVy, lineVz)

	p.Legend.Add("Vx", lineVx)
	p.Legend.Add("Vy", lineVy)
	p.Legend.Add("Vz", lineVz)
	p.Legend.Top = true

	s.clearCharts()
	return p.Save(6*vg.Inch, 4*vg.Inch, filename)
}

func (s *StepIntegrator) clearCharts() {
	s.vxPoints = make(plotter.XYs, 0)
	s.vyPoints = make(plotter.XYs, 0)
	s.vzPoints = make(plotter.XYs, 0)
}
