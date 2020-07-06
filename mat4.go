package glmatrix

import (
	"fmt"
	"math"
)

// NewMat4 creates a new identity mat4
func NewMat4() []float64 {
	return []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Mat4Create creates a new identity mat4
func Mat4Create() []float64 {
	return NewMat4()
}

// Mat4Clone creates a new mat4 initialized with values from an existing matrix
func Mat4Clone(a []float64) []float64 {
	return []float64{
		a[0], a[1], a[2], a[3],
		a[4], a[5], a[6], a[7],
		a[8], a[9], a[10], a[11],
		a[12], a[13], a[14], a[15],
	}
}

// Mat4Copy copy the values from one mat4 to another
func Mat4Copy(out, a []float64) []float64 {
	out[0] = a[0]
	out[1] = a[1]
	out[2] = a[2]
	out[3] = a[3]
	out[4] = a[4]
	out[5] = a[5]
	out[6] = a[6]
	out[7] = a[7]
	out[8] = a[8]
	out[9] = a[9]
	out[10] = a[10]
	out[11] = a[11]
	out[12] = a[12]
	out[13] = a[13]
	out[14] = a[14]
	out[15] = a[15]
	return out
}

// Mat4FromValues create a new mat4 with the given values
func Mat4FromValues(m00, m01, m02, m03, m10, m11, m12, m13, m20, m21, m22, m23, m30, m31, m32, m33 float64) []float64 {
	return []float64{
		m00, m01, m02, m03,
		m10, m11, m12, m13,
		m20, m21, m22, m23,
		m30, m31, m32, m33,
	}
}

// Mat4Set set the components of a mat4 to the given values
func Mat4Set(out []float64, m00, m01, m02, m03, m10, m11, m12, m13, m20, m21, m22, m23, m30, m31, m32, m33 float64) []float64 {
	out[0] = m00
	out[1] = m01
	out[2] = m02
	out[3] = m03
	out[4] = m10
	out[5] = m11
	out[6] = m12
	out[7] = m13
	out[8] = m20
	out[9] = m21
	out[10] = m22
	out[11] = m23
	out[12] = m30
	out[13] = m31
	out[14] = m32
	out[15] = m33
	return out
}

// Mat4Identity set a mat4 to the identity matrix
func Mat4Identity(out []float64) []float64 {
	out[0] = 1
	out[1] = 0
	out[2] = 0
	out[3] = 0
	out[4] = 0
	out[5] = 1
	out[6] = 0
	out[7] = 0
	out[8] = 0
	out[9] = 0
	out[10] = 1
	out[11] = 0
	out[12] = 0
	out[13] = 0
	out[14] = 0
	out[15] = 1
	return out
}

// Mat4Transpose transpose the values of a mat4
func Mat4Transpose(out, a []float64) []float64 {
	// If we are transposing ourselves we can skip a few steps but have to cache some values
	if &(out[0]) == &(a[0]) {
		a01 := a[1]
		a02 := a[2]
		a03 := a[3]
		a12 := a[6]
		a13 := a[7]
		a23 := a[11]

		out[1] = a[4]
		out[2] = a[8]
		out[3] = a[12]
		out[4] = a01
		out[6] = a[9]
		out[7] = a[13]
		out[8] = a02
		out[9] = a12
		out[11] = a[14]
		out[12] = a03
		out[13] = a13
		out[14] = a23
	} else {
		out[0] = a[0]
		out[1] = a[4]
		out[2] = a[8]
		out[3] = a[12]
		out[4] = a[1]
		out[5] = a[5]
		out[6] = a[9]
		out[7] = a[13]
		out[8] = a[2]
		out[9] = a[6]
		out[10] = a[10]
		out[11] = a[14]
		out[12] = a[3]
		out[13] = a[7]
		out[14] = a[11]
		out[15] = a[15]
	}

	return out
}

// Mat4Invert inverts a mat4
func Mat4Invert(out, a []float64) []float64 {
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
	det :=
		b00*b11 - b01*b10 + b02*b09 + b03*b08 - b04*b07 + b05*b06

	if det == 0. {
		return nil
	}
	det = 1.0 / det

	out[0] = (a11*b11 - a12*b10 + a13*b09) * det
	out[1] = (a02*b10 - a01*b11 - a03*b09) * det
	out[2] = (a31*b05 - a32*b04 + a33*b03) * det
	out[3] = (a22*b04 - a21*b05 - a23*b03) * det
	out[4] = (a12*b08 - a10*b11 - a13*b07) * det
	out[5] = (a00*b11 - a02*b08 + a03*b07) * det
	out[6] = (a32*b02 - a30*b05 - a33*b01) * det
	out[7] = (a20*b05 - a22*b02 + a23*b01) * det
	out[8] = (a10*b10 - a11*b08 + a13*b06) * det
	out[9] = (a01*b08 - a00*b10 - a03*b06) * det
	out[10] = (a30*b04 - a31*b02 + a33*b00) * det
	out[11] = (a21*b02 - a20*b04 - a23*b00) * det
	out[12] = (a11*b07 - a10*b09 - a12*b06) * det
	out[13] = (a00*b09 - a01*b07 + a02*b06) * det
	out[14] = (a31*b01 - a30*b03 - a32*b00) * det
	out[15] = (a20*b03 - a21*b01 + a22*b00) * det
	return out
}

// Mat4Adjoint calculates the adjugate of a mat4
func Mat4Adjoint(out, a []float64) []float64 {
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

	out[0] = a11*b11 - a12*b10 + a13*b09
	out[1] = a02*b10 - a01*b11 - a03*b09
	out[2] = a31*b05 - a32*b04 + a33*b03
	out[3] = a22*b04 - a21*b05 - a23*b03
	out[4] = a12*b08 - a10*b11 - a13*b07
	out[5] = a00*b11 - a02*b08 + a03*b07
	out[6] = a32*b02 - a30*b05 - a33*b01
	out[7] = a20*b05 - a22*b02 + a23*b01
	out[8] = a10*b10 - a11*b08 + a13*b06
	out[9] = a01*b08 - a00*b10 - a03*b06
	out[10] = a30*b04 - a31*b02 + a33*b00
	out[11] = a21*b02 - a20*b04 - a23*b00
	out[12] = a11*b07 - a10*b09 - a12*b06
	out[13] = a00*b09 - a01*b07 + a02*b06
	out[14] = a31*b01 - a30*b03 - a32*b00
	out[15] = a20*b03 - a21*b01 + a22*b00
	return out
}

// Mat4Determinant calculates the determinant of a mat4
func Mat4Determinant(a []float64) float64 {
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

	b0 := a00*a11 - a01*a10
	b1 := a00*a12 - a02*a10
	b2 := a01*a12 - a02*a11
	b3 := a20*a31 - a21*a30
	b4 := a20*a32 - a22*a30
	b5 := a21*a32 - a22*a31
	b6 := a00*b5 - a01*b4 + a02*b3
	b7 := a10*b5 - a11*b4 + a12*b3
	b8 := a20*b2 - a21*b1 + a22*b0
	b9 := a30*b2 - a31*b1 + a32*b0

	// Calculate the determinant
	return a13*b6 - a03*b7 + a33*b8 - a23*b9
}

// Mat4Multiply multiplies two mat4s
func Mat4Multiply(out, a, b []float64) []float64 {
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

	// Cache only the current line of the second matrix
	b0 := b[0]
	b1 := b[1]
	b2 := b[2]
	b3 := b[3]
	out[0] = b0*a00 + b1*a10 + b2*a20 + b3*a30
	out[1] = b0*a01 + b1*a11 + b2*a21 + b3*a31
	out[2] = b0*a02 + b1*a12 + b2*a22 + b3*a32
	out[3] = b0*a03 + b1*a13 + b2*a23 + b3*a33

	b0 = b[4]
	b1 = b[5]
	b2 = b[6]
	b3 = b[7]
	out[4] = b0*a00 + b1*a10 + b2*a20 + b3*a30
	out[5] = b0*a01 + b1*a11 + b2*a21 + b3*a31
	out[6] = b0*a02 + b1*a12 + b2*a22 + b3*a32
	out[7] = b0*a03 + b1*a13 + b2*a23 + b3*a33

	b0 = b[8]
	b1 = b[9]
	b2 = b[10]
	b3 = b[11]
	out[8] = b0*a00 + b1*a10 + b2*a20 + b3*a30
	out[9] = b0*a01 + b1*a11 + b2*a21 + b3*a31
	out[10] = b0*a02 + b1*a12 + b2*a22 + b3*a32
	out[11] = b0*a03 + b1*a13 + b2*a23 + b3*a33

	b0 = b[12]
	b1 = b[13]
	b2 = b[14]
	b3 = b[15]
	out[12] = b0*a00 + b1*a10 + b2*a20 + b3*a30
	out[13] = b0*a01 + b1*a11 + b2*a21 + b3*a31
	out[14] = b0*a02 + b1*a12 + b2*a22 + b3*a32
	out[15] = b0*a03 + b1*a13 + b2*a23 + b3*a33
	return out
}

// Mat4Translate translate a mat4 by the given vector
func Mat4Translate(out, a, v []float64) []float64 {
	x := v[0]
	y := v[1]
	z := v[2]

	if &(a[0]) == &(out[0]) {
		out[12] = a[0]*x + a[4]*y + a[8]*z + a[12]
		out[13] = a[1]*x + a[5]*y + a[9]*z + a[13]
		out[14] = a[2]*x + a[6]*y + a[10]*z + a[14]
		out[15] = a[3]*x + a[7]*y + a[11]*z + a[15]
	} else {
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

		out[0] = a00
		out[1] = a01
		out[2] = a02
		out[3] = a03
		out[4] = a10
		out[5] = a11
		out[6] = a12
		out[7] = a13
		out[8] = a20
		out[9] = a21
		out[10] = a22
		out[11] = a23

		out[12] = a00*x + a10*y + a20*z + a[12]
		out[13] = a01*x + a11*y + a21*z + a[13]
		out[14] = a02*x + a12*y + a22*z + a[14]
		out[15] = a03*x + a13*y + a23*z + a[15]
	}

	return out
}

// Mat4Scale scales the mat4 by the dimensions in the given vec3 not using vectorization
func Mat4Scale(out, a, v []float64) []float64 {
	x := v[0]
	y := v[1]
	z := v[2]

	out[0] = a[0] * x
	out[1] = a[1] * x
	out[2] = a[2] * x
	out[3] = a[3] * x
	out[4] = a[4] * y
	out[5] = a[5] * y
	out[6] = a[6] * y
	out[7] = a[7] * y
	out[8] = a[8] * z
	out[9] = a[9] * z
	out[10] = a[10] * z
	out[11] = a[11] * z
	out[12] = a[12]
	out[13] = a[13]
	out[14] = a[14]
	out[15] = a[15]
	return out
}

// Mat4Rotate rotates a mat4 by the given angle around the given axis
func Mat4Rotate(out, a []float64, rad float64, axis []float64) []float64 {
	x := axis[0]
	y := axis[1]
	z := axis[2]
	len := hypot(x, y, z)

	if len < Epsilon {
		return nil
	}

	len = 1 / len
	x *= len
	y *= len
	z *= len

	s := math.Sin(rad)
	c := math.Cos(rad)
	t := 1 - c

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

	// Construct the elements of the rotation matrix
	b00 := x*x*t + c
	b01 := y*x*t + z*s
	b02 := z*x*t - y*s
	b10 := x*y*t - z*s
	b11 := y*y*t + c
	b12 := z*y*t + x*s
	b20 := x*z*t + y*s
	b21 := y*z*t - x*s
	b22 := z*z*t + c

	// Perform rotation-specific matrix multiplication
	out[0] = a00*b00 + a10*b01 + a20*b02
	out[1] = a01*b00 + a11*b01 + a21*b02
	out[2] = a02*b00 + a12*b01 + a22*b02
	out[3] = a03*b00 + a13*b01 + a23*b02
	out[4] = a00*b10 + a10*b11 + a20*b12
	out[5] = a01*b10 + a11*b11 + a21*b12
	out[6] = a02*b10 + a12*b11 + a22*b12
	out[7] = a03*b10 + a13*b11 + a23*b12
	out[8] = a00*b20 + a10*b21 + a20*b22
	out[9] = a01*b20 + a11*b21 + a21*b22
	out[10] = a02*b20 + a12*b21 + a22*b22
	out[11] = a03*b20 + a13*b21 + a23*b22

	if &a != &out {
		// If the source and destination differ, copy the unchanged last row
		out[12] = a[12]
		out[13] = a[13]
		out[14] = a[14]
		out[15] = a[15]
	}
	return out
}

// Mat4RotateX rotates a matrix by the given angle around the X axis
func Mat4RotateX(out, a []float64, rad float64) []float64 {
	s := math.Sin(rad)
	c := math.Cos(rad)
	a10 := a[4]
	a11 := a[5]
	a12 := a[6]
	a13 := a[7]
	a20 := a[8]
	a21 := a[9]
	a22 := a[10]
	a23 := a[11]

	if &a != &out {
		// If the source and destination differ, copy the unchanged rows
		out[0] = a[0]
		out[1] = a[1]
		out[2] = a[2]
		out[3] = a[3]
		out[12] = a[12]
		out[13] = a[13]
		out[14] = a[14]
		out[15] = a[15]
	}

	// Perform axis-specific matrix multiplication
	out[4] = a10*c + a20*s
	out[5] = a11*c + a21*s
	out[6] = a12*c + a22*s
	out[7] = a13*c + a23*s
	out[8] = a20*c - a10*s
	out[9] = a21*c - a11*s
	out[10] = a22*c - a12*s
	out[11] = a23*c - a13*s
	return out
}

// Mat4RotateY rotates a matrix by the given angle around the Y axis
func Mat4RotateY(out, a []float64, rad float64) []float64 {
	s := math.Sin(rad)
	c := math.Cos(rad)
	a00 := a[0]
	a01 := a[1]
	a02 := a[2]
	a03 := a[3]
	a20 := a[8]
	a21 := a[9]
	a22 := a[10]
	a23 := a[11]

	if &a != &out {
		// If the source and destination differ, copy the unchanged rows
		out[4] = a[4]
		out[5] = a[5]
		out[6] = a[6]
		out[7] = a[7]
		out[12] = a[12]
		out[13] = a[13]
		out[14] = a[14]
		out[15] = a[15]
	}

	// Perform axis-specific matrix multiplication
	out[0] = a00*c - a20*s
	out[1] = a01*c - a21*s
	out[2] = a02*c - a22*s
	out[3] = a03*c - a23*s
	out[8] = a00*s + a20*c
	out[9] = a01*s + a21*c
	out[10] = a02*s + a22*c
	out[11] = a03*s + a23*c
	return out
}

// Mat4RotateZ rotates a matrix by the given angle around the Z axis
func Mat4RotateZ(out, a []float64, rad float64) []float64 {
	s := math.Sin(rad)
	c := math.Cos(rad)
	a00 := a[0]
	a01 := a[1]
	a02 := a[2]
	a03 := a[3]
	a10 := a[4]
	a11 := a[5]
	a12 := a[6]
	a13 := a[7]

	if &a != &out {
		// If the source and destination differ, copy the unchanged rows
		out[8] = a[8]
		out[9] = a[9]
		out[10] = a[10]
		out[11] = a[11]
		out[12] = a[12]
		out[13] = a[13]
		out[14] = a[14]
		out[15] = a[15]
	}

	// Perform axis-specific matrix multiplication
	out[0] = a00*c + a10*s
	out[1] = a01*c + a11*s
	out[2] = a02*c + a12*s
	out[3] = a03*c + a13*s
	out[4] = a10*c - a00*s
	out[5] = a11*c - a01*s
	out[6] = a12*c - a02*s
	out[7] = a13*c - a03*s
	return out
}

// Mat4FromTranslation creates a matrix from a vector translation
// This is equivalent to (but much faster than):
//
// - Mat4Identity(dest)
// - Mat4Translate(dest, dest, vec)
func Mat4FromTranslation(out, v []float64) []float64 {
	out[0] = 1
	out[1] = 0
	out[2] = 0
	out[3] = 0
	out[4] = 0
	out[5] = 1
	out[6] = 0
	out[7] = 0
	out[8] = 0
	out[9] = 0
	out[10] = 1
	out[11] = 0
	out[12] = v[0]
	out[13] = v[1]
	out[14] = v[2]
	out[15] = 1
	return out
}

// Mat4FromScaling creates a matrix from a vector scaling
// This is equivalent to (but much faster than):
//
// - Mat4Identity(dest)
// - Mat4Scale(dest, dest, vec)
func Mat4FromScaling(out, v []float64) []float64 {
	out[0] = v[0]
	out[1] = 0
	out[2] = 0
	out[3] = 0
	out[4] = 0
	out[5] = v[1]
	out[6] = 0
	out[7] = 0
	out[8] = 0
	out[9] = 0
	out[10] = v[2]
	out[11] = 0
	out[12] = 0
	out[13] = 0
	out[14] = 0
	out[15] = 1
	return out
}

// Mat4FromRotation creates a matrix from a given angle around a given axis
// This is equivalent to (but much faster than):
//
// - Mat4Identity(dest)
// - Mat4Rotate(dest, dest, rad, axis)
func Mat4FromRotation(out []float64, rad float64, axis []float64) []float64 {
	x := axis[0]
	y := axis[1]
	z := axis[2]
	len := hypot(x, y, z)

	if len < Epsilon {
		return nil
	}

	len = 1 / len
	x *= len
	y *= len
	z *= len

	s := math.Sin(rad)
	c := math.Cos(rad)
	t := 1 - c

	// Perform rotation-specific matrix multiplication
	out[0] = x*x*t + c
	out[1] = y*x*t + z*s
	out[2] = z*x*t - y*s
	out[3] = 0
	out[4] = x*y*t - z*s
	out[5] = y*y*t + c
	out[6] = z*y*t + x*s
	out[7] = 0
	out[8] = x*z*t + y*s
	out[9] = y*z*t - x*s
	out[10] = z*z*t + c
	out[11] = 0
	out[12] = 0
	out[13] = 0
	out[14] = 0
	out[15] = 1
	return out
}

// Mat4FromXRotation creates a matrix from the given angle around the X axis
// This is equivalent to (but much faster than):
//
// - Mat4Identity(dest)
// - Mat4RotateX(dest, dest, rad)
func Mat4FromXRotation(out []float64, rad float64) []float64 {
	s := math.Sin(rad)
	c := math.Cos(rad)

	// Perform axis-specific matrix multiplication
	out[0] = 1
	out[1] = 0
	out[2] = 0
	out[3] = 0
	out[4] = 0
	out[5] = c
	out[6] = s
	out[7] = 0
	out[8] = 0
	out[9] = -s
	out[10] = c
	out[11] = 0
	out[12] = 0
	out[13] = 0
	out[14] = 0
	out[15] = 1
	return out
}

// Mat4FromYRotation creates a matrix from the given angle around the Y axis
// This is equivalent to (but much faster than):
//
// - Mat4Identity(dest)
// - Mat4RotateY(dest, dest, rad)
func Mat4FromYRotation(out []float64, rad float64) []float64 {
	s := math.Sin(rad)
	c := math.Cos(rad)

	// Perform axis-specific matrix multiplication
	out[0] = c
	out[1] = 0
	out[2] = -s
	out[3] = 0
	out[4] = 0
	out[5] = 1
	out[6] = 0
	out[7] = 0
	out[8] = s
	out[9] = 0
	out[10] = c
	out[11] = 0
	out[12] = 0
	out[13] = 0
	out[14] = 0
	out[15] = 1
	return out
}

// Mat4FromZRotation creates a matrix from the given angle around the Z axis
// This is equivalent to (but much faster than):
//
// - Mat4Identity(dest)
// - Mat4RotateZ(dest, dest, rad)
func Mat4FromZRotation(out []float64, rad float64) []float64 {
	s := math.Sin(rad)
	c := math.Cos(rad)

	// Perform axis-specific matrix multiplication
	out[0] = c
	out[1] = s
	out[2] = 0
	out[3] = 0
	out[4] = -s
	out[5] = c
	out[6] = 0
	out[7] = 0
	out[8] = 0
	out[9] = 0
	out[10] = 1
	out[11] = 0
	out[12] = 0
	out[13] = 0
	out[14] = 0
	out[15] = 1
	return out
}

// Mat4FromRotationTranslation creates a matrix from a quaternion rotation and vector translation
// This is equivalent to (but much faster than):
//
// - Mat4Identity(dest)
// - Mat4Translate(dest, vec)
// - quatMat := Mat4Create()
// - Quat4ToMat4(quat, quatMat)
// - Mat4Multiply(dest, quatMat)
func Mat4FromRotationTranslation(out, q, v []float64) []float64 {
	// Quaternion math
	x := q[0]
	y := q[1]
	z := q[2]
	w := q[3]
	x2 := x + x
	y2 := y + y
	z2 := z + z

	xx := x * x2
	xy := x * y2
	xz := x * z2
	yy := y * y2
	yz := y * z2
	zz := z * z2
	wx := w * x2
	wy := w * y2
	wz := w * z2

	out[0] = 1 - (yy + zz)
	out[1] = xy + wz
	out[2] = xz - wy
	out[3] = 0
	out[4] = xy - wz
	out[5] = 1 - (xx + zz)
	out[6] = yz + wx
	out[7] = 0
	out[8] = xz + wy
	out[9] = yz - wx
	out[10] = 1 - (xx + yy)
	out[11] = 0
	out[12] = v[0]
	out[13] = v[1]
	out[14] = v[2]
	out[15] = 1

	return out
}

// Mat4FromQuat2 creates a new mat4 from a dual quat.
func Mat4FromQuat2(out, a []float64) []float64 {
	translation := []float64{0., 0., 0.}
	bx := -a[0]
	by := -a[1]
	bz := -a[2]
	bw := a[3]
	ax := a[4]
	ay := a[5]
	az := a[6]
	aw := a[7]

	magnitude := bx*bx + by*by + bz*bz + bw*bw
	//Only scale if it makes sense
	if magnitude > 0 {
		translation[0] = ((ax*bw + aw*bx + ay*bz - az*by) * 2) / magnitude
		translation[1] = ((ay*bw + aw*by + az*bx - ax*bz) * 2) / magnitude
		translation[2] = ((az*bw + aw*bz + ax*by - ay*bx) * 2) / magnitude
	} else {
		translation[0] = (ax*bw + aw*bx + ay*bz - az*by) * 2
		translation[1] = (ay*bw + aw*by + az*bx - ax*bz) * 2
		translation[2] = (az*bw + aw*bz + ax*by - ay*bx) * 2
	}
	Mat4FromRotationTranslation(out, a, translation)
	return out
}

// Mat4GetTranslation returns the translation vector component of a transformation
// matrix. If a matrix is built with fromRotationTranslation,
// the returned vector will be the same as the translation vector
// originally supplied.
func Mat4GetTranslation(out, mat []float64) []float64 {
	out[0] = mat[12]
	out[1] = mat[13]
	out[2] = mat[14]

	return out
}

// Mat4GetScaling returns the scaling factor component of a transformation
// matrix. If a matrix is built with fromRotationTranslationScale
// with a normalized Quaternion parameter, the returned vector will be
// the same as the scaling vector
// originally supplied.
func Mat4GetScaling(out, mat []float64) []float64 {
	m11 := mat[0]
	m12 := mat[1]
	m13 := mat[2]
	m21 := mat[4]
	m22 := mat[5]
	m23 := mat[6]
	m31 := mat[8]
	m32 := mat[9]
	m33 := mat[10]

	out[0] = hypot(m11, m12, m13)
	out[1] = hypot(m21, m22, m23)
	out[2] = hypot(m31, m32, m33)

	return out
}

// Mat4GetRotation returns a quaternion representing the rotational component
// of a transformation matrix. If a matrix is built with
// fromRotationTranslation, the returned quaternion will be the
// same as the quaternion originally supplied.
func Mat4GetRotation(out, mat []float64) []float64 {
	scaling := []float64{0., 0., 0.}
	Mat4GetScaling(scaling, mat)

	is1 := 1 / scaling[0]
	is2 := 1 / scaling[1]
	is3 := 1 / scaling[2]

	sm11 := mat[0] * is1
	sm12 := mat[1] * is2
	sm13 := mat[2] * is3
	sm21 := mat[4] * is1
	sm22 := mat[5] * is2
	sm23 := mat[6] * is3
	sm31 := mat[8] * is1
	sm32 := mat[9] * is2
	sm33 := mat[10] * is3

	trace := sm11 + sm22 + sm33
	S := 0.

	if trace > 0 {
		S = math.Sqrt(trace+1.0) * 2
		out[3] = 0.25 * S
		out[0] = (sm23 - sm32) / S
		out[1] = (sm31 - sm13) / S
		out[2] = (sm12 - sm21) / S
	} else if sm11 > sm22 && sm11 > sm33 {
		S = math.Sqrt(1.0+sm11-sm22-sm33) * 2
		out[3] = (sm23 - sm32) / S
		out[0] = 0.25 * S
		out[1] = (sm12 + sm21) / S
		out[2] = (sm31 + sm13) / S
	} else if sm22 > sm33 {
		S = math.Sqrt(1.0+sm22-sm11-sm33) * 2
		out[3] = (sm31 - sm13) / S
		out[0] = (sm12 + sm21) / S
		out[1] = 0.25 * S
		out[2] = (sm23 + sm32) / S
	} else {
		S = math.Sqrt(1.0+sm33-sm11-sm22) * 2
		out[3] = (sm12 - sm21) / S
		out[0] = (sm31 + sm13) / S
		out[1] = (sm23 + sm32) / S
		out[2] = 0.25 * S
	}

	return out
}

// Mat4FromRotationTranslationScale creates a matrix from a quaternion rotation, vector translation and vector scale
// This is equivalent to (but much faster than):
//
// - Mat4Identity(dest)
// - Mat4Translate(dest, vec)
// - quatMat := Mat4Create()
// - Quat4ToMat4(quat, quatMat)
// - Mat4Multiply(dest, quatMat)
// - Mat4Scale(dest, scale)
func Mat4FromRotationTranslationScale(out, q, v, s []float64) []float64 {
	// Quaternion math
	x := q[0]
	y := q[1]
	z := q[2]
	w := q[3]
	x2 := x + x
	y2 := y + y
	z2 := z + z

	xx := x * x2
	xy := x * y2
	xz := x * z2
	yy := y * y2
	yz := y * z2
	zz := z * z2
	wx := w * x2
	wy := w * y2
	wz := w * z2
	sx := s[0]
	sy := s[1]
	sz := s[2]

	out[0] = (1 - (yy + zz)) * sx
	out[1] = (xy + wz) * sx
	out[2] = (xz - wy) * sx
	out[3] = 0
	out[4] = (xy - wz) * sy
	out[5] = (1 - (xx + zz)) * sy
	out[6] = (yz + wx) * sy
	out[7] = 0
	out[8] = (xz + wy) * sz
	out[9] = (yz - wx) * sz
	out[10] = (1 - (xx + yy)) * sz
	out[11] = 0
	out[12] = v[0]
	out[13] = v[1]
	out[14] = v[2]
	out[15] = 1

	return out
}

// Mat4FromRotationTranslationScaleOrigin creates a matrix from a quaternion rotation, vector translation and vector scale, rotating and scaling around the given origin
// This is equivalent to (but much faster than):
//
// - Mat4Identity(dest)
// - Mat4Translate(dest, vec)
// - Mat4Translate(dest, origin)
// - quatMat := Mat4Create()
// - Quat4ToMat4(quat, quatMat)
// - Mat4Multiply(dest, quatMat)
// - Mat4Scale(dest, scale)
// - Mat4Translate(dest, negativeOrigin)
func Mat4FromRotationTranslationScaleOrigin(out, q, v, s, o []float64) []float64 {
	// Quaternion math
	x := q[0]
	y := q[1]
	z := q[2]
	w := q[3]
	x2 := x + x
	y2 := y + y
	z2 := z + z

	xx := x * x2
	xy := x * y2
	xz := x * z2
	yy := y * y2
	yz := y * z2
	zz := z * z2
	wx := w * x2
	wy := w * y2
	wz := w * z2

	sx := s[0]
	sy := s[1]
	sz := s[2]

	ox := o[0]
	oy := o[1]
	oz := o[2]

	out0 := (1 - (yy + zz)) * sx
	out1 := (xy + wz) * sx
	out2 := (xz - wy) * sx
	out4 := (xy - wz) * sy
	out5 := (1 - (xx + zz)) * sy
	out6 := (yz + wx) * sy
	out8 := (xz + wy) * sz
	out9 := (yz - wx) * sz
	out10 := (1 - (xx + yy)) * sz

	out[0] = out0
	out[1] = out1
	out[2] = out2
	out[3] = 0
	out[4] = out4
	out[5] = out5
	out[6] = out6
	out[7] = 0
	out[8] = out8
	out[9] = out9
	out[10] = out10
	out[11] = 0
	out[12] = v[0] + ox - (out0*ox + out4*oy + out8*oz)
	out[13] = v[1] + oy - (out1*ox + out5*oy + out9*oz)
	out[14] = v[2] + oz - (out2*ox + out6*oy + out10*oz)
	out[15] = 1

	return out
}

// Mat4FromQuat calculates a 4x4 matrix from the given quaternion
func Mat4FromQuat(out, q []float64) []float64 {
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
	out[1] = yx + wz
	out[2] = zx - wy
	out[3] = 0

	out[4] = yx - wz
	out[5] = 1 - xx - zz
	out[6] = zy + wx
	out[7] = 0

	out[8] = zx + wy
	out[9] = zy - wx
	out[10] = 1 - xx - yy
	out[11] = 0

	out[12] = 0
	out[13] = 0
	out[14] = 0
	out[15] = 1

	return out
}

// Mat4Frustum generates a frustum matrix with the given bounds
func Mat4Frustum(out []float64, left, right, bottom, top, near, far float64) []float64 {
	rl := 1. / (right - left)
	tb := 1. / (top - bottom)
	nf := 1. / (near - far)
	out[0] = near * 2 * rl
	out[1] = 0
	out[2] = 0
	out[3] = 0
	out[4] = 0
	out[5] = near * 2 * tb
	out[6] = 0
	out[7] = 0
	out[8] = (right + left) * rl
	out[9] = (top + bottom) * tb
	out[10] = (far + near) * nf
	out[11] = -1
	out[12] = 0
	out[13] = 0
	out[14] = far * near * 2 * nf
	out[15] = 0
	return out
}

// Mat4Perspective generates a perspective projection matrix with the given bounds.
func Mat4Perspective(out []float64, fovy, aspect, near, far float64) []float64 {
	f := 1.0 / math.Tan(fovy/2)
	out[0] = f / aspect
	out[1] = 0
	out[2] = 0
	out[3] = 0
	out[4] = 0
	out[5] = f
	out[6] = 0
	out[7] = 0
	out[8] = 0
	out[9] = 0
	out[11] = -1
	out[12] = 0
	out[13] = 0
	out[15] = 0
	if !math.IsInf(far, 1) {
		nf := 1 / (near - far)
		out[10] = (far + near) * nf
		out[14] = 2 * far * near * nf
	} else {
		out[10] = -1
		out[14] = -2 * near
	}
	return out
}

// Fov represent Field of View
type Fov struct {
	UpDegrees    float64
	DownDegrees  float64
	LeftDegrees  float64
	RightDegrees float64
}

// Mat4PerspectiveFromFieldOfView generates a perspective projection matrix with the given field of view.
// This is primarily useful for generating projection matrices to be used
// with the still experiemental WebVR API.
func Mat4PerspectiveFromFieldOfView(out []float64, fov *Fov, near, far float64) []float64 {
	upTan := math.Tan((fov.UpDegrees * math.Pi) / 180.0)
	downTan := math.Tan((fov.DownDegrees * math.Pi) / 180.0)
	leftTan := math.Tan((fov.LeftDegrees * math.Pi) / 180.0)
	rightTan := math.Tan((fov.RightDegrees * math.Pi) / 180.0)
	xScale := 2.0 / (leftTan + rightTan)
	yScale := 2.0 / (upTan + downTan)

	out[0] = xScale
	out[1] = 0.0
	out[2] = 0.0
	out[3] = 0.0
	out[4] = 0.0
	out[5] = yScale
	out[6] = 0.0
	out[7] = 0.0
	out[8] = -((leftTan - rightTan) * xScale * 0.5)
	out[9] = (upTan - downTan) * yScale * 0.5
	out[10] = far / (near - far)
	out[11] = -1.0
	out[12] = 0.0
	out[13] = 0.0
	out[14] = (far * near) / (near - far)
	out[15] = 0.0
	return out
}

// Mat4Ortho generates a orthogonal projection matrix with the given bounds
func Mat4Ortho(out []float64, left, right, bottom, top, near, far float64) []float64 {
	lr := 1 / (left - right)
	bt := 1 / (bottom - top)
	nf := 1 / (near - far)
	out[0] = -2 * lr
	out[1] = 0
	out[2] = 0
	out[3] = 0
	out[4] = 0
	out[5] = -2 * bt
	out[6] = 0
	out[7] = 0
	out[8] = 0
	out[9] = 0
	out[10] = 2 * nf
	out[11] = 0
	out[12] = (left + right) * lr
	out[13] = (top + bottom) * bt
	out[14] = (far + near) * nf
	out[15] = 1
	return out
}

// Mat4LookAt generates a look-at matrix with the given eye position, focal point, and up axis.
// If you want a matrix that actually makes an object look at another object, you should use targetTo instead.
func Mat4LookAt(out, eye, center, up []float64) []float64 {
	eyex := eye[0]
	eyey := eye[1]
	eyez := eye[2]
	upx := up[0]
	upy := up[1]
	upz := up[2]
	centerx := center[0]
	centery := center[1]
	centerz := center[2]

	if math.Abs(eyex-centerx) < Epsilon && math.Abs(eyey-centery) < Epsilon && math.Abs(eyez-centerz) < Epsilon {
		return Mat4Identity(out)
	}

	z0 := eyex - centerx
	z1 := eyey - centery
	z2 := eyez - centerz

	len := 1. / hypot(z0, z1, z2)
	z0 *= len
	z1 *= len
	z2 *= len

	x0 := upy*z2 - upz*z1
	x1 := upz*z0 - upx*z2
	x2 := upx*z1 - upy*z0
	len = hypot(x0, x1, x2)
	if len == 0. {
		x0 = 0
		x1 = 0
		x2 = 0
	} else {
		len = 1 / len
		x0 *= len
		x1 *= len
		x2 *= len
	}

	y0 := z1*x2 - z2*x1
	y1 := z2*x0 - z0*x2
	y2 := z0*x1 - z1*x0

	len = hypot(y0, y1, y2)
	if len == 0. {
		y0 = 0
		y1 = 0
		y2 = 0
	} else {
		len = 1 / len
		y0 *= len
		y1 *= len
		y2 *= len
	}

	out[0] = x0
	out[1] = y0
	out[2] = z0
	out[3] = 0
	out[4] = x1
	out[5] = y1
	out[6] = z1
	out[7] = 0
	out[8] = x2
	out[9] = y2
	out[10] = z2
	out[11] = 0
	out[12] = -(x0*eyex + x1*eyey + x2*eyez)
	out[13] = -(y0*eyex + y1*eyey + y2*eyez)
	out[14] = -(z0*eyex + z1*eyey + z2*eyez)
	out[15] = 1

	return out
}

// Mat4TargetTo generates a matrix that makes something look at something else.
func Mat4TargetTo(out, eye, target, up []float64) []float64 {
	eyex := eye[0]
	eyey := eye[1]
	eyez := eye[2]
	upx := up[0]
	upy := up[1]
	upz := up[2]

	z0 := eyex - target[0]
	z1 := eyey - target[1]
	z2 := eyez - target[2]

	len := z0*z0 + z1*z1 + z2*z2
	if len > 0 {
		len = 1 / math.Sqrt(len)
		z0 *= len
		z1 *= len
		z2 *= len
	}

	x0 := upy*z2 - upz*z1
	x1 := upz*z0 - upx*z2
	x2 := upx*z1 - upy*z0

	len = x0*x0 + x1*x1 + x2*x2
	if len > 0 {
		len = 1 / math.Sqrt(len)
		x0 *= len
		x1 *= len
		x2 *= len
	}

	out[0] = x0
	out[1] = x1
	out[2] = x2
	out[3] = 0
	out[4] = z1*x2 - z2*x1
	out[5] = z2*x0 - z0*x2
	out[6] = z0*x1 - z1*x0
	out[7] = 0
	out[8] = z0
	out[9] = z1
	out[10] = z2
	out[11] = 0
	out[12] = eyex
	out[13] = eyey
	out[14] = eyez
	out[15] = 1
	return out
}

// Mat4Str returns a string representation of a mat4
func Mat4Str(a []float64) string {
	return fmt.Sprintf("mat4(%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v)",
		a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8],
		a[9], a[10], a[11], a[12], a[13], a[14], a[15])
}

// Mat4Frob returns Frobenius norm of a mat4
func Mat4Frob(a []float64) float64 {
	return hypot(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15])
}

// Mat4Add adds two mat4's
func Mat4Add(out, a, b []float64) []float64 {
	out[0] = a[0] + b[0]
	out[1] = a[1] + b[1]
	out[2] = a[2] + b[2]
	out[3] = a[3] + b[3]
	out[4] = a[4] + b[4]
	out[5] = a[5] + b[5]
	out[6] = a[6] + b[6]
	out[7] = a[7] + b[7]
	out[8] = a[8] + b[8]
	out[9] = a[9] + b[9]
	out[10] = a[10] + b[10]
	out[11] = a[11] + b[11]
	out[12] = a[12] + b[12]
	out[13] = a[13] + b[13]
	out[14] = a[14] + b[14]
	out[15] = a[15] + b[15]
	return out
}

// Mat4Subtract subtracts matrix b from matrix a
func Mat4Subtract(out, a, b []float64) []float64 {
	out[0] = a[0] - b[0]
	out[1] = a[1] - b[1]
	out[2] = a[2] - b[2]
	out[3] = a[3] - b[3]
	out[4] = a[4] - b[4]
	out[5] = a[5] - b[5]
	out[6] = a[6] - b[6]
	out[7] = a[7] - b[7]
	out[8] = a[8] - b[8]
	out[9] = a[9] - b[9]
	out[10] = a[10] - b[10]
	out[11] = a[11] - b[11]
	out[12] = a[12] - b[12]
	out[13] = a[13] - b[13]
	out[14] = a[14] - b[14]
	out[15] = a[15] - b[15]
	return out
}

// Mat4MultiplyScalar multiply each element of the matrix by a scalar.
func Mat4MultiplyScalar(out, a []float64, b float64) []float64 {
	out[0] = a[0] * b
	out[1] = a[1] * b
	out[2] = a[2] * b
	out[3] = a[3] * b
	out[4] = a[4] * b
	out[5] = a[5] * b
	out[6] = a[6] * b
	out[7] = a[7] * b
	out[8] = a[8] * b
	out[9] = a[9] * b
	out[10] = a[10] * b
	out[11] = a[11] * b
	out[12] = a[12] * b
	out[13] = a[13] * b
	out[14] = a[14] * b
	out[15] = a[15] * b
	return out
}

// Mat4MultiplyScalarAndAdd adds two mat4's after multiplying each element of the second operand by a scalar value.
func Mat4MultiplyScalarAndAdd(out, a, b []float64, scale float64) []float64 {
	out[0] = a[0] + b[0]*scale
	out[1] = a[1] + b[1]*scale
	out[2] = a[2] + b[2]*scale
	out[3] = a[3] + b[3]*scale
	out[4] = a[4] + b[4]*scale
	out[5] = a[5] + b[5]*scale
	out[6] = a[6] + b[6]*scale
	out[7] = a[7] + b[7]*scale
	out[8] = a[8] + b[8]*scale
	out[9] = a[9] + b[9]*scale
	out[10] = a[10] + b[10]*scale
	out[11] = a[11] + b[11]*scale
	out[12] = a[12] + b[12]*scale
	out[13] = a[13] + b[13]*scale
	out[14] = a[14] + b[14]*scale
	out[15] = a[15] + b[15]*scale
	return out
}

// Mat4ExactEquals returns whether or not the matrices have exactly the same elements in the same position (when compared with ===)
func Mat4ExactEquals(a, b []float64) bool {
	return a[0] == b[0] &&
		a[1] == b[1] &&
		a[2] == b[2] &&
		a[3] == b[3] &&
		a[4] == b[4] &&
		a[5] == b[5] &&
		a[6] == b[6] &&
		a[7] == b[7] &&
		a[8] == b[8] &&
		a[9] == b[9] &&
		a[10] == b[10] &&
		a[11] == b[11] &&
		a[12] == b[12] &&
		a[13] == b[13] &&
		a[14] == b[14] &&
		a[15] == b[15]
}

// Mat4Equals returns whether or not the matrices have approximately the same elements in the same position.
func Mat4Equals(a, b []float64) bool {
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]
	a4 := a[4]
	a5 := a[5]
	a6 := a[6]
	a7 := a[7]
	a8 := a[8]
	a9 := a[9]
	a10 := a[10]
	a11 := a[11]
	a12 := a[12]
	a13 := a[13]
	a14 := a[14]
	a15 := a[15]

	b0 := b[0]
	b1 := b[1]
	b2 := b[2]
	b3 := b[3]
	b4 := b[4]
	b5 := b[5]
	b6 := b[6]
	b7 := b[7]
	b8 := b[8]
	b9 := b[9]
	b10 := b[10]
	b11 := b[11]
	b12 := b[12]
	b13 := b[13]
	b14 := b[14]
	b15 := b[15]

	return equals(a0, b0) &&
		equals(a1, b1) &&
		equals(a2, b2) &&
		equals(a3, b3) &&
		equals(a4, b4) &&
		equals(a5, b5) &&
		equals(a6, b6) &&
		equals(a7, b7) &&
		equals(a8, b8) &&
		equals(a9, b9) &&
		equals(a10, b10) &&
		equals(a11, b11) &&
		equals(a12, b12) &&
		equals(a13, b13) &&
		equals(a14, b14) &&
		equals(a15, b15)
}

// Mat4Mul alias for Mat4Multiply
var Mat4Mul = Mat4Multiply

// Mat4Sub alias for Mat4Subtract
var Mat4Sub = Mat4Subtract
