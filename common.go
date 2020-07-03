package glmatrix

import "math"

// Epsilon is a tolerant value
const Epsilon = 0.000001

const degree = math.Pi / 180

// ToRadian convert Degree To Radian
func ToRadian(a float64) float64 {
	return a * degree
}

func equals(a, b float64) bool {
	return math.Abs(a-b) <= Epsilon*math.Max(1.0, math.Max(math.Abs(a), math.Abs(b)))
}

func hypot(vals ...float64) float64 {
	sum := 0.
	for _, val := range vals {
		sum += val * val
	}
	return math.Sqrt(sum)
}
