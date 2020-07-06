package glmatrix

import (
	"testing"
)

var quat2A = []float64{
	1, 2, 3, 4,
	2, 5, 6, -2,
}

var quat2B = []float64{
	5, 6, 7, 8,
	9, 8, 6, -4,
}

var identityq2 = []float64{
	0, 0, 0, 1,
	0, 0, 0, 0,
}

func TestQuat2Create(t *testing.T) {
	actual := Quat2Create()
	if !testSlice(actual, identityq2) {
		t.Errorf("create: %v", actual)
	}
}

func TestQuat2Clone(t *testing.T) {
	actual := Quat2Clone(quat2A)
	expect := quat2A
	if !testSlice(actual, expect) {
		t.Errorf("clone: %v", actual)
	}
}

func TestQuat2Copy(t *testing.T) {
	actual := Quat2Create()
	Quat2Copy(actual, quat2A)
	expect := quat2A
	if !testSlice(actual, expect) {
		t.Errorf("copy: %v", actual)
	}
}

func TestQuat2Identity(t *testing.T) {
	actual := Quat2Create()
	Quat2Identity(actual)
	expect := identityq2
	if !testSlice(actual, expect) {
		t.Errorf("identity: %v", actual)
	}
}

func TestQuat2Invert(t *testing.T) {
	actual := Quat2Create()
	Quat2Invert(actual, quat2A)
	expect := []float64{
		-0.0333333333, -0.06666666666, -0.1, 0.13333333333,
		-2. / 30, -5. / 30, -6. / 30, -2. / 30,
	}
	if !testSlice(actual, expect) {
		t.Errorf("invert: %v", actual)
	}
}

func TestQuat2Multiply(t *testing.T) {
	actual := Quat2Create()
	Quat2Multiply(actual, quat2A, quat2B)
	expect := []float64{
		24, 48, 48, -6,
		25, 89, 23, -157,
	}
	if !testSlice(actual, expect) {
		t.Errorf("multiply: %v", actual)
	}
}

func TestQuat2Translate(t *testing.T) {
	quat2A := Quat2Normalize(Quat2Create(), quat2A)
	matrixA := Mat4FromQuat2(Mat4Create(), quat2A)

	actual := Quat2Translate(Quat2Create(), quat2A, vec)
	matOut := Mat4Translate(Mat4Create(), matrixA, vec)
	expect := Quat2FromMat4(Quat2Create(), matOut)
	if !testSlice(actual, expect) {
		t.Errorf("translate: %v", actual)
	}
}

func TestQuat2Scale(t *testing.T) {
	actual := Quat2Create()
	Quat2Scale(actual, quat2A, 2.)
	expect := []float64{
		2, 4, 6, 8,
		4, 10, 12, -4,
	}
	if !testSlice(actual, expect) {
		t.Errorf("scale: %v", actual)
	}
}

/*
func TestQuat2RotateX(t *testing.T) {
	quat2A := Quat2Normalize(Quat2Create(), quat2A)
	matrixA := Mat4FromQuat2(Mat4Create(), quat2A)

	actual := Quat2RotateX(Quat2Create(), quat2A, 5)
	matOut := Mat4RotateX(Mat4Create(), matrixA, 5)
	expect := Quat2FromMat4(Quat2Create(), matOut)
	if !testSlice(actual, expect) {
		t.Errorf("rotate x: \n%v \n%v", actual, expect)
	}
}

func TestQuat2RotateY(t *testing.T) {
	rad := math.Pi * 0.5
	actual := Quat2RotateY(Quat2Create(), quat2A, rad)
	expect := []float64{
		math.Cos(rad), 0, -math.Sin(rad), 0,
		0, 1, 0, 0,
		math.Sin(rad), 0, math.Cos(rad), 0,
		1, 2, 3, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("rotate y: %v", actual)
	}
}

func TestQuat2RotateZ(t *testing.T) {
	rad := math.Pi * 0.5
	actual := Quat2RotateZ(Quat2Create(), quat2A, rad)
	expect := []float64{
		math.Cos(rad), math.Sin(rad), 0, 0,
		-math.Sin(rad), math.Cos(rad), 0, 0,
		0, 0, 1, 0,
		1, 2, 3, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("rotate z: %v", actual)
	}
}
*/

func TestQuat2GetTranslation(t *testing.T) {
	quat2A := Quat2FromTranslation(Quat2Create(), []float64{1, 2, 3})
	actual := Quat2GetTranslation(Vec3Create(), quat2A)
	expect := []float64{1, 2, 3}
	if !testSlice(actual, expect) {
		t.Errorf("get translation: %v", actual)
	}
}

func TestQuat2Str(t *testing.T) {
	actual := Quat2Str(quat2A)
	expect := "quat2(1, 2, 3, 4, 2, 5, 6, -2)"
	if actual != expect {
		t.Errorf("str: %v", actual)
	}
}

func TestQuat2Add(t *testing.T) {
	actual := Quat2Add(Quat2Create(), quat2A, quat2B)
	expect := []float64{
		6, 8, 10, 12,
		11, 13, 12, -6,
	}
	if !testSlice(actual, expect) {
		t.Errorf("add: %v", actual)
	}
}

func TestQuat2FromValues(t *testing.T) {
	actual := Quat2FromValues(1, 2, 3, 4, 5, 7, 8, -2)
	expect := []float64{
		1, 2, 3, 4,
		5, 7, 8, -2,
	}
	if !testSlice(actual, expect) {
		t.Errorf("from values: %v", actual)
	}
}

func TestQuat2Set(t *testing.T) {
	actual := Quat2Create()
	Quat2Set(actual, 1, 2, 3, 4, 2, 5, 6, -2)
	expect := []float64{
		1, 2, 3, 4,
		2, 5, 6, -2,
	}
	if !testSlice(actual, expect) {
		t.Errorf("set: %v", actual)
	}
}

func TestQuat2ExactEquals(t *testing.T) {
	quat2A := []float64{
		1, 1, 1, 1,
		1, 1, 1, 1,
	}
	quat2B := []float64{
		1, 1, 1, 1,
		1, 1, 1, 1,
	}
	matC := []float64{
		1, 1, 1, 1,
		1, 1, 1, 1 + 1e-10,
	}
	if !Quat2ExactEquals(quat2A, quat2B) {
		t.Errorf("exact equal")
	}
	if Quat2ExactEquals(quat2A, matC) {
		t.Errorf("exact equal")
	}
}

func TestQuat2Equals(t *testing.T) {
	quat2A := []float64{
		1, 1, 1, 1,
		1, 1, 1, 1,
	}
	quat2B := []float64{
		1, 1, 1, 1,
		1, 1, 1, 1,
	}
	matC := []float64{
		1, 1, 1, 1,
		1, 1, 1, 1 + 1e-10,
	}
	if !Quat2Equals(quat2A, quat2B) {
		t.Errorf("equal")
	}
	if !Quat2Equals(quat2A, matC) {
		t.Errorf("equal")
	}
}
