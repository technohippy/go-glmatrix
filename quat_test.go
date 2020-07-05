package glmatrix

import (
	"math"
	"testing"
)

var quatA = []float64{1, 2, 3, 4}
var quatB = []float64{5, 6, 7, 8}
var vec = []float64{1, 1, -1}
var id = []float64{0, 0, 0, 1}

func TestQuatSlerp(t *testing.T) {
	actual := QuatSlerp(QuatCreate(), []float64{0, 0, 0, 1}, []float64{0, 1, 0, 0}, 0.5)
	expect := []float64{0, 0.707106, 0, 0.707106}
	if !testSlice(actual, expect) {
		t.Errorf("slerp: %v", actual)
	}
}

func TestQuatPow(t *testing.T) {
	actual := QuatPow(QuatCreate(), id, 2.1)
	expect := id
	if !testSlice(actual, expect) {
		t.Errorf("pow: %v", actual)
	}
}

func TestQuatRotateX(t *testing.T) {
	vec := []float64{1, 1, -1}
	actual := QuatRotateX(QuatCreate(), id, math.Pi/2)
	Vec3TransformQuat(vec, []float64{0, 0, -1}, actual)
	expect := []float64{0, 1, 0}
	if !testSlice(vec, expect) {
		t.Errorf("rotate x: %v", actual)
	}
}

func TestQuatRotateY(t *testing.T) {
	vec := []float64{1, 1, -1}
	actual := QuatRotateY(QuatCreate(), id, math.Pi/2)
	Vec3TransformQuat(vec, []float64{0, 0, -1}, actual)
	expect := []float64{-1, 0, 0}
	if !testSlice(vec, expect) {
		t.Errorf("rotate y: %v", actual)
	}
}

func TestQuatRotateZ(t *testing.T) {
	vec := []float64{1, 1, -1}
	actual := QuatRotateZ(QuatCreate(), id, math.Pi/2)
	Vec3TransformQuat(vec, []float64{0, 1, 0}, actual)
	expect := []float64{-1, 0, 0}
	if !testSlice(vec, expect) {
		t.Errorf("rotate z: %v", vec)
	}
}

func TestQuatFromMat3(t *testing.T) {
	matr := []float64{
		1, 0, 0,
		0, 0, -1,
		0, 1, 0,
	}
	actual := QuatFromMat3(QuatCreate(), matr)
	expect := []float64{-0.707106, 0, 0, 0.707106}
	if !testSlice(actual, expect) {
		t.Errorf("from mat3: %v", actual)
	}
}

func TestQuatFromEuler(t *testing.T) {
	actual := QuatFromEuler(QuatCreate(), -90, 0, 0)
	expect := []float64{-0.707106, 0, 0, 0.707106}
	if !testSlice(actual, expect) {
		t.Errorf("from euler: %v", actual)
	}
}

func TestQuatSetAxes(t *testing.T) {
	view := []float64{-1, 0, 0}
	up := []float64{0, 1, 0}
	right := []float64{0, 0, -1}
	quat := QuatSetAxes(QuatCreate(), view, right, up)
	actual := Vec3TransformQuat(Vec3Create(), []float64{1, 0, 0}, quat)
	expect := []float64{0, 0, 1}
	if !testSlice(actual, expect) {
		t.Errorf("set axes: %v", actual)
	}
}

func TestQuatRotationTo(t *testing.T) {
	actual := QuatRotationTo(QuatCreate(), []float64{0, 1, 0}, []float64{1, 0, 0})
	expect := []float64{0, 0, -0.707106, 0.707106}
	if !testSlice(actual, expect) {
		t.Errorf("rotation to: %v", actual)
	}
}

func TestQuatCreate(t *testing.T) {
	actual := QuatCreate()
	expect := []float64{0, 0, 0, 1}
	if !testSlice(actual, expect) {
		t.Errorf("create: %v", actual)
	}
}

func TestQuatClone(t *testing.T) {
	actual := QuatClone(quatA)
	expect := quatA
	if !testSlice(actual, expect) {
		t.Errorf("clone: %v", actual)
	}
	if &actual == &expect {
		t.Errorf("clone: %v", actual)
	}
}

func TestQuatFromValues(t *testing.T) {
	actual := QuatFromValues(1, 2, 3, 4)
	expect := []float64{1, 2, 3, 4}
	if !testSlice(actual, expect) {
		t.Errorf("from values: %v", actual)
	}
}

func TestQuatCopy(t *testing.T) {
	actual := QuatCopy(QuatCreate(), quatA)
	expect := []float64{1, 2, 3, 4}
	if !testSlice(actual, expect) {
		t.Errorf("copy: %v", actual)
	}
}

func TestQuatSet(t *testing.T) {
	actual := QuatSet(QuatCreate(), 1, 2, 3, 4)
	expect := []float64{1, 2, 3, 4}
	if !testSlice(actual, expect) {
		t.Errorf("set: %v", actual)
	}
}

func TestQuatIdentity(t *testing.T) {
	actual := QuatIdentity(QuatCreate())
	expect := []float64{0, 0, 0, 1}
	if !testSlice(actual, expect) {
		t.Errorf("identity: %v", actual)
	}
}

func TestQuatSetAxisAngle(t *testing.T) {
	actual := QuatSetAxisAngle(QuatCreate(), []float64{1, 0, 0}, math.Pi*0.5)
	expect := []float64{0.707106, 0, 0, 0.707106}
	if !testSlice(actual, expect) {
		t.Errorf("set axis angle: %v", actual)
	}
}

func TestQuatGetAxisAngle(t *testing.T) {
	out := QuatSetAxisAngle(QuatCreate(), []float64{1, 0, 0}, 0.7778)
	actual := QuatGetAxisAngle(vec, out)
	expect := 0.7778
	if !equals(actual, expect) {
		t.Errorf("get axis angle: %v", actual)
	}
}

func TestQuatGetAngle(t *testing.T) {
	a1 := QuatNormalize(QuatCreate(), quatA)
	a2 := QuatRotateX(QuatCreate(), a1, math.Pi/4)
	actual := QuatGetAngle(a1, a2)
	expect := math.Pi / 4
	if !equals(actual, expect) {
		t.Errorf("get angle: %v, %v", actual, expect)
	}
}

func TestQuatAdd(t *testing.T) {
	actual := QuatAdd(QuatCreate(), quatA, quatB)
	expect := []float64{6, 8, 10, 12}
	if !testSlice(actual, expect) {
		t.Errorf("add: %v", actual)
	}
}

func TestQuatMultiply(t *testing.T) {
	actual := QuatMultiply(QuatCreate(), quatA, quatB)
	expect := []float64{24, 48, 48, -6}
	if !testSlice(actual, expect) {
		t.Errorf("multiply: %v", actual)
	}
}

func TestQuatScale(t *testing.T) {
	actual := QuatScale(QuatCreate(), quatA, 2.)
	expect := []float64{2, 4, 6, 8}
	if !testSlice(actual, expect) {
		t.Errorf("scale: %v", actual)
	}
}

func TestQuatLength(t *testing.T) {
	actual := QuatLength(quatA)
	expect := 5.477225
	if !equals(actual, expect) {
		t.Errorf("length: %v", actual)
	}
	actual = QuatLen(quatA)
	if !equals(actual, expect) {
		t.Errorf("len: %v", actual)
	}
}

func TestQuatSquaredLength(t *testing.T) {
	actual := QuatSquaredLength(quatA)
	expect := 30.
	if !equals(actual, expect) {
		t.Errorf("squared length: %v", actual)
	}
	actual = QuatSqrLen(quatA)
	if !equals(actual, expect) {
		t.Errorf("sqrlen: %v", actual)
	}
}

func TestQuatNormalize(t *testing.T) {
	quatA := []float64{5, 0, 0, 0}
	actual := QuatNormalize(QuatCreate(), quatA)
	expect := []float64{1, 0, 0, 0}
	if !testSlice(actual, expect) {
		t.Errorf("normalize: %v", actual)
	}
}

func TestQuatLerp(t *testing.T) {
	actual := QuatLerp(QuatCreate(), quatA, quatB, 0.5)
	expect := []float64{3, 4, 5, 6}
	if !testSlice(actual, expect) {
		t.Errorf("lerp: %v", actual)
	}
}

/*
func TestQuatRandom(t *testing.T) {
	actual := QuatRandom(QuatCreate())
	expect := QuatNormalize(QuatCreate(), actual)
	if !testSlice(actual, expect) {
		t.Errorf("random: %v %v", actual, expect)
	}
}
*/

func TestQuatInvert(t *testing.T) {
	actual := QuatInvert(QuatCreate(), quatA)
	expect := []float64{-0.033333, -0.066666, -0.1, 0.133333}
	if !testSlice(actual, expect) {
		t.Errorf("invert: %v", actual)
	}
}

func TestQuatConjugate(t *testing.T) {
	actual := QuatConjugate(QuatCreate(), quatA)
	expect := []float64{-1, -2, -3, 4}
	if !testSlice(actual, expect) {
		t.Errorf("conjugate: %v", actual)
	}
}

func TestQuatExactEquals(t *testing.T) {
	q1 := []float64{0, 0, 0, 1}
	q2 := []float64{0, 0, 0, 1}
	q3 := []float64{0, 0, 0, 1 + 1e-10}
	if !QuatExactEquals(q1, q2) {
		t.Errorf("exact equals: %v %v", q1, q2)
	}
	if QuatExactEquals(q1, q3) {
		t.Errorf("exact equals: %v %v", q1, q3)
	}
}

func TestQuatEquals(t *testing.T) {
	q1 := []float64{0, 0, 0, 1}
	q2 := []float64{0, 0, 0, 1}
	q3 := []float64{0, 0, 0, 1 + 1e-10}
	if !QuatEquals(q1, q2) {
		t.Errorf("exact equals: %v %v", q1, q2)
	}
	if !QuatEquals(q1, q3) {
		t.Errorf("exact equals: %v %v", q1, q3)
	}
}
