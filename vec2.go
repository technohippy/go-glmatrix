package glmatrix

import (
	"fmt"
	"math"
	"math/rand"
)

// NewVec2 creates a new, empty vec2
func NewVec2() []float64 {
	return []float64{0., 0.}
}

// Vec2Create creates a new vec2 initialized with values from an existing vector
func Vec2Create() []float64 {
	return NewVec2()
}

// Vec2Clone creates a new vec2 initialized with the given values
func Vec2Clone(a []float64) []float64 {
	return []float64{a[0], a[1]}
}

// Vec2FromValues creates a new vec2 initialized with the given values
func Vec2FromValues(x, y float64) []float64 {
	return []float64{x, y}
}

// Vec2Copy copy the values from one vec2 to another
func Vec2Copy(out, a []float64) []float64 {
	out[0] = a[0]
	out[1] = a[1]
	return out
}

// Vec2Set set the components of a vec2 to the given values
func Vec2Set(out []float64, x, y float64) []float64 {
	out[0] = x
	out[1] = y
	return out
}

// Vec2Add adds two vec2's
func Vec2Add(out, a, b []float64) []float64 {
	out[0] = a[0] + b[0]
	out[1] = a[1] + b[1]
	return out
}

// Vec2Subtract subtracts vector b from vector a
func Vec2Subtract(out, a, b []float64) []float64 {
	out[0] = a[0] - b[0]
	out[1] = a[1] - b[1]
	return out
}

// Vec2Multiply multiplies two vec2's
func Vec2Multiply(out, a, b []float64) []float64 {
	out[0] = a[0] * b[0]
	out[1] = a[1] * b[1]
	return out
}

// Vec2Divide divides two vec2's
func Vec2Divide(out, a, b []float64) []float64 {
	out[0] = a[0] / b[0]
	out[1] = a[1] / b[1]
	return out
}

// Vec2Ceil math.ceil the components of a vec2
func Vec2Ceil(out, a []float64) []float64 {
	out[0] = math.Ceil(a[0])
	out[1] = math.Ceil(a[1])
	return out
}

// Vec2Floor math.floor the components of a vec2
func Vec2Floor(out, a []float64) []float64 {
	out[0] = math.Floor(a[0])
	out[1] = math.Floor(a[1])
	return out
}

// Vec2Min returns the minimum of two vec2's
func Vec2Min(out, a, b []float64) []float64 {
	out[0] = math.Min(a[0], b[0])
	out[1] = math.Min(a[1], b[1])
	return out
}

// Vec2Max returns the maximum of two vec2's
func Vec2Max(out, a, b []float64) []float64 {
	out[0] = math.Max(a[0], b[0])
	out[1] = math.Max(a[1], b[1])
	return out
}

// Vec2Round math.round the components of a vec2
func Vec2Round(out, a []float64) []float64 {
	out[0] = math.Round(a[0])
	out[1] = math.Round(a[1])
	return out
}

// Vec2Scale scales a vec2 by a scalar number
func Vec2Scale(out, a []float64, scale float64) []float64 {
	out[0] = a[0] * scale
	out[1] = a[1] * scale
	return out
}

// Vec2ScaleAndAdd adds two vec2's after scaling the second operand by a scalar value
func Vec2ScaleAndAdd(out, a, b []float64, scale float64) []float64 {
	out[0] = a[0] + b[0]*scale
	out[1] = a[1] + b[1]*scale
	return out
}

// Vec2Distance calculates the euclidian distance between two vec2's
func Vec2Distance(a, b []float64) float64 {
	x := b[0] - a[0]
	y := b[1] - a[1]
	return math.Hypot(x, y)
}

// Vec2SquaredDistance calculates the squared euclidian distance between two vec2's
func Vec2SquaredDistance(a, b []float64) float64 {
	x := b[0] - a[0]
	y := b[1] - a[1]
	return x*x + y*y
}

// Vec2Length calculates the length of a vec2
func Vec2Length(out []float64) float64 {
	x := out[0]
	y := out[1]
	return math.Hypot(x, y)
}

// Vec2SquaredLength calculates the squared length of a vec2
func Vec2SquaredLength(out []float64) float64 {
	x := out[0]
	y := out[1]
	return x*x + y*y
}

// Vec2Negate negates the components of a vec2
func Vec2Negate(out, a []float64) []float64 {
	out[0] = -a[0]
	out[1] = -a[1]
	return out
}

// Vec2Inverse returns the inverse of the components of a vec2
func Vec2Inverse(out, a []float64) []float64 {
	out[0] = 1. / a[0]
	out[1] = 1. / a[1]
	return out
}

// Vec2Normalize normalize a vec2
func Vec2Normalize(out, a []float64) []float64 {
	len := Vec2Length(a)
	if 0 < len {
		len = 1. / len
	}
	out[0] = a[0] * len
	out[1] = a[1] * len
	return out
}

// Vec2Dot calculates the dot product of two vec2's
func Vec2Dot(a, b []float64) float64 {
	return a[0]*b[0] + a[1]*b[1]
}

// Vec2Cross computes the cross product of two vec2's
// Note that the cross product must by definition produce a 3D vector
func Vec2Cross(out, a, b []float64) []float64 {
	z := a[0]*b[1] - a[1]*b[0]
	out[0] = 0
	out[1] = 0
	out[2] = z
	return out
}

// Vec2Lerp performs a linear interpolation between two vec2's
func Vec2Lerp(out, a, b []float64, t float64) []float64 {
	ax := a[0]
	ay := a[1]
	out[0] = ax + t*(b[0]-ax)
	out[1] = ay + t*(b[1]-ay)
	return out
}

// Vec2Random generates a random vector with the given scale
func Vec2Random(out []float64, scale float64) []float64 {
	r := rand.Float64() * 2.0 * math.Pi
	out[0] = math.Cos(r) * scale
	out[1] = math.Sin(r) * scale
	return out
}

// Vec2TransformMat2 transforms the vec2 with a mat2
func Vec2TransformMat2(out, a, m []float64) []float64 {
	x := a[0]
	y := a[1]
	out[0] = m[0]*x + m[2]*y
	out[1] = m[1]*x + m[3]*y
	return out
}

// Vec2TransformMat2d transforms the vec2 with a mat2d
func Vec2TransformMat2d(out, a, m []float64) []float64 {
	x := a[0]
	y := a[1]
	out[0] = m[0]*x + m[2]*y + m[4]
	out[1] = m[1]*x + m[3]*y + m[5]
	return out
}

// Vec2TransformMat3 transforms the vec2 with a mat3
// 3rd vector component is implicitly '1'
func Vec2TransformMat3(out, a, m []float64) []float64 {
	x := a[0]
	y := a[1]
	out[0] = m[0]*x + m[2]*y + m[6]
	out[1] = m[1]*x + m[3]*y + m[7]
	return out
}

// Vec2TransformMat4 transforms the vec2 with a mat4
// 3rd vector component is implicitly '0'
// 4th vector component is implicitly '1'
func Vec2TransformMat4(out, a, m []float64) []float64 {
	x := a[0]
	y := a[1]
	out[0] = m[0]*x + m[2]*y + m[12]
	out[1] = m[1]*x + m[3]*y + m[13]
	return out
}

// Vec2Rotate rotate a 2D vector
func Vec2Rotate(out, p, c []float64, rad float64) []float64 {
	p0 := p[0] - c[0]
	p1 := p[1] - c[1]
	sinC := math.Sin(rad)
	cosC := math.Cos(rad)

	out[0] = p0*cosC - p1*sinC + c[0]
	out[1] = p0*sinC + p1*cosC + c[1]
	return out
}

// Vec2Angle get the angle between two 2D vectors
func Vec2Angle(a, b []float64) float64 {
	x1 := a[0]
	y1 := a[1]
	x2 := b[0]
	y2 := b[1]
	cosine := math.Sqrt(x1*x1+y1*y1) * math.Sqrt(x2*x2+y2*y2)
	if cosine != 0 {
		cosine = (x1*x2 + y1*y2) / cosine
	}
	return math.Acos(math.Min(math.Max(cosine, -1), 1))
}

// Vec2Zero set the components of a vec2 to zero
func Vec2Zero(out []float64) []float64 {
	out[0] = 0.
	out[1] = 0.
	return out
}

// Vec2Str returns a string representation of a vector
func Vec2Str(out []float64) string {
	return fmt.Sprintf("vec2(%v, %v)", out[0], out[1])
}

// Vec2ExactEquals returns whether or not the vectors exactly have the same elements in the same position (when compared with ===)
func Vec2ExactEquals(a, b []float64) bool {
	return a[0] == b[0] && a[1] == b[1]
}

// Vec2Equals returns whether or not the vectors have approximately the same elements in the same position.
func Vec2Equals(a, b []float64) bool {
	return equals(a[0], b[0]) && equals(a[1], b[1])
}

// Vec2Len alias for Vec2Length
var Vec2Len = Vec2Length

// Vec2Sub alias for Vec2Subtract
var Vec2Sub = Vec2Subtract

// Vec2Mul alias for Vec2Multiply
var Vec2Mul = Vec2Multiply

// Vec2Div alias for Vec2Divide
var Vec2Div = Vec2Divide

// Vec2Dist alias for Vec2Distance
var Vec2Dist = Vec2Distance

// Vec2SqrDist alias for Vec2SquaredDistance
var Vec2SqrDist = Vec2SquaredDistance

// Vec2SqrLen alias for Vec2SquaredLength
var Vec2SqrLen = Vec2SquaredLength

// Vec2ForEach perform some operation over an array of vec2s.
func Vec2ForEach(a []float64, stride, offset, count int, fn func([]float64, []float64, []float64), arg []float64) []float64 {
	if stride <= 0 {
		stride = 2
	}
	if offset <= 0 {
		offset = 0
	}
	var l int
	if 0 < count {
		l = int(math.Min(float64(count*stride+offset), float64(len(a))))
	} else {
		l = len(a)
	}

	for i := offset; i < l; i += stride {
		vec := []float64{a[i], a[i+1]}
		fn(vec, vec, arg)
		a[i] = vec[0]
		a[i+1] = vec[1]
	}
	return a
}
