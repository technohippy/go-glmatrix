package gomatrix

import (
	"fmt"
	"math"
)

// Mat2dCreate creates a new identity mat2d
func Mat2dCreate() []float64 {
	return []float64{
		1, 0,
		0, 1,
		0, 0,
	}
}

// Mat2dClone creates a new mat2d initialized with values from an existing matrix
func Mat2dClone(a []float64) []float64 {
	return []float64{
		a[0], a[1],
		a[2], a[3],
		a[4], a[5],
	}
}

// Mat2dCopy copy the values from one mat2d to another
func Mat2dCopy(out, a []float64) []float64 {
	out[0] = a[0]
	out[1] = a[1]
	out[2] = a[2]
	out[3] = a[3]
	out[4] = a[4]
	out[5] = a[5]
	return out
}

// Mat2dIdentity set a mat2d to the identity matrix
func Mat2dIdentity(out []float64) []float64 {
	out[0] = 1
	out[1] = 0
	out[2] = 0
	out[3] = 1
	out[4] = 0
	out[5] = 0
	return out
}

// Mat2dFromValues create a new mat2d with the given values
func Mat2dFromValues(a, b, c, d, tx, ty float64) []float64 {
	return []float64{
		a, b,
		c, d,
		tx, ty,
	}
}

// Mat2dSet set the components of a mat2d to the given values
func Mat2dSet(out []float64, a, b, c, d, tx, ty float64) []float64 {
	out[0] = a
	out[1] = b
	out[2] = c
	out[3] = d
	out[4] = tx
	out[5] = ty
	return out
}

// Mat2dInvert inverts a mat2d
func Mat2dInvert(out, a []float64) []float64 {
	aa := a[0]
	ab := a[1]
	ac := a[2]
	ad := a[3]
	atx := a[4]
	aty := a[5]

	det := aa*ad - ab*ac
	if det == 0. {
		return nil
	}
	det = 1.0 / det

	out[0] = ad * det
	out[1] = -ab * det
	out[2] = -ac * det
	out[3] = aa * det
	out[4] = (ac*aty - ad*atx) * det
	out[5] = (ab*atx - aa*aty) * det
	return out
}

// Mat2dDeterminant calculates the determinant of a mat2d
func Mat2dDeterminant(a []float64) float64 {
	return a[0]*a[3] - a[1]*a[2]
}

// Mat2dMultiply multiplies two mat2d's
func Mat2dMultiply(out, a, b []float64) []float64 {
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]
	a4 := a[4]
	a5 := a[5]
	b0 := b[0]
	b1 := b[1]
	b2 := b[2]
	b3 := b[3]
	b4 := b[4]
	b5 := b[5]
	out[0] = a0*b0 + a2*b1
	out[1] = a1*b0 + a3*b1
	out[2] = a0*b2 + a2*b3
	out[3] = a1*b2 + a3*b3
	out[4] = a0*b4 + a2*b5 + a4
	out[5] = a1*b4 + a3*b5 + a5
	return out
}

// Mat2dRotate rotates a mat2d by the given angle
func Mat2dRotate(out, a []float64, rad float64) []float64 {
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]
	a4 := a[4]
	a5 := a[5]
	s := math.Sin(rad)
	c := math.Cos(rad)
	out[0] = a0*c + a2*s
	out[1] = a1*c + a3*s
	out[2] = a0*-s + a2*c
	out[3] = a1*-s + a3*c
	out[4] = a4
	out[5] = a5
	return out
}

// Mat2dScale scales the mat2d by the dimensions in the given vec2
func Mat2dScale(out, a, v []float64) []float64 {
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]
	a4 := a[4]
	a5 := a[5]
	v0 := v[0]
	v1 := v[1]
	out[0] = a0 * v0
	out[1] = a1 * v0
	out[2] = a2 * v1
	out[3] = a3 * v1
	out[4] = a4
	out[5] = a5
	return out
}

// Mat2dTranslate translates the mat2d by the dimensions in the given vec2
func Mat2dTranslate(out, a, v []float64) []float64 {
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]
	a4 := a[4]
	a5 := a[5]
	v0 := v[0]
	v1 := v[1]
	out[0] = a0
	out[1] = a1
	out[2] = a2
	out[3] = a3
	out[4] = a0*v0 + a2*v1 + a4
	out[5] = a1*v0 + a3*v1 + a5
	return out
}

// Mat2dFromRotation creates a matrix from a given angle
// This is equivalent to (but much faster than):
//
// - Mat2dIdentity(dest)
// - Mat2dRotate(dest, dest, rad)
func Mat2dFromRotation(out []float64, rad float64) []float64 {
	s := math.Sin(rad)
	c := math.Cos(rad)
	out[0] = c
	out[1] = s
	out[2] = -s
	out[3] = c
	out[4] = 0
	out[5] = 0
	return out
}

// Mat2dFromScaling creates a matrix from a vector scaling
// This is equivalent to (but much faster than):
//
// - Mat2dIdentity(dest)
// - Mat2dScale(dest, dest, vec)
func Mat2dFromScaling(out, v []float64) []float64 {
	out[0] = v[0]
	out[1] = 0
	out[2] = 0
	out[3] = v[1]
	out[4] = 0
	out[5] = 0
	return out
}

// Mat2dFromTranslation creates a matrix from a vector translation
// This is equivalent to (but much faster than):
//
// - Mat2dIdentity(dest)
// - Mat2dTranslate(dest, dest, vec)
func Mat2dFromTranslation(out, v []float64) []float64 {
	out[0] = 1
	out[1] = 0
	out[2] = 0
	out[3] = 1
	out[4] = v[0]
	out[5] = v[1]
	return out
}

// Mat2dStr returns a string representation of a mat2d
func Mat2dStr(a []float64) string {
	return fmt.Sprintf("mat2d(%v, %v, %v, %v, %v, %v)", a[0], a[1], a[2], a[3], a[4], a[5])
}

// Mat2dFrob returns Frobenius norm of a mat2d
func Mat2dFrob(a []float64) float64 {
	return hypot(a[0], a[1], a[2], a[3], a[4], a[5], 1)
}

// Mat2dAdd adds two mat2d's
func Mat2dAdd(out, a, b []float64) []float64 {
	out[0] = a[0] + b[0]
	out[1] = a[1] + b[1]
	out[2] = a[2] + b[2]
	out[3] = a[3] + b[3]
	out[4] = a[4] + b[4]
	out[5] = a[5] + b[5]
	return out
}

// Mat2dSubtract subtracts matrix b from matrix a
func Mat2dSubtract(out, a, b []float64) []float64 {
	out[0] = a[0] - b[0]
	out[1] = a[1] - b[1]
	out[2] = a[2] - b[2]
	out[3] = a[3] - b[3]
	out[4] = a[4] - b[4]
	out[5] = a[5] - b[5]
	return out
}

// Mat2dMultiplyScalar multiply each element of the matrix by a scalar.
func Mat2dMultiplyScalar(out, a []float64, b float64) []float64 {
	out[0] = a[0] * b
	out[1] = a[1] * b
	out[2] = a[2] * b
	out[3] = a[3] * b
	out[4] = a[4] * b
	out[5] = a[5] * b
	return out
}

// Mat2dMultiplyScalarAndAdd adds two mat2d's after multiplying each element of the second operand by a scalar value.
func Mat2dMultiplyScalarAndAdd(out, a, b []float64, scale float64) []float64 {
	out[0] = a[0] + b[0]*scale
	out[1] = a[1] + b[1]*scale
	out[2] = a[2] + b[2]*scale
	out[3] = a[3] + b[3]*scale
	out[4] = a[4] + b[4]*scale
	out[5] = a[5] + b[5]*scale
	return out
}

// Mat2dExactEquals returns whether or not the matrices have exactly the same elements in the same position (when compared with ==)
func Mat2dExactEquals(a, b []float64) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3] && a[4] == b[4] && a[5] == b[5]
}

// Mat2dEquals returns whether or not the matrices have approximately the same elements in the same position.
func Mat2dEquals(a, b []float64) bool {
	return equals(a[0], b[0]) && equals(a[1], b[1]) && equals(a[2], b[2]) && equals(a[3], b[3]) && equals(a[4], b[4]) && equals(a[5], b[5])
}

// Mat2dMul alias for Mat2dMultiply
var Mat2dMul = Mat2dMultiply

// Mat2dSub alias for Mat2dSubtract
var Mat2dSub = Mat2dSubtract
