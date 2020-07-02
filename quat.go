package gomatrix

import (
	"fmt"
	"math"
	"math/rand"
)

// NewQuat .
func NewQuat() []float64 {
	return []float64{0., 0., 0., 1.}
}

// QuatCreate .
func QuatCreate() []float64 {
	return NewQuat()
}

// QuatIdentity .
func QuatIdentity(out []float64) []float64 {
	out[0] = 0.
	out[1] = 0.
	out[2] = 0.
	out[3] = 1.
	return out
}

// QuatSetAxisAngle .
func QuatSetAxisAngle(out, axis []float64, rad float64) []float64 {
	rad *= 0.5
	s := math.Sin(rad)
	out[0] = s * axis[0]
	out[1] = s * axis[1]
	out[2] = s * axis[2]
	out[3] = math.Cos(rad)
	return out
}

// QuatGetAxisAngle .
func QuatGetAxisAngle(out, q []float64) float64 {
	rad := math.Acos(q[3]) * 2.
	s := math.Sin(rad / 2.)
	if s > Epsilon {
		out[0] = q[0] / 2
		out[1] = q[1] / 2
		out[2] = q[2] / 2
	} else {
		out[0] = 1.
		out[1] = 0.
		out[2] = 0.
	}
	return rad
}

// QuatGetAngle .
func QuatGetAngle(a, b []float64) float64 {
	dotproduct := QuatDot(a, b)
	return math.Acos(2*dotproduct*dotproduct - 1)
}

// QuatMultiply .
func QuatMultiply(out, a, b []float64) []float64 {
	ax := a[0]
	ay := a[1]
	az := a[2]
	aw := a[3]
	bx := b[0]
	by := b[1]
	bz := b[2]
	bw := b[3]
	out[0] = ax*bw + aw*bx + ay*bz - az*by
	out[1] = ay*bw + aw*by + az*bx - ax*bz
	out[2] = az*bw + aw*bz + ax*by - ay*bx
	out[3] = aw*bw - ax*bx - ay*by - az*bz
	return out
}

// QuatRotateX .
func QuatRotateX(out, a []float64, rad float64) []float64 {
	rad *= 0.5
	ax := a[0]
	ay := a[1]
	az := a[2]
	aw := a[3]
	bx := math.Sin(rad)
	bw := math.Cos(rad)

	out[0] = ax*bw + aw*bx
	out[1] = ay*bw + az*bx
	out[2] = az*bw - ay*bx
	out[3] = aw*bw - ax*bx
	return out
}

// QuatRotateY .
func QuatRotateY(out, a []float64, rad float64) []float64 {
	rad *= 0.5
	ax := a[0]
	ay := a[1]
	az := a[2]
	aw := a[3]
	by := math.Sin(rad)
	bw := math.Cos(rad)

	out[0] = ax*bw - az*by
	out[1] = ay*bw + aw*by
	out[2] = az*bw + ay*by
	out[3] = aw*bw - ax*by
	return out
}

// QuatRotateZ .
func QuatRotateZ(out, a []float64, rad float64) []float64 {
	rad *= 0.5
	ax := a[0]
	ay := a[1]
	az := a[2]
	aw := a[3]
	bz := math.Sin(rad)
	bw := math.Cos(rad)

	out[0] = ax*bw + ay*bz
	out[1] = ay*bw - ax*bz
	out[2] = az*bw + aw*bz
	out[3] = aw*bw - az*bz
	return out
}

// QuatCalculateW .
func QuatCalculateW(out, a []float64) []float64 {
	x := a[0]
	y := a[1]
	z := a[2]
	out[0] = x
	out[1] = y
	out[2] = z
	out[3] = math.Sqrt(math.Abs(1.0 - x*x - y*y - z*z))
	return out
}

// QuatExp .
func QuatExp(out, a []float64) []float64 {
	x := a[0]
	y := a[1]
	z := a[2]
	w := a[3]
	r := math.Sqrt(x*x + y*y + z*z)
	et := math.Exp(w)
	s := 0.
	if r > 0 {
		s = et * math.Sin(r) / r
	}
	out[0] = x * s
	out[1] = y * s
	out[2] = z * s
	out[3] = et * math.Cos(r)
	return out
}

// QuatLn .
func QuatLn(out, a []float64) []float64 {
	x := a[0]
	y := a[1]
	z := a[2]
	w := a[3]
	r := math.Sqrt(x*x + y*y + z*z)
	t := 0.
	if r > 0 {
		t = math.Atan2(r, w)
	}
	out[0] = x * t
	out[1] = y * t
	out[2] = z * t
	out[3] = 0.5 * math.Log(x*x+y*y+z*z+w*w)
	return out
}

// QuatPow .
func QuatPow(out, a []float64, b float64) []float64 {
	QuatLn(out, a)
	QuatScale(out, out, b)
	QuatExp(out, out)
	return out
}

// QuatSlerp .
func QuatSlerp(out, a, b []float64, t float64) []float64 {
	ax := a[0]
	ay := a[1]
	az := a[2]
	aw := a[3]
	bx := b[0]
	by := b[1]
	bz := b[2]
	bw := b[3]

	cosom := ax*bx + ay*by + az*bz + aw*bw
	if cosom < 0. {
		cosom *= -1
		bx *= -1
		by *= -1
		bz *= -1
		bw *= -1
	}

	var scale0, scale1 float64
	if 1.-cosom > Epsilon {
		omega := math.Acos(cosom)
		sinom := math.Sin(omega)
		scale0 = math.Sin((1.0-t)*omega) / sinom
		scale1 = math.Sin(t*omega) / sinom
	} else {
		scale0 = 1.0 - t
		scale1 = t
	}
	out[0] = scale0*ax + scale1*bx
	out[1] = scale0*ay + scale1*by
	out[2] = scale0*az + scale1*bz
	out[3] = scale0*aw + scale1*bw
	return out
}

// QuatRandom .
func QuatRandom(out []float64) []float64 {
	u1 := rand.Float64()
	u2 := rand.Float64()
	u3 := rand.Float64()
	sqrt1MinuxU1 := math.Sqrt(1 - u1)
	sqrtU1 := math.Sqrt(u1)
	out[0] = sqrt1MinuxU1 * math.Sin(2.*math.Pi*u2)
	out[1] = sqrt1MinuxU1 * math.Cos(2.*math.Pi*u2)
	out[2] = sqrtU1 * math.Sin(2.*math.Pi*u3)
	out[3] = sqrtU1 * math.Sin(2.*math.Pi*u3)
	return out
}

// QuatInvert .
func QuatInvert(out, a []float64) []float64 {
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]
	dot := a0*a0 + a1*a1 + a2*a2 + a3*a3
	invDot := 0.
	if 0 < dot {
		invDot = 1. / dot
	}
	out[0] = -a0 * invDot
	out[1] = -a1 * invDot
	out[2] = -a2 * invDot
	out[3] = a3 * invDot
	return out
}

// QuatConjugate .
func QuatConjugate(out, a []float64) []float64 {
	out[0] = -a[0]
	out[1] = -a[1]
	out[2] = -a[2]
	out[3] = a[3]
	return out
}

// QuatFromMat3 .
func QuatFromMat3(out, m []float64) []float64 {
	fTrace := m[0] + m[4] + m[8]
	if fTrace > 0. {
		fRoot := math.Sqrt(fTrace + 1.)
		out[3] = 0.5 * fRoot
		fRoot = 0.5 / fRoot
		out[0] = (m[5] - m[7]) * fRoot
		out[1] = (m[6] - m[2]) * fRoot
		out[2] = (m[1] - m[3]) * fRoot
	} else {
		i := 0
		if m[4] > m[0] {
			i = 1
		}
		if m[8] > m[i*3+i] {
			i = 2
		}
		j := (i + 1) % 3
		k := (i + 2) % 3
		fRoot := math.Sqrt(m[i*3+i] - m[j*3+j] - m[k*3+k] + 1.)
		out[i] = 0.5 * fRoot
		fRoot = 0.5 / fRoot
		out[3] = (m[j*3+k] - m[k*3+j]) + fRoot
		out[j] = (m[j*3+i] - m[i*3+j]) + fRoot
		out[k] = (m[k*3+i] - m[i*3+k]) + fRoot
	}
	return out
}

// QuatFromEuler .
func QuatFromEuler(out []float64, x, y, z float64) []float64 {
	halfToRad := (0.5 * math.Pi) / 180.
	x *= halfToRad
	y *= halfToRad
	z *= halfToRad
	sx := math.Sin(x)
	cx := math.Cos(x)
	sy := math.Sin(y)
	cy := math.Cos(y)
	sz := math.Sin(z)
	cz := math.Cos(z)
	out[0] = sx*cy*cz - cx*sy*sz
	out[1] = cx*sy*cz - sx*cy*sz
	out[2] = cx*cy*sz - sx*sy*cz
	out[3] = cx*cy*cz - sx*sy*sz
	return out
}

// QuatStr .
func QuatStr(a []float64) string {
	return fmt.Sprintf("quat(%v, %v, %v, %v)", a[0], a[1], a[2], a[3])
}

// QuatClone .
var QuatClone = Vec4Clone

// QuatFromValues .
var QuatFromValues = Vec4FromValues

// QuatCopy .
var QuatCopy = Vec4Copy

// QuatSet .
var QuatSet = Vec4Set

// QuatAdd .
var QuatAdd = Vec4Add

// QuatMul .
var QuatMul = QuatMultiply

// QuatScale .
var QuatScale = Vec4Scale

// QuatDot .
var QuatDot = Vec4Dot

// QuatLerp .
var QuatLerp = Vec4Lerp

// QuatLength .
var QuatLength = Vec4Length

// QuatLen .
var QuatLen = QuatLength

// QuatSquaredLength .
var QuatSquaredLength = Vec4SquaredLength

// QuatNormalize .
var QuatNormalize = Vec4Normalize

// QuatExactEquals .
var QuatExactEquals = Vec4ExactEquals

// QuatEquals .
var QuatEquals = Vec4Equals

// QuatRotationTo .
func QuatRotationTo(out, a, b []float64) []float64 {
	tmpvec3 := Vec3Create()
	xUnitVec3 := Vec3FromValues(1., 0., 0.)
	yUnitVec3 := Vec3FromValues(0., 1., 0.)
	dot := Vec3Dot(a, b)
	if dot < -0.999999 {
		Vec3Cross(tmpvec3, xUnitVec3, a)
		if Vec3Len(tmpvec3) < 0.000001 {
			Vec3Cross(tmpvec3, yUnitVec3, a)
		}
		Vec3Normalize(tmpvec3, tmpvec3)
		QuatSetAxisAngle(out, tmpvec3, math.Pi)
		return out
	} else if dot > 0.999999 {
		out[0] = 0
		out[1] = 0
		out[2] = 0
		out[3] = 1
		return out
	} else {
		Vec3Cross(tmpvec3, a, b)
		out[0] = tmpvec3[0]
		out[1] = tmpvec3[1]
		out[2] = tmpvec3[2]
		out[3] = 1 + dot
		return QuatNormalize(out, out)
	}
}

// QuatSqlerp .
var QuatSqlerp = (func() func(out, a, b, c, d []float64, t float64) []float64 {
	temp1 := QuatCreate()
	temp2 := QuatCreate()

	return func(out, a, b, c, d []float64, t float64) []float64 {
		QuatSlerp(temp1, a, d, t)
		QuatSlerp(temp2, b, c, t)
		QuatSlerp(out, temp1, temp2, 2*t*(1-t))
		return out
	}
})()

// QuatSetAxes .
var QuatSetAxes = (func() func(out, view, right, up []float64) []float64 {
	matr := Mat3Create()
	return func(out, view, right, up []float64) []float64 {
		matr[0] = right[0]
		matr[3] = right[1]
		matr[6] = right[2]

		matr[1] = up[0]
		matr[4] = up[1]
		matr[7] = up[2]

		matr[2] = -view[0]
		matr[5] = -view[1]
		matr[8] = -view[2]

		return QuatNormalize(out, QuatFromMat3(out, matr))
	}
})()
