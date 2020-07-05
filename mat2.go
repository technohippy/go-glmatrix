package glmatrix

import (
	"fmt"
	"math"
)

// Mat2Create creates a new identity mat2
func Mat2Create() []float64 {
	return []float64{
		1, 0,
		0, 1,
	}
}

// Mat2Clone creates a new mat2 initialized with values from an existing matrix
func Mat2Clone(a []float64) []float64 {
	return []float64{
		a[0], a[1],
		a[2], a[3],
	}
}

// Mat2Copy copy the values from one mat2 to another
func Mat2Copy(out, a []float64) []float64 {
	out[0] = a[0]
	out[1] = a[1]
	out[2] = a[2]
	out[3] = a[3]
	return out
}

// Mat2Identity set a mat2 to the identity matrix
func Mat2Identity(out []float64) []float64 {
	out[0] = 1
	out[1] = 0
	out[2] = 0
	out[3] = 1
	return out
}

// Mat2FromValues create a new mat2 with the given values
func Mat2FromValues(m00, m01, m10, m11 float64) []float64 {
	return []float64{
		m00, m01,
		m10, m11,
	}
}

// Mat2Set set the components of a mat2 to the given values
func Mat2Set(out []float64, m00, m01, m10, m11 float64) []float64 {
	out[0] = m00
	out[1] = m01
	out[2] = m10
	out[3] = m11
	return out
}

// Mat2Transpose transpose the values of a mat2
func Mat2Transpose(out, a []float64) []float64 {
	// If we are transposing ourselves we can skip a few steps but have to cache
	// some values
	if &out == &a {
		a1 := a[1]
		out[1] = a[2]
		out[2] = a1
	} else {
		out[0] = a[0]
		out[1] = a[2]
		out[2] = a[1]
		out[3] = a[3]
	}

	return out
}

// Mat2Invert inverts a mat2
func Mat2Invert(out, a []float64) []float64 {
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]

	// Calculate the determinant
	det := a0*a3 - a2*a1

	if det == 0. {
		return nil
	}
	det = 1.0 / det

	out[0] = a3 * det
	out[1] = -a1 * det
	out[2] = -a2 * det
	out[3] = a0 * det

	return out
}

// Mat2Adjoint calculates the adjugate of a mat2
func Mat2Adjoint(out, a []float64) []float64 {
	// Caching this value is nessesary if out == a
	a0 := a[0]
	out[0] = a[3]
	out[1] = -a[1]
	out[2] = -a[2]
	out[3] = a0

	return out
}

// Mat2Determinant calculates the determinant of a mat2
func Mat2Determinant(a []float64) float64 {
	return a[0]*a[3] - a[2]*a[1]
}

// Mat2Multiply multiplies two mat2's
func Mat2Multiply(out, a, b []float64) []float64 {
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]
	b0 := b[0]
	b1 := b[1]
	b2 := b[2]
	b3 := b[3]
	out[0] = a0*b0 + a2*b1
	out[1] = a1*b0 + a3*b1
	out[2] = a0*b2 + a2*b3
	out[3] = a1*b2 + a3*b3
	return out
}

// Mat2Rotate rotates a mat2 by the given angle
func Mat2Rotate(out, a []float64, rad float64) []float64 {
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]
	s := math.Sin(rad)
	c := math.Cos(rad)
	out[0] = a0*c + a2*s
	out[1] = a1*c + a3*s
	out[2] = a0*-s + a2*c
	out[3] = a1*-s + a3*c
	return out
}

// Mat2Scale scales the mat2 by the dimensions in the given vec2
func Mat2Scale(out, a, v []float64) []float64 {
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]
	v0 := v[0]
	v1 := v[1]
	out[0] = a0 * v0
	out[1] = a1 * v0
	out[2] = a2 * v1
	out[3] = a3 * v1
	return out
}

// Mat2FromRotation creates a matrix from a given angle
// This is equivalent to (but much faster than):
//
// - Mat2Identity(dest)
// - Mat2Rotate(dest, dest, rad)
func Mat2FromRotation(out []float64, rad float64) []float64 {
	s := math.Sin(rad)
	c := math.Cos(rad)
	out[0] = c
	out[1] = s
	out[2] = -s
	out[3] = c
	return out
}

// Mat2FromScaling creates a matrix from a vector scaling
// This is equivalent to (but much faster than):
//
// - Mat2Identity(dest)
// - Mat2Scale(dest, dest, vec)
func Mat2FromScaling(out, v []float64) []float64 {
	out[0] = v[0]
	out[1] = 0
	out[2] = 0
	out[3] = v[1]
	return out
}

// Mat2Str returns a string representation of a mat2
func Mat2Str(a []float64) string {
	return fmt.Sprintf("mat2(%v, %v, %v, %v)", a[0], a[1], a[2], a[3])
}

// Mat2Frob returns Frobenius norm of a mat2
func Mat2Frob(a []float64) float64 {
	return hypot(a[0], a[1], a[2], a[3])
}

// Mat2LDU returns L, D and U matrices (Lower triangular, Diagonal and Upper triangular) by factorizing the input matrix
func Mat2LDU(L, D, U, a []float64) [][]float64 {
	L[2] = a[2] / a[0]
	U[0] = a[0]
	U[1] = a[1]
	U[3] = a[3] - L[2]*U[1]
	return [][]float64{L, D, U}
}

// Mat2Add adds two mat2's
func Mat2Add(out, a, b []float64) []float64 {
	out[0] = a[0] + b[0]
	out[1] = a[1] + b[1]
	out[2] = a[2] + b[2]
	out[3] = a[3] + b[3]
	return out
}

// Mat2Subtract subtracts matrix b from matrix a
func Mat2Subtract(out, a, b []float64) []float64 {
	out[0] = a[0] - b[0]
	out[1] = a[1] - b[1]
	out[2] = a[2] - b[2]
	out[3] = a[3] - b[3]
	return out
}

// Mat2ExactEquals returns whether or not the matrices have exactly the same elements in the same position (when compared with ==)
func Mat2ExactEquals(a, b []float64) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3]
}

// Mat2Equals returns whether or not the matrices have approximately the same elements in the same position.
func Mat2Equals(a, b []float64) bool {
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]
	b0 := b[0]
	b1 := b[1]
	b2 := b[2]
	b3 := b[3]
	return equals(a0, b0) && equals(a1, b1) && equals(a2, b2) && equals(a3, b3)
}

// Mat2MultiplyScalar multiply each element of the matrix by a scalar.
func Mat2MultiplyScalar(out, a []float64, b float64) []float64 {
	out[0] = a[0] * b
	out[1] = a[1] * b
	out[2] = a[2] * b
	out[3] = a[3] * b
	return out
}

// Mat2MultiplyScalarAndAdd adds two mat2's after multiplying each element of the second operand by a scalar value.
func Mat2MultiplyScalarAndAdd(out, a, b []float64, scale float64) []float64 {
	out[0] = a[0] + b[0]*scale
	out[1] = a[1] + b[1]*scale
	out[2] = a[2] + b[2]*scale
	out[3] = a[3] + b[3]*scale
	return out
}

// Mat2Mul alias for Mat2Multiply
var Mat2Mul = Mat2Multiply

// Mat2Sub alias for Mat2Subtract
var Mat2Sub = Mat2Subtract
