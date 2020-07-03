package glmatrix

import (
	"fmt"
	"math"
)

// Quat2Create creates a new identity dual quat
func Quat2Create() []float64 {
	return []float64{
		0, 0, 0, 1,
		0, 0, 0, 0,
	}
}

// Quat2Clone creates a new quat initialized with values from an existing quaternion
func Quat2Clone(a []float64) []float64 {
	return []float64{
		a[0], a[1], a[2], a[3],
		a[4], a[5], a[6], a[7],
	}
}

// Quat2FromValues creates a new dual quat initialized with the given values
func Quat2FromValues(x1, y1, z1, w1, x2, y2, z2, w2 float64) []float64 {
	return []float64{
		x1, y1, z1, w1,
		x2, y2, z2, w2,
	}
}

// Quat2FromRotationTranslationValues creates a new dual quat from the given values (quat and translation)
func Quat2FromRotationTranslationValues(x1, y1, z1, w1, x2, y2, z2 float64) []float64 {
	dq := []float64{
		x1, y1, z1, w1,
		0, 0, 0, 0,
	}
	ax := x2 * 0.5
	ay := y2 * 0.5
	az := z2 * 0.5
	dq[4] = ax*w1 + ay*z1 - az*y1
	dq[5] = ay*w1 + az*x1 - ax*z1
	dq[6] = az*w1 + ax*y1 - ay*x1
	dq[7] = -ax*x1 - ay*y1 - az*z1
	return dq
}

// Quat2FromRotationTranslation creates a dual quat from a quaternion and a translation
func Quat2FromRotationTranslation(out, q, t []float64) []float64 {
	ax := t[0] * 0.5
	ay := t[1] * 0.5
	az := t[2] * 0.5
	bx := q[0]
	by := q[1]
	bz := q[2]
	bw := q[3]
	out[0] = bx
	out[1] = by
	out[2] = bz
	out[3] = bw
	out[4] = ax*bw + ay*bz - az*by
	out[5] = ay*bw + az*bx - ax*bz
	out[6] = az*bw + ax*by - ay*bx
	out[7] = -ax*bx - ay*by - az*bz
	return out
}

// Quat2FromTranslation creates a dual quat from a translation
func Quat2FromTranslation(out, t []float64) []float64 {
	out[0] = 0
	out[1] = 0
	out[2] = 0
	out[3] = 1
	out[4] = t[0] * 0.5
	out[5] = t[1] * 0.5
	out[6] = t[2] * 0.5
	out[7] = 0
	return out
}

// Quat2FromRotation creates a dual quat from a quaternion
func Quat2FromRotation(out, q []float64) []float64 {
	out[0] = q[0]
	out[1] = q[1]
	out[2] = q[2]
	out[3] = q[3]
	out[4] = 0
	out[5] = 0
	out[6] = 0
	out[7] = 0
	return out
}

// Quat2FromMat4 creates a new dual quat from a matrix (4x4)
func Quat2FromMat4(out, a []float64) []float64 {
	//TODO Optimize this
	outer := QuatCreate()
	Mat4GetRotation(outer, a)
	t := []float64{0, 0, 0}
	Mat4GetTranslation(t, a)
	Quat2FromRotationTranslation(out, outer, t)
	return out
}

// Quat2Copy copy the values from one dual quat to another
func Quat2Copy(out, a []float64) []float64 {
	out[0] = a[0]
	out[1] = a[1]
	out[2] = a[2]
	out[3] = a[3]
	out[4] = a[4]
	out[5] = a[5]
	out[6] = a[6]
	out[7] = a[7]
	return out
}

// Quat2Identity set a dual quat to the identity dual quaternion
func Quat2Identity(out []float64) []float64 {
	out[0] = 0
	out[1] = 0
	out[2] = 0
	out[3] = 1
	out[4] = 0
	out[5] = 0
	out[6] = 0
	out[7] = 0
	return out
}

// Quat2Set set the components of a dual quat to the given values
func Quat2Set(out []float64, x1, y1, z1, w1, x2, y2, z2, w2 float64) []float64 {
	out[0] = x1
	out[1] = y1
	out[2] = z1
	out[3] = w1

	out[4] = x2
	out[5] = y2
	out[6] = z2
	out[7] = w2
	return out
}

// Quat2GetReal gets the real part of a dual quat
var Quat2GetReal = QuatCopy

// Quat2GetDual gets the dual part of a dual quat
func Quat2GetDual(out, a []float64) []float64 {
	out[0] = a[4]
	out[1] = a[5]
	out[2] = a[6]
	out[3] = a[7]
	return out
}

// Quat2SetReal set the real component of a dual quat to the given quaternion
var Quat2SetReal = QuatCopy

// Quat2SetDual set the dual component of a dual quat to the given quaternion
func Quat2SetDual(out, q []float64) []float64 {
	out[4] = q[0]
	out[5] = q[1]
	out[6] = q[2]
	out[7] = q[3]
	return out
}

// Quat2GetTranslation gets the translation of a normalized dual quat
func Quat2GetTranslation(out, a []float64) []float64 {
	ax := a[4]
	ay := a[5]
	az := a[6]
	aw := a[7]
	bx := -a[0]
	by := -a[1]
	bz := -a[2]
	bw := a[3]
	out[0] = (ax*bw + aw*bx + ay*bz - az*by) * 2
	out[1] = (ay*bw + aw*by + az*bx - ax*bz) * 2
	out[2] = (az*bw + aw*bz + ax*by - ay*bx) * 2
	return out
}

// Quat2Translate translates a dual quat by the given vector
func Quat2Translate(out, a, v []float64) []float64 {
	ax1 := a[0]
	ay1 := a[1]
	az1 := a[2]
	aw1 := a[3]
	bx1 := v[0] * 0.5
	by1 := v[1] * 0.5
	bz1 := v[2] * 0.5
	ax2 := a[4]
	ay2 := a[5]
	az2 := a[6]
	aw2 := a[7]
	out[0] = ax1
	out[1] = ay1
	out[2] = az1
	out[3] = aw1
	out[4] = aw1*bx1 + ay1*bz1 - az1*by1 + ax2
	out[5] = aw1*by1 + az1*bx1 - ax1*bz1 + ay2
	out[6] = aw1*bz1 + ax1*by1 - ay1*bx1 + az2
	out[7] = -ax1*bx1 - ay1*by1 - az1*bz1 + aw2
	return out
}

// Quat2RotateX rotates a dual quat around the X axis
func Quat2RotateX(out, a []float64, rad float64) []float64 {
	bx := -a[0]
	by := -a[1]
	bz := -a[2]
	bw := a[3]
	ax := a[4]
	ay := a[5]
	az := a[6]
	aw := a[7]
	ax1 := ax*bw + aw*bx + ay*bz - az*by
	ay1 := ay*bw + aw*by + az*bx - ax*bz
	az1 := az*bw + aw*bz + ax*by - ay*bx
	aw1 := aw*bw - ax*bx - ay*by - az*bz
	QuatRotateX(out, a, rad)
	bx = out[0]
	by = out[1]
	bz = out[2]
	bw = out[3]
	out[4] = ax1*bw + aw1*bx + ay1*bz - az1*by
	out[5] = ay1*bw + aw1*by + az1*bx - ax1*bz
	out[6] = az1*bw + aw1*bz + ax1*by - ay1*bx
	out[7] = aw1*bw - ax1*bx - ay1*by - az1*bz
	return out
}

// Quat2RotateY rotates a dual quat around the Y axis
func Quat2RotateY(out, a []float64, rad float64) []float64 {
	bx := -a[0]
	by := -a[1]
	bz := -a[2]
	bw := a[3]
	ax := a[4]
	ay := a[5]
	az := a[6]
	aw := a[7]
	ax1 := ax*bw + aw*bx + ay*bz - az*by
	ay1 := ay*bw + aw*by + az*bx - ax*bz
	az1 := az*bw + aw*bz + ax*by - ay*bx
	aw1 := aw*bw - ax*bx - ay*by - az*bz
	QuatRotateY(out, a, rad)
	bx = out[0]
	by = out[1]
	bz = out[2]
	bw = out[3]
	out[4] = ax1*bw + aw1*bx + ay1*bz - az1*by
	out[5] = ay1*bw + aw1*by + az1*bx - ax1*bz
	out[6] = az1*bw + aw1*bz + ax1*by - ay1*bx
	out[7] = aw1*bw - ax1*bx - ay1*by - az1*bz
	return out
}

// Quat2RotateZ rotates a dual quat around the Z axis
func Quat2RotateZ(out, a []float64, rad float64) []float64 {
	bx := -a[0]
	by := -a[1]
	bz := -a[2]
	bw := a[3]
	ax := a[4]
	ay := a[5]
	az := a[6]
	aw := a[7]
	ax1 := ax*bw + aw*bx + ay*bz - az*by
	ay1 := ay*bw + aw*by + az*bx - ax*bz
	az1 := az*bw + aw*bz + ax*by - ay*bx
	aw1 := aw*bw - ax*bx - ay*by - az*bz
	QuatRotateZ(out, a, rad)
	bx = out[0]
	by = out[1]
	bz = out[2]
	bw = out[3]
	out[4] = ax1*bw + aw1*bx + ay1*bz - az1*by
	out[5] = ay1*bw + aw1*by + az1*bx - ax1*bz
	out[6] = az1*bw + aw1*bz + ax1*by - ay1*bx
	out[7] = aw1*bw - ax1*bx - ay1*by - az1*bz
	return out
}

// Quat2RotateByQuatAppend rotates a dual quat by a given quaternion (a * q)
func Quat2RotateByQuatAppend(out, a, q []float64) []float64 {
	qx := q[0]
	qy := q[1]
	qz := q[2]
	qw := q[3]
	ax := a[0]
	ay := a[1]
	az := a[2]
	aw := a[3]

	out[0] = ax*qw + aw*qx + ay*qz - az*qy
	out[1] = ay*qw + aw*qy + az*qx - ax*qz
	out[2] = az*qw + aw*qz + ax*qy - ay*qx
	out[3] = aw*qw - ax*qx - ay*qy - az*qz
	ax = a[4]
	ay = a[5]
	az = a[6]
	aw = a[7]
	out[4] = ax*qw + aw*qx + ay*qz - az*qy
	out[5] = ay*qw + aw*qy + az*qx - ax*qz
	out[6] = az*qw + aw*qz + ax*qy - ay*qx
	out[7] = aw*qw - ax*qx - ay*qy - az*qz
	return out
}

// Quat2RotateByQuatPrepend rotates a dual quat by a given quaternion (q * a)
func Quat2RotateByQuatPrepend(out, q, a []float64) []float64 {
	qx := q[0]
	qy := q[1]
	qz := q[2]
	qw := q[3]
	bx := a[0]
	by := a[1]
	bz := a[2]
	bw := a[3]

	out[0] = qx*bw + qw*bx + qy*bz - qz*by
	out[1] = qy*bw + qw*by + qz*bx - qx*bz
	out[2] = qz*bw + qw*bz + qx*by - qy*bx
	out[3] = qw*bw - qx*bx - qy*by - qz*bz
	bx = a[4]
	by = a[5]
	bz = a[6]
	bw = a[7]
	out[4] = qx*bw + qw*bx + qy*bz - qz*by
	out[5] = qy*bw + qw*by + qz*bx - qx*bz
	out[6] = qz*bw + qw*bz + qx*by - qy*bx
	out[7] = qw*bw - qx*bx - qy*by - qz*bz
	return out
}

// Quat2RotateAroundAxis rotates a dual quat around a given axis. Does the normalisation automatically
func Quat2RotateAroundAxis(out, a, axis []float64, rad float64) []float64 {
	//Special case for rad = 0
	if equals(rad, 0) {
		return Quat2Copy(out, a)
	}
	axisLength := hypot(axis[0], axis[1], axis[2])

	rad = rad * 0.5
	s := math.Sin(rad)
	bx := (s * axis[0]) / axisLength
	by := (s * axis[1]) / axisLength
	bz := (s * axis[2]) / axisLength
	bw := math.Cos(rad)

	ax1 := a[0]
	ay1 := a[1]
	az1 := a[2]
	aw1 := a[3]
	out[0] = ax1*bw + aw1*bx + ay1*bz - az1*by
	out[1] = ay1*bw + aw1*by + az1*bx - ax1*bz
	out[2] = az1*bw + aw1*bz + ax1*by - ay1*bx
	out[3] = aw1*bw - ax1*bx - ay1*by - az1*bz

	ax := a[4]
	ay := a[5]
	az := a[6]
	aw := a[7]
	out[4] = ax*bw + aw*bx + ay*bz - az*by
	out[5] = ay*bw + aw*by + az*bx - ax*bz
	out[6] = az*bw + aw*bz + ax*by - ay*bx
	out[7] = aw*bw - ax*bx - ay*by - az*bz

	return out
}

// Quat2Add adds two dual quat's
func Quat2Add(out, a, b []float64) []float64 {
	out[0] = a[0] + b[0]
	out[1] = a[1] + b[1]
	out[2] = a[2] + b[2]
	out[3] = a[3] + b[3]
	out[4] = a[4] + b[4]
	out[5] = a[5] + b[5]
	out[6] = a[6] + b[6]
	out[7] = a[7] + b[7]
	return out
}

// Quat2Multiply multiplies two dual quat's
func Quat2Multiply(out, a, b []float64) []float64 {
	ax0 := a[0]
	ay0 := a[1]
	az0 := a[2]
	aw0 := a[3]
	bx1 := b[4]
	by1 := b[5]
	bz1 := b[6]
	bw1 := b[7]
	ax1 := a[4]
	ay1 := a[5]
	az1 := a[6]
	aw1 := a[7]
	bx0 := b[0]
	by0 := b[1]
	bz0 := b[2]
	bw0 := b[3]
	out[0] = ax0*bw0 + aw0*bx0 + ay0*bz0 - az0*by0
	out[1] = ay0*bw0 + aw0*by0 + az0*bx0 - ax0*bz0
	out[2] = az0*bw0 + aw0*bz0 + ax0*by0 - ay0*bx0
	out[3] = aw0*bw0 - ax0*bx0 - ay0*by0 - az0*bz0
	out[4] =
		ax0*bw1 +
			aw0*bx1 +
			ay0*bz1 -
			az0*by1 +
			ax1*bw0 +
			aw1*bx0 +
			ay1*bz0 -
			az1*by0
	out[5] =
		ay0*bw1 +
			aw0*by1 +
			az0*bx1 -
			ax0*bz1 +
			ay1*bw0 +
			aw1*by0 +
			az1*bx0 -
			ax1*bz0
	out[6] =
		az0*bw1 +
			aw0*bz1 +
			ax0*by1 -
			ay0*bx1 +
			az1*bw0 +
			aw1*bz0 +
			ax1*by0 -
			ay1*bx0
	out[7] =
		aw0*bw1 -
			ax0*bx1 -
			ay0*by1 -
			az0*bz1 +
			aw1*bw0 -
			ax1*bx0 -
			ay1*by0 -
			az1*bz0
	return out
}

// Quat2Mul alias for Quat2Multiply
var Quat2Mul = Quat2Multiply

// Quat2Scale scales a dual quat by a scalar number
func Quat2Scale(out, a []float64, b float64) []float64 {
	out[0] = a[0] * b
	out[1] = a[1] * b
	out[2] = a[2] * b
	out[3] = a[3] * b
	out[4] = a[4] * b
	out[5] = a[5] * b
	out[6] = a[6] * b
	out[7] = a[7] * b
	return out
}

// Quat2Dot calculates the dot product of two dual quat's (The dot product of the real parts)
var Quat2Dot = QuatDot

// Quat2Lerp performs a linear interpolation between two dual quats's
// NOTE: The resulting dual quaternions won't always be normalized (The error is most noticeable when t = 0.5)
func Quat2Lerp(out, a, b []float64, t float64) []float64 {
	mt := 1 - t
	if Quat2Dot(a, b) < 0 {
		t = -t
	}

	out[0] = a[0]*mt + b[0]*t
	out[1] = a[1]*mt + b[1]*t
	out[2] = a[2]*mt + b[2]*t
	out[3] = a[3]*mt + b[3]*t
	out[4] = a[4]*mt + b[4]*t
	out[5] = a[5]*mt + b[5]*t
	out[6] = a[6]*mt + b[6]*t
	out[7] = a[7]*mt + b[7]*t

	return out
}

// Quat2Invert calculates the inverse of a dual quat. If they are normalized, conjugate is cheaper
func Quat2Invert(out, a []float64) []float64 {
	sqlen := Quat2SquaredLength(a)
	out[0] = -a[0] / sqlen
	out[1] = -a[1] / sqlen
	out[2] = -a[2] / sqlen
	out[3] = a[3] / sqlen
	out[4] = -a[4] / sqlen
	out[5] = -a[5] / sqlen
	out[6] = -a[6] / sqlen
	out[7] = a[7] / sqlen
	return out
}

// Quat2Conjugate calculates the conjugate of a dual quat
// If the dual quaternion is normalized, this function is faster than quat2.inverse and produces the same result.
func Quat2Conjugate(out, a []float64) []float64 {
	out[0] = -a[0]
	out[1] = -a[1]
	out[2] = -a[2]
	out[3] = a[3]
	out[4] = -a[4]
	out[5] = -a[5]
	out[6] = -a[6]
	out[7] = a[7]
	return out
}

// Quat2Length calculates the length of a dual quat
var Quat2Length = QuatLength

// Quat2Len alias for Quat2Length
var Quat2Len = Quat2Length

// Quat2SquaredLength calculates the squared length of a dual quat
var Quat2SquaredLength = QuatSquaredLength

// Quat2SqrLen alias for Quat2SquaredLength
var Quat2SqrLen = Quat2SquaredLength

// Quat2Normalize normalize a dual quat
func Quat2Normalize(out, a []float64) []float64 {
	magnitude := Quat2SquaredLength(a)
	if magnitude > 0 {
		magnitude = math.Sqrt(magnitude)

		a0 := a[0] / magnitude
		a1 := a[1] / magnitude
		a2 := a[2] / magnitude
		a3 := a[3] / magnitude

		b0 := a[4]
		b1 := a[5]
		b2 := a[6]
		b3 := a[7]

		dotAB := a0*b0 + a1*b1 + a2*b2 + a3*b3

		out[0] = a0
		out[1] = a1
		out[2] = a2
		out[3] = a3

		out[4] = (b0 - a0*dotAB) / magnitude
		out[5] = (b1 - a1*dotAB) / magnitude
		out[6] = (b2 - a2*dotAB) / magnitude
		out[7] = (b3 - a3*dotAB) / magnitude
	}
	return out
}

// Quat2Str returns a string representation of a dual quatenion
func Quat2Str(a []float64) string {
	return fmt.Sprintf("quat2(%v, %v, %v, %v, %v, %v, %v, %v)", a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7])
}

// Quat2ExactEquals returns whether or not the dual quaternions have exactly the same elements in the same position (when compared with ==)
func Quat2ExactEquals(a, b []float64) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3] && a[4] == b[4] && a[5] == b[5] && a[6] == b[6] && a[7] == b[7]
}

// Quat2Equals returns whether or not the dual quaternions have approximately the same elements in the same position.
func Quat2Equals(a, b []float64) bool {
	return equals(a[0], b[0]) &&
		equals(a[1], b[1]) &&
		equals(a[2], b[2]) &&
		equals(a[3], b[3]) &&
		equals(a[4], b[4]) &&
		equals(a[5], b[5]) &&
		equals(a[6], b[6]) &&
		equals(a[7], b[7])
}
