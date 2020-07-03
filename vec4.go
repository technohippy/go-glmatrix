package glmatrix

import (
	"fmt"
	"math"
	"math/rand"
)

// NewVec4 creates a new, empty  Vec4
func NewVec4() []float64 {
	return []float64{0., 0., 0., 0.}
}

// Vec4Create creates a new  Vec4 initialized with values from an existing vector
func Vec4Create() []float64 {
	return NewVec4()
}

// Vec4Clone creates a new  Vec4 initialized with the given values
func Vec4Clone(a []float64) []float64 {
	return []float64{a[0], a[1], a[2], a[3]}
}

// Vec4FromValues creates a new  Vec4 initialized with the given values
func Vec4FromValues(x, y, z, w float64) []float64 {
	return []float64{x, y, z, w}
}

// Vec4Copy copy the values from one  Vec4 to another
func Vec4Copy(out, a []float64) []float64 {
	out[0] = a[0]
	out[1] = a[1]
	out[2] = a[2]
	out[3] = a[3]
	return out
}

// Vec4Set set the components of a vec4 to the given values
func Vec4Set(out []float64, x, y, z, w float64) []float64 {
	out[0] = x
	out[1] = y
	out[2] = z
	out[3] = w
	return out
}

// Vec4Add adds two  Vec4's
func Vec4Add(out, a, b []float64) []float64 {
	out[0] = a[0] + b[0]
	out[1] = a[1] + b[1]
	out[2] = a[2] + b[2]
	out[3] = a[3] + b[3]
	return out
}

// Vec4Subtract subtracts vector b from vector a
func Vec4Subtract(out, a, b []float64) []float64 {
	out[0] = a[0] - b[0]
	out[1] = a[1] - b[1]
	out[2] = a[2] - b[2]
	out[3] = a[3] - b[3]
	return out
}

// Vec4Multiply multiplies two  Vec4's
func Vec4Multiply(out, a, b []float64) []float64 {
	out[0] = a[0] * b[0]
	out[1] = a[1] * b[1]
	out[2] = a[2] * b[2]
	out[3] = a[3] * b[3]
	return out
}

// Vec4Divide divides two  Vec4's
func Vec4Divide(out, a, b []float64) []float64 {
	out[0] = a[0] / b[0]
	out[1] = a[1] / b[1]
	out[2] = a[2] / b[2]
	out[3] = a[3] / b[3]
	return out
}

// Vec4Ceil math.ceil the components of a vec4
func Vec4Ceil(out, a []float64) []float64 {
	out[0] = math.Ceil(a[0])
	out[1] = math.Ceil(a[1])
	out[2] = math.Ceil(a[2])
	out[3] = math.Ceil(a[3])
	return out
}

// Vec4Floor math.floor the components of a vec4
func Vec4Floor(out, a []float64) []float64 {
	out[0] = math.Floor(a[0])
	out[1] = math.Floor(a[1])
	out[2] = math.Floor(a[2])
	out[3] = math.Floor(a[3])
	return out
}

// Vec4Min returns the minimum of two  Vec4's
func Vec4Min(out, a, b []float64) []float64 {
	out[0] = math.Min(a[0], b[0])
	out[1] = math.Min(a[1], b[1])
	out[2] = math.Min(a[2], b[2])
	out[3] = math.Min(a[3], b[3])
	return out
}

// Vec4Max returns the maximum of two  Vec4's
func Vec4Max(out, a, b []float64) []float64 {
	out[0] = math.Max(a[0], b[0])
	out[1] = math.Max(a[1], b[1])
	out[2] = math.Max(a[2], b[2])
	out[3] = math.Max(a[3], b[3])
	return out
}

// Vec4Round math.round the components of a vec4
func Vec4Round(out, a []float64) []float64 {
	out[0] = math.Round(a[0])
	out[1] = math.Round(a[1])
	out[2] = math.Round(a[2])
	out[3] = math.Round(a[3])
	return out
}

// Vec4Scale scales a vec4 by a scalar number
func Vec4Scale(out, a []float64, scale float64) []float64 {
	out[0] = a[0] * scale
	out[1] = a[1] * scale
	out[2] = a[2] * scale
	out[3] = a[3] * scale
	return out
}

// Vec4ScaleAndAdd adds two  Vec4's after scaling the second operand by a scalar value
func Vec4ScaleAndAdd(out, a, b []float64, scale float64) []float64 {
	out[0] = a[0] + b[0]*scale
	out[1] = a[1] + b[1]*scale
	out[2] = a[2] + b[2]*scale
	out[3] = a[3] + b[3]*scale
	return out
}

// Vec4Distance calculates the euclidian distance between two  Vec4's
func Vec4Distance(a, b []float64) float64 {
	x := b[0] - a[0]
	y := b[1] - a[1]
	z := b[2] - a[2]
	w := b[3] - a[3]
	return hypot(x, y, z, w)
}

// Vec4SquaredDistance calculates the squared euclidian distance between two  Vec4's
func Vec4SquaredDistance(a, b []float64) float64 {
	x := b[0] - a[0]
	y := b[1] - a[1]
	z := b[2] - a[2]
	w := b[3] - a[3]
	return x*x + y*y + z*z + w*w
}

// Vec4Length calculates the length of a vec4
func Vec4Length(out []float64) float64 {
	x := out[0]
	y := out[1]
	z := out[2]
	w := out[3]
	return hypot(x, y, z, w)
}

// Vec4SquaredLength calculates the squared length of a vec4
func Vec4SquaredLength(out []float64) float64 {
	x := out[0]
	y := out[1]
	z := out[2]
	w := out[3]
	return x*x + y*y + z*z + w*w
}

// Vec4Negate negates the components of a vec4
func Vec4Negate(out, a []float64) []float64 {
	out[0] = -a[0]
	out[1] = -a[1]
	out[2] = -a[2]
	out[3] = -a[3]
	return out
}

// Vec4Inverse returns the inverse of the components of a vec4
func Vec4Inverse(out, a []float64) []float64 {
	out[0] = 1. / a[0]
	out[1] = 1. / a[1]
	out[2] = 1. / a[2]
	out[3] = 1. / a[3]
	return out
}

// Vec4Normalize normalize a vec4
func Vec4Normalize(out, a []float64) []float64 {
	len := Vec4Length(a)
	if 0 < len {
		len = 1. / len
	}
	out[0] = a[0] * len
	out[1] = a[1] * len
	out[2] = a[2] * len
	out[3] = a[3] * len
	return out
}

// Vec4Dot calculates the dot product of two  Vec4's
func Vec4Dot(a, b []float64) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]
}

// Vec4Cross computes the cross product of two  Vec4's
func Vec4Cross(out, u, v, w []float64) []float64 {
	A := v[0]*w[1] - v[1]*w[0]
	B := v[0]*w[2] - v[2]*w[0]
	C := v[0]*w[3] - v[3]*w[0]
	D := v[1]*w[2] - v[2]*w[1]
	E := v[1]*w[3] - v[3]*w[1]
	F := v[2]*w[3] - v[3]*w[2]

	G := u[0]
	H := u[1]
	I := u[2]
	J := u[3]

	out[0] = H*F - I*E + J*D
	out[1] = -(G * F) + I*C - J*B
	out[2] = G*E - H*C + J*A
	out[3] = -(G * D) + H*B - I*A
	return out
}

// Vec4Lerp performs a linear interpolation between two  Vec4's
func Vec4Lerp(out, a, b []float64, t float64) []float64 {
	ax := a[0]
	ay := a[1]
	az := a[2]
	aw := a[3]
	out[0] = ax + t*(b[0]-ax)
	out[1] = ay + t*(b[1]-ay)
	out[2] = az + t*(b[2]-az)
	out[3] = aw + t*(b[3]-aw)
	return out
}

// Vec4Random generates a random vector with the given scale
func Vec4Random(out []float64, scale float64) []float64 {
	if scale == 0. {
		scale = 1.
	}

	var v1, v2, v3, v4, s1, s2 float64
	for {
		v1 = rand.Float64()*2 - 1
		v2 = rand.Float64()*2 - 1
		s1 = v1*v1 + v2*v2
		if s1 >= 1 {
			break
		}
	}
	for {
		v3 = rand.Float64()*2 - 1
		v4 = rand.Float64()*2 - 1
		s2 = v3*v3 + v4*v4
		if s2 >= 1 {
			break
		}
	}

	d := math.Sqrt((1 - s1) / s2)
	out[0] = scale * v1
	out[1] = scale * v2
	out[2] = scale * v3 * d
	out[3] = scale * v4 * d
	return out
}

// Vec4TransformMat4 transforms the vec4 with a mat4
func Vec4TransformMat4(out, a, m []float64) []float64 {
	x := a[0]
	y := a[1]
	z := a[2]
	w := a[3]
	out[0] = (m[0]*x + m[4]*y + m[8]*z + m[12]) / w
	out[1] = (m[1]*x + m[5]*y + m[9]*z + m[13]) / w
	out[2] = (m[2]*x + m[6]*y + m[10]*z + m[14]) / w
	out[3] = (m[3]*x + m[7]*y + m[11]*z + m[15]) / w
	return out
}

// Vec4TransformQuat transforms the vec4 with a quat
func Vec4TransformQuat(out, a, q []float64) []float64 {
	qx := q[0]
	qy := q[1]
	qz := q[2]
	qw := q[3]
	x := a[0]
	y := a[1]
	z := a[2]
	ix := qw*x + qy*z - qz*y
	iy := qw*y + qy*x - qz*z
	iz := qw*z + qy*y - qz*x
	iw := -qw*x - qy*y - qz*z

	out[0] = ix*qw + iw*-qx + iy*-qz - iz*-qy
	out[1] = iy*qw + iw*-qy + iz*-qx - ix*-qz
	out[2] = iz*qw + iw*-qz + ix*-qy - iy*-qx
	out[3] = a[3]
	return out
}

// Vec4Zero set the components of a vec4 to zero
func Vec4Zero(out []float64) []float64 {
	out[0] = 0.
	out[1] = 0.
	out[2] = 0.
	out[3] = 0.
	return out
}

// Vec4Str returns a string representation of a vector
func Vec4Str(a []float64) string {
	return fmt.Sprintf("vec4(%v, %v, %v, %v)", a[0], a[1], a[2], a[3])
}

// Vec4ExactEquals returns whether or not the vectors exactly have the same elements in the same position (when compared with ===)
func Vec4ExactEquals(a, b []float64) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[2]
}

// Vec4Equals returns whether or not the vectors have approximately the same elements in the same position.
func Vec4Equals(a, b []float64) bool {
	return equals(a[0], b[0]) && equals(a[1], b[1]) && equals(a[2], b[2]) && equals(a[3], b[3])
}

// Vec4Len alias for  Vec4Length
var Vec4Len = Vec4Length

// Vec4Sub alias for  Vec4Subtract
var Vec4Sub = Vec4Subtract

// Vec4Mul alias for  Vec4Multiply
var Vec4Mul = Vec4Multiply

// Vec4Div alias for  Vec4Divide
var Vec4Div = Vec4Divide

// Vec4Dist alias for  Vec4Distance
var Vec4Dist = Vec4Distance

// Vec4SqrDist alias for  Vec4SquaredDistance
var Vec4SqrDist = Vec4SquaredDistance

// Vec4SqrLen alias for Vec4SquaredLength
var Vec4SqrLen = Vec4SquaredLength

// Vec4ForEach perform some operation over an array of  Vec4s.
func Vec4ForEach(a []float64, stride, offset, count int, fn func([]float64, []float64, []interface{}), arg []interface{}) []float64 {
	if stride < 0 {
		stride = 4
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
		vec := []float64{a[i], a[i+1], a[i+2], a[i+3]}
		fn(vec, vec, arg)
		a[i] = vec[0]
		a[i+1] = vec[1]
		a[i+2] = vec[2]
		a[i+3] = vec[3]
	}
	return a
}
