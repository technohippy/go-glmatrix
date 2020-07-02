package gomatrix

import (
	"fmt"
	"math"
)

// NewMat3 creates a new identity mat3
func NewMat3() []float64 {
	return []float64{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

// Mat3Create creates a new identity mat3
func Mat3Create() []float64 {
	return NewMat3()
}

// Mat3FromMat4 copies the upper-left 3x3 values into the given mat3.
func Mat3FromMat4(out, a []float64) []float64 {
	out[0] = a[0]
	out[1] = a[1]
	out[2] = a[2]
	out[3] = a[4]
	out[4] = a[5]
	out[5] = a[6]
	out[6] = a[8]
	out[7] = a[9]
	out[8] = a[10]
	return out
}

// Mat3Clone creates a new mat3 initialized with values from an existing matrix
func Mat3Clone(a []float64) []float64 {
	return []float64{
		a[0], a[1], a[2],
		a[3], a[4], a[5],
		a[6], a[7], a[8],
	}
}

// Mat3Copy copy the values from one mat3 to another
func Mat3Copy(out, a []float64) []float64 {
	out[0] = a[0]
	out[1] = a[1]
	out[2] = a[2]
	out[3] = a[3]
	out[4] = a[4]
	out[5] = a[5]
	out[6] = a[6]
	out[7] = a[7]
	out[8] = a[8]
	return out
}

// Mat3FromValues create a new mat3 with the given values
func Mat3FromValues(m00, m01, m02, m10, m11, m12, m20, m21, m22 float64) []float64 {
	return []float64{
		m00, m01, m02,
		m10, m11, m12,
		m20, m21, m22,
	}
}

// Mat3Set set the components of a mat3 to the given values
func Mat3Set(out []float64, m00, m01, m02, m10, m11, m12, m20, m21, m22 float64) []float64 {
	out[0] = m00
	out[1] = m01
	out[2] = m02
	out[3] = m10
	out[4] = m11
	out[5] = m12
	out[6] = m20
	out[7] = m21
	out[8] = m22
	return out
}

// Mat3Identity set a mat3 to the identity matrix
func Mat3Identity(out []float64) []float64 {
	out[0] = 1
	out[1] = 0
	out[2] = 0
	out[3] = 0
	out[4] = 1
	out[5] = 0
	out[6] = 0
	out[7] = 0
	out[8] = 1
	return out
}

// Mat3Transpose transpose the values of a mat3
func Mat3Transpose(out, a []float64) []float64 {
	// If we are transposing ourselves we can skip a few steps but have to cache some values
	if &out == &a {
		a01 := a[1]
		a02 := a[2]
		a12 := a[5]
		out[1] = a[3]
		out[2] = a[6]
		out[3] = a01
		out[5] = a[7]
		out[6] = a02
		out[7] = a12
	} else {
		out[0] = a[0]
		out[1] = a[3]
		out[2] = a[6]
		out[3] = a[1]
		out[4] = a[4]
		out[5] = a[7]
		out[6] = a[2]
		out[7] = a[5]
		out[8] = a[8]
	}

	return out
}

// Mat3Invert inverts a mat3
func Mat3Invert(out, a []float64) []float64 {
	a00 := a[0]
	a01 := a[1]
	a02 := a[2]
	a10 := a[3]
	a11 := a[4]
	a12 := a[5]
	a20 := a[6]
	a21 := a[7]
	a22 := a[8]

	b01 := a22*a11 - a12*a21
	b11 := -a22*a10 + a12*a20
	b21 := a21*a10 - a11*a20

	// Calculate the determinant
	det := a00*b01 + a01*b11 + a02*b21

	if det == 0. {
		return nil
	}
	det = 1.0 / det

	out[0] = b01 * det
	out[1] = (-a22*a01 + a02*a21) * det
	out[2] = (a12*a01 - a02*a11) * det
	out[3] = b11 * det
	out[4] = (a22*a00 - a02*a20) * det
	out[5] = (-a12*a00 + a02*a10) * det
	out[6] = b21 * det
	out[7] = (-a21*a00 + a01*a20) * det
	out[8] = (a11*a00 - a01*a10) * det
	return out
}

// Mat3Adjoint calculates the adjugate of a mat3
func Mat3Adjoint(out, a []float64) []float64 {
	a00 := a[0]
	a01 := a[1]
	a02 := a[2]
	a10 := a[3]
	a11 := a[4]
	a12 := a[5]
	a20 := a[6]
	a21 := a[7]
	a22 := a[8]

	out[0] = a11*a22 - a12*a21
	out[1] = a02*a21 - a01*a22
	out[2] = a01*a12 - a02*a11
	out[3] = a12*a20 - a10*a22
	out[4] = a00*a22 - a02*a20
	out[5] = a02*a10 - a00*a12
	out[6] = a10*a21 - a11*a20
	out[7] = a01*a20 - a00*a21
	out[8] = a00*a11 - a01*a10
	return out
}

// Mat3Determinant calculates the determinant of a mat3
func Mat3Determinant(a []float64) float64 {
	a00 := a[0]
	a01 := a[1]
	a02 := a[2]
	a10 := a[3]
	a11 := a[4]
	a12 := a[5]
	a20 := a[6]
	a21 := a[7]
	a22 := a[8]

	return a00*(a22*a11-a12*a21) +
		a01*(-a22*a10+a12*a20) +
		a02*(a21*a10-a11*a20)
}

// Mat3Multiply multiplies two mat3's
func Mat3Multiply(out, a, b []float64) []float64 {
	a00 := a[0]
	a01 := a[1]
	a02 := a[2]
	a10 := a[3]
	a11 := a[4]
	a12 := a[5]
	a20 := a[6]
	a21 := a[7]
	a22 := a[8]

	b00 := b[0]
	b01 := b[1]
	b02 := b[2]
	b10 := b[3]
	b11 := b[4]
	b12 := b[5]
	b20 := b[6]
	b21 := b[7]
	b22 := b[8]

	out[0] = b00*a00 + b01*a10 + b02*a20
	out[1] = b00*a01 + b01*a11 + b02*a21
	out[2] = b00*a02 + b01*a12 + b02*a22

	out[3] = b10*a00 + b11*a10 + b12*a20
	out[4] = b10*a01 + b11*a11 + b12*a21
	out[5] = b10*a02 + b11*a12 + b12*a22

	out[6] = b20*a00 + b21*a10 + b22*a20
	out[7] = b20*a01 + b21*a11 + b22*a21
	out[8] = b20*a02 + b21*a12 + b22*a22
	return out
}

// Mat3Translate translate a mat3 by the given vector
func Mat3Translate(out, a, v []float64) []float64 {
	a00 := a[0]
	a01 := a[1]
	a02 := a[2]
	a10 := a[3]
	a11 := a[4]
	a12 := a[5]
	a20 := a[6]
	a21 := a[7]
	a22 := a[8]
	x := v[0]
	y := v[1]

	out[0] = a00
	out[1] = a01
	out[2] = a02

	out[3] = a10
	out[4] = a11
	out[5] = a12

	out[6] = x*a00 + y*a10 + a20
	out[7] = x*a01 + y*a11 + a21
	out[8] = x*a02 + y*a12 + a22
	return out
}

// Mat3Rotate rotates a mat3 by the given angle
func Mat3Rotate(out, a []float64, rad float64) []float64 {
	a00 := a[0]
	a01 := a[1]
	a02 := a[2]
	a10 := a[3]
	a11 := a[4]
	a12 := a[5]
	a20 := a[6]
	a21 := a[7]
	a22 := a[8]
	s := math.Sin(rad)
	c := math.Cos(rad)

	out[0] = c*a00 + s*a10
	out[1] = c*a01 + s*a11
	out[2] = c*a02 + s*a12

	out[3] = c*a10 - s*a00
	out[4] = c*a11 - s*a01
	out[5] = c*a12 - s*a02

	out[6] = a20
	out[7] = a21
	out[8] = a22
	return out
}

// Mat3Scale scales the mat3 by the dimensions in the given vec2
func Mat3Scale(out, a, v []float64) []float64 {
	x := v[0]
	y := v[1]

	out[0] = x * a[0]
	out[1] = x * a[1]
	out[2] = x * a[2]

	out[3] = y * a[3]
	out[4] = y * a[4]
	out[5] = y * a[5]

	out[6] = a[6]
	out[7] = a[7]
	out[8] = a[8]
	return out
}

// Mat3FromTranslation creates a matrix from a vector translation
// This is equivalent to (but much faster than):
//
// - Mat3Identity(dest)
// - Mat3Translate(dest, dest, vec)
func Mat3FromTranslation(out, v []float64) []float64 {
	out[0] = 1
	out[1] = 0
	out[2] = 0
	out[3] = 0
	out[4] = 1
	out[5] = 0
	out[6] = v[0]
	out[7] = v[1]
	out[8] = 1
	return out
}

// Mat3FromRotation creates a matrix from a given angle
// This is equivalent to (but much faster than):
//
// - Mat3Identity(dest)
// - Mat3Rotate(dest, dest, rad)
func Mat3FromRotation(out []float64, rad float64) []float64 {
	s := math.Sin(rad)
	c := math.Cos(rad)

	out[0] = c
	out[1] = s
	out[2] = 0

	out[3] = -s
	out[4] = c
	out[5] = 0

	out[6] = 0
	out[7] = 0
	out[8] = 1
	return out
}

// Mat3FromScaling creates a matrix from a vector scaling
// This is equivalent to (but much faster than):
//
// - Mat3Identity(dest)
// - Mat3Scale(dest, dest, vec)
func Mat3FromScaling(out, v []float64) []float64 {
	out[0] = v[0]
	out[1] = 0
	out[2] = 0

	out[3] = 0
	out[4] = v[1]
	out[5] = 0

	out[6] = 0
	out[7] = 0
	out[8] = 1
	return out
}

// Mat3FromMat2d copies the values from a mat2d into a mat3
func Mat3FromMat2d(out, a []float64) []float64 {
	out[0] = a[0]
	out[1] = a[1]
	out[2] = 0

	out[3] = a[2]
	out[4] = a[3]
	out[5] = 0

	out[6] = a[4]
	out[7] = a[5]
	out[8] = 1
	return out
}

// Mat3FromQuat calculates a 3x3 matrix from the given quaternion
func Mat3FromQuat(out, q []float64) []float64 {
	x := q[0]
	y := q[1]
	z := q[2]
	w := q[3]
	x2 := x + x
	y2 := y + y
	z2 := z + z

	xx := x * x2
	yx := y * x2
	yy := y * y2
	zx := z * x2
	zy := z * y2
	zz := z * z2
	wx := w * x2
	wy := w * y2
	wz := w * z2

	out[0] = 1 - yy - zz
	out[3] = yx - wz
	out[6] = zx + wy

	out[1] = yx + wz
	out[4] = 1 - xx - zz
	out[7] = zy - wx

	out[2] = zx - wy
	out[5] = zy + wx
	out[8] = 1 - xx - yy

	return out
}

// Mat3NormalFromMat4 calculates a 3x3 normal matrix (transpose inverse) from the 4x4 matrix
func Mat3NormalFromMat4(out, a []float64) []float64 {
	a00 := a[0]
	a01 := a[1]
	a02 := a[2]
	a03 := a[3]
	a10 := a[4]
	a11 := a[5]
	a12 := a[6]
	a13 := a[7]
	a20 := a[8]
	a21 := a[9]
	a22 := a[10]
	a23 := a[11]
	a30 := a[12]
	a31 := a[13]
	a32 := a[14]
	a33 := a[15]

	b00 := a00*a11 - a01*a10
	b01 := a00*a12 - a02*a10
	b02 := a00*a13 - a03*a10
	b03 := a01*a12 - a02*a11
	b04 := a01*a13 - a03*a11
	b05 := a02*a13 - a03*a12
	b06 := a20*a31 - a21*a30
	b07 := a20*a32 - a22*a30
	b08 := a20*a33 - a23*a30
	b09 := a21*a32 - a22*a31
	b10 := a21*a33 - a23*a31
	b11 := a22*a33 - a23*a32

	// Calculate the determinant
	det := b00*b11 - b01*b10 + b02*b09 + b03*b08 - b04*b07 + b05*b06

	if det == 0. {
		return nil
	}
	det = 1.0 / det

	out[0] = (a11*b11 - a12*b10 + a13*b09) * det
	out[1] = (a12*b08 - a10*b11 - a13*b07) * det
	out[2] = (a10*b10 - a11*b08 + a13*b06) * det

	out[3] = (a02*b10 - a01*b11 - a03*b09) * det
	out[4] = (a00*b11 - a02*b08 + a03*b07) * det
	out[5] = (a01*b08 - a00*b10 - a03*b06) * det

	out[6] = (a31*b05 - a32*b04 + a33*b03) * det
	out[7] = (a32*b02 - a30*b05 - a33*b01) * det
	out[8] = (a30*b04 - a31*b02 + a33*b00) * det

	return out
}

// Mat3Projection generates a 2D projection matrix with the given bounds
func Mat3Projection(out []float64, width, height float64) []float64 {
	out[0] = 2 / width
	out[1] = 0
	out[2] = 0
	out[3] = 0
	out[4] = -2 / height
	out[5] = 0
	out[6] = -1
	out[7] = 1
	out[8] = 1
	return out
}

// Mat3Str returns a string representation of a mat3
func Mat3Str(a []float64) string {
	return fmt.Sprintf("mat3(%v, %v, %v, %v, %v, %v, %v, %v, %v)",
		a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8])
}

// Mat3Frob returns Frobenius norm of a mat3
func Mat3Frob(a []float64) float64 {
	return hypot(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7])
}

// Mat3Add adds two mat3's
func Mat3Add(out, a, b []float64) []float64 {
	out[0] = a[0] + b[0]
	out[1] = a[1] + b[1]
	out[2] = a[2] + b[2]
	out[3] = a[3] + b[3]
	out[4] = a[4] + b[4]
	out[5] = a[5] + b[5]
	out[6] = a[6] + b[6]
	out[7] = a[7] + b[7]
	out[8] = a[8] + b[8]
	return out
}

// Mat3Subtract subtracts matrix b from matrix a
func Mat3Subtract(out, a, b []float64) []float64 {
	out[0] = a[0] - b[0]
	out[1] = a[1] - b[1]
	out[2] = a[2] - b[2]
	out[3] = a[3] - b[3]
	out[4] = a[4] - b[4]
	out[5] = a[5] - b[5]
	out[6] = a[6] - b[6]
	out[7] = a[7] - b[7]
	out[8] = a[8] - b[8]
	return out
}

// Mat3MultiplyScalar multiply each element of the matrix by a scalar.
func Mat3MultiplyScalar(out, a []float64, b float64) []float64 {
	out[0] = a[0] * b
	out[1] = a[1] * b
	out[2] = a[2] * b
	out[3] = a[3] * b
	out[4] = a[4] * b
	out[5] = a[5] * b
	out[6] = a[6] * b
	out[7] = a[7] * b
	out[8] = a[8] * b
	return out
}

// Mat3MultiplyScalarAndAdd adds two mat3's after multiplying each element of the second operand by a scalar value.
func Mat3MultiplyScalarAndAdd(out, a, b []float64, scale float64) []float64 {
	out[0] = a[0] + b[0]*scale
	out[1] = a[1] + b[1]*scale
	out[2] = a[2] + b[2]*scale
	out[3] = a[3] + b[3]*scale
	out[4] = a[4] + b[4]*scale
	out[5] = a[5] + b[5]*scale
	out[6] = a[6] + b[6]*scale
	out[7] = a[7] + b[7]*scale
	out[8] = a[8] + b[8]*scale
	return out
}

// Mat3ExactEquals returns whether or not the matrices have exactly the same elements in the same position (when compared with ===)
func Mat3ExactEquals(a, b []float64) bool {
	return a[0] == b[0] &&
		a[1] == b[1] &&
		a[2] == b[2] &&
		a[3] == b[3] &&
		a[4] == b[4] &&
		a[5] == b[5] &&
		a[6] == b[6] &&
		a[7] == b[7] &&
		a[8] == b[8]
}

// Mat3Equals returns whether or not the matrices have approximately the same elements in the same position.
func Mat3Equals(a, b []float64) bool {
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]
	a4 := a[4]
	a5 := a[5]
	a6 := a[6]
	a7 := a[7]
	a8 := a[8]
	b0 := b[0]
	b1 := b[1]
	b2 := b[2]
	b3 := b[3]
	b4 := b[4]
	b5 := b[5]
	b6 := b[6]
	b7 := b[7]
	b8 := b[8]
	return equals(a0, b0) &&
		equals(a1, b1) &&
		equals(a2, b2) &&
		equals(a3, b3) &&
		equals(a4, b4) &&
		equals(a5, b5) &&
		equals(a6, b6) &&
		equals(a7, b7) &&
		equals(a8, b8)
}

// Mat3Mul alias for Mat3Multiply
var Mat3Mul = Mat3Multiply

// Mat3Sub alias for Mat3Subtract
var Mat3Sub = Mat3Subtract
