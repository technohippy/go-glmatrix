package gomatrix

import (
	"fmt"
	"math"
	"math/rand"
)

// NewVec2 .
func NewVec2() []float64 {
	return []float64{0., 0.}
}

// Vec2Create .
func Vec2Create() []float64 {
	return NewVec2()
}

// Vec2Clone .
func Vec2Clone(a []float64) []float64 {
	return []float64{a[0], a[1]}
}

// Vec2FromValues .
func Vec2FromValues(x, y float64) []float64 {
	return []float64{x, y}
}

// Vec2Copy .
func Vec2Copy(out, a []float64) []float64 {
	out[0] = a[0]
	out[1] = a[1]
	return out
}

// Vec2Set .
func Vec2Set(out []float64, x, y float64) []float64 {
	out[0] = x
	out[1] = y
	return out
}

// Vec2Add .
func Vec2Add(out, a, b []float64) []float64 {
	out[0] = a[0] + b[0]
	out[1] = a[1] + b[1]
	return out
}

// Vec2Subtract .
func Vec2Subtract(out, a, b []float64) []float64 {
	out[0] = a[0] - b[0]
	out[1] = a[1] - b[1]
	return out
}

// Vec2Multiply .
func Vec2Multiply(out, a, b []float64) []float64 {
	out[0] = a[0] * b[0]
	out[1] = a[1] * b[1]
	return out
}

// Vec2Divide .
func Vec2Divide(out, a, b []float64) []float64 {
	out[0] = a[0] / b[0]
	out[1] = a[1] / b[1]
	return out
}

// Vec2Ceil .
func Vec2Ceil(out, a []float64) []float64 {
	out[0] = math.Ceil(a[0])
	out[1] = math.Ceil(a[1])
	return out
}

// Vec2Floor .
func Vec2Floor(out, a []float64) []float64 {
	out[0] = math.Floor(a[0])
	out[1] = math.Floor(a[1])
	return out
}

// Vec2Min .
func Vec2Min(out, a, b []float64) []float64 {
	out[0] = math.Min(a[0], b[0])
	out[1] = math.Min(a[1], b[1])
	return out
}

// Vec2Max .
func Vec2Max(out, a, b []float64) []float64 {
	out[0] = math.Max(a[0], b[0])
	out[1] = math.Max(a[1], b[1])
	return out
}

// Vec2Round .
func Vec2Round(out, a []float64) []float64 {
	out[0] = math.Round(a[0])
	out[1] = math.Round(a[1])
	return out
}

// Vec2Scale .
func Vec2Scale(out, a []float64, scale float64) []float64 {
	out[0] = a[0] * scale
	out[1] = a[1] * scale
	return out
}

// Vec2ScaleAndAdd .
func Vec2ScaleAndAdd(out, a, b []float64, scale float64) []float64 {
	out[0] = a[0] + b[0]*scale
	out[1] = a[1] + b[1]*scale
	return out
}

// Vec2Distance .
func Vec2Distance(a, b []float64) float64 {
	x := b[0] - a[0]
	y := b[1] - a[1]
	return math.Hypot(x, y)
}

// Vec2SquaredDistance .
func Vec2SquaredDistance(a, b []float64) float64 {
	x := b[0] - a[0]
	y := b[1] - a[1]
	return x*x + y*y
}

// Vec2Length .
func Vec2Length(out []float64) float64 {
	x := out[0]
	y := out[1]
	return math.Hypot(x, y)
}

// Vec2SquaredLength .
func Vec2SquaredLength(out []float64) float64 {
	x := out[0]
	y := out[1]
	return x*x + y*y
}

// Vec2Negate .
func Vec2Negate(out, a []float64) []float64 {
	out[0] = -a[0]
	out[1] = -a[1]
	return out
}

// Vec2Inverse .
func Vec2Inverse(out, a []float64) []float64 {
	out[0] = 1. / a[0]
	out[1] = 1. / a[1]
	return out
}

// Vec2Normalize .
func Vec2Normalize(out, a []float64) []float64 {
	len := Vec2Length(a)
	if 0 < len {
		len = 1. / len
	}
	out[0] = a[0] * len
	out[1] = a[1] * len
	return out
}

// Vec2Dot .
func Vec2Dot(a, b []float64) float64 {
	return a[0]*b[0] + a[1]*b[1]
}

// Vec2Cross .
// Vec2TODO

// Vec2Lerp .
func Vec2Lerp(out, a, b []float64, t float64) []float64 {
	ax := out[0]
	ay := out[1]
	out[0] = ax + t*(b[0]-ax)
	out[1] = ay + t*(b[1]-ay)
	return out
}

// Vec2Random .
func Vec2Random(out []float64, scale float64) []float64 {
	r := rand.Float64() * 2.0 * math.Pi
	out[0] = math.Cos(r) * scale
	out[1] = math.Sin(r) * scale
	return out
}

// Vec2TransformMat2 .
func Vec2TransformMat2(out, a, m []float64) []float64 {
	x := a[0]
	y := a[1]
	out[0] = m[0]*x + m[2]*y
	out[1] = m[1]*x + m[3]*y
	return out
}

// Vec2TransformMat2d .
func Vec2TransformMat2d(out, a, m []float64) []float64 {
	x := a[0]
	y := a[1]
	out[0] = m[0]*x + m[2]*y + m[4]
	out[1] = m[1]*x + m[3]*y + m[5]
	return out
}

// Vec2TransformMat3 .
func Vec2TransformMat3(out, a, m []float64) []float64 {
	x := a[0]
	y := a[1]
	out[0] = m[0]*x + m[2]*y + m[6]
	out[1] = m[1]*x + m[3]*y + m[7]
	return out
}

// Vec2TransformMat4 .
func Vec2TransformMat4(out, a, m []float64) []float64 {
	x := a[0]
	y := a[1]
	out[0] = m[0]*x + m[2]*y + m[12]
	out[1] = m[1]*x + m[3]*y + m[13]
	return out
}

// Vec2Rotate .
func Vec2Rotate(out, a, b []float64, rad float64) []float64 {
	p0 := a[0] - b[0]
	p1 := a[1] - b[1]
	sinC := math.Sin(rad)
	cosC := math.Cos(rad)

	out[0] = p0*cosC - p1*sinC + b[0]
	out[1] = p0*sinC + p1*cosC + b[1]
	return out
}

// Vec2Angle .
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

// Vec2Zero .
func Vec2Zero(out []float64) []float64 {
	out[0] = 0.
	out[1] = 0.
	return out
}

// Vec2Str .
func Vec2Str(out []float64) string {
	return fmt.Sprintf("vec2(%v, %v)", out[0], out[1])
}

// Vec2ExactEquals .
func Vec2ExactEquals(a, b []float64) bool {
	return a[0] == b[0] && a[1] == b[1]
}

// Vec2Equals .
func Vec2Equals(a, b []float64) bool {
	return equals(a[0], b[0]) && equals(a[1], b[1])
}

// Vec2Len .
var Vec2Len = Vec2Length

// Vec2Sub .
var Vec2Sub = Vec2Subtract

// Vec2Mul .
var Vec2Mul = Vec2Multiply

// Vec2Div .
var Vec2Div = Vec2Divide

// Vec2Dist .
var Vec2Dist = Vec2Distance

// Vec2SqrDist .
var Vec2SqrDist = Vec2SquaredDistance

// Vec2SqrLen .
var Vec2SqrLen = Vec2SquaredLength

// Vec2ForEach .
func Vec2ForEach(a []float64, stride, offset, count int, fn func([]float64, []float64, []interface{}), arg []interface{}) []float64 {
	if stride < 0 {
		stride = 2
	}
	if offset < 0 {
		offset = 0
	}
	var l int
	if 0 <= count {
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
