package vectors

import (
	"lab1/constants"
	"math"
)

type AccelerationVector struct {
	A1x, A1y, A1z float64
}

func AccelerationFromCoords(x, y, z float64) AccelerationVector {
	r := proccesRadius(x, y, z)
	return AccelerationVector{
		A1x: proccesForCoord(x, r),
		A1y: proccesForCoord(y, r),
		A1z: proccesForCoord(z, r),
	}
}

func proccesForCoord(value, r float64) float64 {
	numerator := -1 * constants.Gravity * value
	return numerator / math.Pow(r, 3)
}

func proccesRadius(x, y, z float64) float64 {
	return math.Sqrt(x*x + y*y + z*z)
}
