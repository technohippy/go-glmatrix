package gomatrix

import (
	"fmt"
	"math"
	"math/rand"
)

// NewVec3 .
func NewVec3() []float64 {
	return []float64{0., 0., 0.}
}

// Vec3Create .
func Vec3Create() []float64 {
	return NewVec3()
}

// Vec3Clone .
func Vec3Clone(a []float64) []float64 {
	return []float64{a[0], a[1], a[2]}
}

// Vec3FromValues .
func Vec3FromValues(x, y, z float64) []float64 {
	return []float64{x, y, z}
}

// Vec3Copy .
func Vec3Copy(out, a []float64) []float64 {
	out[0] = a[0]
	out[1] = a[1]
	out[2] = a[2]
	return out
}

// Vec3Set .
func Vec3Set(out []float64, x, y, z float64) []float64 {
	out[0] = x
	out[1] = y
	out[2] = z
	return out
}

// Vec3Add .
func Vec3Add(out, a, b []float64) []float64 {
	out[0] = a[0] + b[0]
	out[1] = a[1] + b[1]
	out[2] = a[2] + b[2]
	return out
}

// Vec3Subtract .
func Vec3Subtract(out, a, b []float64) []float64 {
	out[0] = a[0] - b[0]
	out[1] = a[1] - b[1]
	out[2] = a[2] - b[2]
	return out
}

// Vec3Multiply .
func Vec3Multiply(out, a, b []float64) []float64 {
	out[0] = a[0] * b[0]
	out[1] = a[1] * b[1]
	out[2] = a[2] * b[2]
	return out
}

// Vec3Divide .
func Vec3Divide(out, a, b []float64) []float64 {
	out[0] = a[0] / b[0]
	out[1] = a[1] / b[1]
	out[2] = a[2] / b[2]
	return out
}

// Vec3Ceil .
func Vec3Ceil(out, a []float64) []float64 {
	out[0] = math.Ceil(a[0])
	out[1] = math.Ceil(a[1])
	out[2] = math.Ceil(a[2])
	return out
}

// Vec3Floor .
func Vec3Floor(out, a []float64) []float64 {
	out[0] = math.Floor(a[0])
	out[1] = math.Floor(a[1])
	out[2] = math.Floor(a[2])
	return out
}

// Vec3Min .
func Vec3Min(out, a, b []float64) []float64 {
	out[0] = math.Min(a[0], b[0])
	out[1] = math.Min(a[1], b[1])
	out[2] = math.Min(a[2], b[2])
	return out
}

// Vec3Max .
func Vec3Max(out, a, b []float64) []float64 {
	out[0] = math.Max(a[0], b[0])
	out[1] = math.Max(a[1], b[1])
	out[2] = math.Max(a[2], b[2])
	return out
}

// Vec3Round .
func Vec3Round(out, a []float64) []float64 {
	out[0] = math.Round(a[0])
	out[1] = math.Round(a[1])
	out[2] = math.Round(a[2])
	return out
}

// Vec3Scale .
func Vec3Scale(out, a []float64, scale float64) []float64 {
	out[0] = a[0] * scale
	out[1] = a[1] * scale
	out[2] = a[2] * scale
	return out
}

// Vec3ScaleAndAdd .
func Vec3ScaleAndAdd(out, a, b []float64, scale float64) []float64 {
	out[0] = a[0] + b[0]*scale
	out[1] = a[1] + b[1]*scale
	out[2] = a[2] + b[2]*scale
	return out
}

// Vec3Distance .
func Vec3Distance(a, b []float64) float64 {
	x := b[0] - a[0]
	y := b[1] - a[1]
	z := b[2] - a[2]
	return hypot(x, y, z)
}

// Vec3SquaredDistance .
func Vec3SquaredDistance(a, b []float64) float64 {
	x := b[0] - a[0]
	y := b[1] - a[1]
	z := b[2] - a[2]
	return x*x + y*y + z*z
}

// Vec3Length .
func Vec3Length(out []float64) float64 {
	x := out[0]
	y := out[1]
	z := out[2]
	return hypot(x, y, z)
}

// Vec3SquaredLength .
func Vec3SquaredLength(out []float64) float64 {
	x := out[0]
	y := out[1]
	z := out[2]
	return x*x + y*y + z*z
}

// Vec3Negate .
func Vec3Negate(out, a []float64) []float64 {
	out[0] = -a[0]
	out[1] = -a[1]
	out[2] = -a[2]
	return out
}

// Vec3Inverse .
func Vec3Inverse(out, a []float64) []float64 {
	out[0] = 1. / a[0]
	out[1] = 1. / a[1]
	out[2] = 1. / a[2]
	return out
}

// Vec3Normalize .
func Vec3Normalize(out, a []float64) []float64 {
	len := Vec3Length(a)
	if 0 < len {
		len = 1. / len
	}
	out[0] = a[0] * len
	out[1] = a[1] * len
	out[2] = a[2] * len
	return out
}

// Vec3Dot .
func Vec3Dot(a, b []float64) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

// Vec3Cross .
func Vec3Cross(out, a, b []float64) []float64 {
	ax := a[0]
	ay := a[1]
	az := a[2]
	bx := b[0]
	by := b[1]
	bz := b[2]
	out[0] = ay*bz - az*by
	out[1] = az*bx - ax*bz
	out[2] = ax*by - ay*bx
	return out
}

// Vec3Lerp .
func Vec3Lerp(out, a, b []float64, t float64) []float64 {
	ax := a[0]
	ay := a[1]
	az := a[2]
	out[0] = ax + t*(b[0]-ax)
	out[1] = ay + t*(b[1]-ay)
	out[2] = az + t*(b[2]-az)
	return out
}

// Vec3Slerp .
func Vec3Slerp(out, a, b []float64, t float64) []float64 {
	angle := math.Acos(math.Min(math.Max(Vec3Dot(a, b), -1), 1))
	sinTotal := math.Sin(angle)
	ratioA := math.Sin((1-t)*angle) / sinTotal
	ratioB := math.Sin(t*angle) / sinTotal
	out[0] = ratioA*a[0] + ratioB*b[0]
	out[1] = ratioA*a[1] + ratioB*b[1]
	out[2] = ratioA*a[2] + ratioB*b[2]
	return out
}

// Vec3Hermite .
func Vec3Hermite(out, a, b, c, d []float64, t float64) []float64 {
	factorTimes2 := t * t
	factor1 := factorTimes2*(2*t-3) + 1
	factor2 := factorTimes2*(t-2) + t
	factor3 := factorTimes2 * (t - 1)
	factor4 := factorTimes2 * (3 - 2*t)
	out[0] = a[0]*factor1 + b[0]*factor2 + c[0]*factor3 + d[0]*factor4
	out[1] = a[1]*factor1 + b[1]*factor2 + c[1]*factor3 + d[1]*factor4
	out[2] = a[2]*factor1 + b[2]*factor2 + c[2]*factor3 + d[2]*factor4
	return out
}

// Vec3Bezier .
func Vec3Bezier(out, a, b, c, d []float64, t float64) []float64 {
	inverseFactor := 1 - t
	inverseFactorTimesTwo := inverseFactor * inverseFactor
	factorTimes2 := t * t
	factor1 := inverseFactorTimesTwo * inverseFactor
	factor2 := 3 * t * inverseFactorTimesTwo
	factor3 := 3 * factorTimes2 * inverseFactor
	factor4 := factorTimes2 * t
	out[0] = a[0]*factor1 + b[0]*factor2 + c[0]*factor3 + d[0]*factor4
	out[1] = a[1]*factor1 + b[1]*factor2 + c[1]*factor3 + d[1]*factor4
	out[2] = a[2]*factor1 + b[2]*factor2 + c[2]*factor3 + d[2]*factor4
	return out
}

// Vec3Random .
func Vec3Random(out []float64, scale float64) []float64 {
	r := rand.Float64() * 2.0 * math.Pi
	z := rand.Float64()*2.0 - 1.0
	zScale := math.Sqrt(1.-z*z) * scale

	out[0] = math.Cos(r) * zScale
	out[1] = math.Sin(r) * zScale
	out[2] = z * scale
	return out
}

// Vec3TransformMat3 .
func Vec3TransformMat3(out, a, m []float64) []float64 {
	x := a[0]
	y := a[1]
	z := a[2]
	out[0] = x*m[0] + y*m[3] + z*m[6]
	out[1] = x*m[1] + y*m[4] + z*m[7]
	out[2] = x*m[2] + y*m[5] + z*m[8]
	return out
}

// Vec3TransformMat4 .
func Vec3TransformMat4(out, a, m []float64) []float64 {
	x := a[0]
	y := a[1]
	z := a[2]
	w := m[3]*x + m[7]*y + m[11]*z + m[15]
	if w == 0. {
		w = 1.
	}
	out[0] = (m[0]*x + m[4]*y + m[8]*z + m[12]) / w
	out[1] = (m[1]*x + m[5]*y + m[9]*z + m[13]) / w
	out[2] = (m[2]*x + m[6]*y + m[10]*z + m[14]) / w
	return out
}

// Vec3TransformQuat .
func Vec3TransformQuat(out, a, q []float64) []float64 {
	qx := q[0]
	qy := q[1]
	qz := q[2]
	qw := q[3]
	x := a[0]
	y := a[1]
	z := a[2]
	uvx := qy*z - qz*y
	uvy := qz*x - qx*z
	uvz := qx*y - qy*x
	uuvx := qy*uvz - qz*uvy
	uuvy := qz*uvx - qx*uvz
	uuvz := qx*uvy - qy*uvx
	w2 := qw * 2
	uvx *= w2
	uvy *= w2
	uvz *= w2
	uuvx *= 2
	uuvy *= 2
	uuvz *= 2
	out[0] = x + uvx + uuvx
	out[1] = y + uvy + uuvy
	out[2] = z + uvz + uuvz
	return out
}

// Vec3RotateX .
func Vec3RotateX(out, a, b []float64, rad float64) []float64 {
	p := []float64{
		a[0] - b[0],
		a[1] - b[1],
		a[2] - b[2],
	}
	r := []float64{
		p[0],
		p[1]*math.Cos(rad) - p[2]*math.Sin(rad),
		p[1]*math.Sin(rad) + p[2]*math.Cos(rad),
	}
	out[0] = r[0] + b[0]
	out[1] = r[1] + b[1]
	out[2] = r[2] + b[2]
	return out
}

// Vec3RotateY .
func Vec3RotateY(out, a, b []float64, rad float64) []float64 {
	p := []float64{
		a[0] - b[0],
		a[1] - b[1],
		a[2] - b[2],
	}
	r := []float64{
		p[2]*math.Cos(rad) - p[0]*math.Sin(rad),
		p[1],
		p[2]*math.Sin(rad) + p[0]*math.Cos(rad),
	}
	out[0] = r[0] + b[0]
	out[1] = r[1] + b[1]
	out[2] = r[2] + b[2]
	return out
}

// Vec3RotateZ .
func Vec3RotateZ(out, a, b []float64, rad float64) []float64 {
	p := []float64{
		a[0] - b[0],
		a[1] - b[1],
		a[2] - b[2],
	}
	r := []float64{
		p[0]*math.Cos(rad) - p[1]*math.Sin(rad),
		p[0]*math.Sin(rad) + p[1]*math.Cos(rad),
		p[2],
	}
	out[0] = r[0] + b[0]
	out[1] = r[1] + b[1]
	out[2] = r[2] + b[2]
	return out
}

// Vec3Angle .
func Vec3Angle(a, b []float64) float64 {
	ax := a[0]
	ay := a[1]
	az := a[2]
	bx := b[0]
	by := b[1]
	bz := b[2]
	mag1 := math.Sqrt(ax*ax + ay*ay + az*az)
	mag2 := math.Sqrt(bx*bx + by*by + bz*bz)
	mag := mag1 * mag2
	cosine := mag
	if cosine != 0 {
		cosine = Vec3Dot(a, b) / mag
	}
	return math.Acos(math.Min(math.Max(cosine, -1), 1))
}

// Vec3Zero .
func Vec3Zero(out []float64) []float64 {
	out[0] = 0.
	out[1] = 0.
	out[2] = 0.
	return out
}

// Vec3Str .
func Vec3Str(out []float64) string {
	return fmt.Sprintf("vec3(%v, %v, %v)", out[0], out[1], out[2])
}

// Vec3ExactEquals .
func Vec3ExactEquals(a, b []float64) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2]
}

// Vec3Equals .
func Vec3Equals(a, b []float64) bool {
	return equals(a[0], b[0]) && equals(a[1], b[1]) && equals(a[2], b[2])
}

// Vec3Len .
var Vec3Len = Vec3Length

// Vec3Sub .
var Vec3Sub = Vec3Subtract

// Vec3Mul .
var Vec3Mul = Vec3Multiply

// Vec3Div .
var Vec3Div = Vec3Divide

// Vec3Dist .
var Vec3Dist = Vec3Distance

// Vec3SqrDist .
var Vec3SqrDist = Vec3SquaredDistance

// Vec3SqrLen .
var Vec3SqrLen = Vec3SquaredLength

// Vec3ForEach .
func Vec3ForEach(a []float64, stride, offset, count int, fn func([]float64, []float64, []interface{}), arg []interface{}) []float64 {
	if stride < 0 {
		stride = 3
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
		vec := []float64{a[i], a[i+1], a[i+2]}
		fn(vec, vec, arg)
		a[i] = vec[0]
		a[i+1] = vec[1]
		a[i+2] = vec[2]
	}
	return a
}
