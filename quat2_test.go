package glmatrix

import (
	"testing"
)

func equalsQuat2(q1, q2 []float64) bool {
	allSignsFlipped := false
	if len(q1) != len(q2) {
		return false
	}
	for i := 0; i < len(q1); i++ {
		if allSignsFlipped {
			if !equals(q1[i], -q2[i]) {
				return false
			}
		} else {
			if !equals(q1[i], q2[i]) {
				allSignsFlipped = true
				i = 0
			}
		}
	}
	return true
}

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

func TestQuat2RotateByQuatPrepend(t *testing.T) {
	rotationQuat := Quat2Create()
	rotationQuat[0] = 2
	rotationQuat[1] = 5
	rotationQuat[2] = 2
	rotationQuat[3] = -10
	expect := Quat2Multiply(Quat2Create(), rotationQuat, quat2A)
	actual := Quat2RotateByQuatPrepend(Quat2Create(), Quat2GetReal(Quat2Create(), rotationQuat), quat2A)
	if !equalsQuat2(actual, expect) {
		t.Errorf("rotate by quat prepend: %v", actual)
	}
}

func TestQuat2RotateByQuatAppend(t *testing.T) {
	actual := Quat2RotateByQuatAppend(Quat2Create(), quat2A, []float64{2, 5, 2, -10})
	rotationQuat := Quat2Create()
	rotationQuat[0] = 2
	rotationQuat[1] = 5
	rotationQuat[2] = 2
	rotationQuat[3] = -10
	expect := Quat2Multiply(Quat2Create(), quat2A, rotationQuat)
	if !equalsQuat2(actual, expect) {
		t.Errorf("rotate by quat append: %v", actual)
	}
}

func TestQuat2RotateAroundAxis(t *testing.T) {
	ax := []float64{1, 4, 2}
	quat2A := Quat2FromRotationTranslation(Quat2Create(), []float64{1, 2, 3, 4}, []float64{-5, 4, 10})
	Quat2Normalize(quat2A, quat2A)
	matrixA := Mat4FromQuat2(Mat4Create(), quat2A)
	actual := Quat2RotateAroundAxis(Quat2Create(), quat2A, ax, 5)
	matOut := Mat4Rotate(Mat4Create(), matrixA, 5, ax)
	var quat2B = []float64{
		5, 6, 7, 8,
		9, 8, 6, -4,
	}
	expect := Quat2FromMat4(quat2B, matOut)
	if !equalsQuat2(actual, expect) {
		t.Errorf("rotate around axis: \n%v \n%v", actual, expect)
	}
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

func TestQuat2GetDual(t *testing.T) {
	actual := Quat2GetDual(Quat2Create(), quat2A)
	expect := []float64{2, 5, 6, -2, 0, 0, 0, 0}
	if !testSlice(actual, expect) {
		t.Errorf("get dual: %v", actual)
	}
}

func TestQuat2SetReal(t *testing.T) {
	quat2A := []float64{
		1, 2, 3, 4,
		2, 5, 6, -2,
	}
	actual := Quat2SetReal(quat2A, []float64{4, 6, 8, -100})
	expect := []float64{4, 6, 8, -100, 2, 5, 6, -2}
	if !testSlice(actual, expect) {
		t.Errorf("set real: %v", actual)
	}
}

func TestQuat2SetDual(t *testing.T) {
	quat2A := []float64{
		1, 2, 3, 4,
		2, 5, 6, -2,
	}
	actual := Quat2SetDual(quat2A, []float64{4.3, 6, 8, -100})
	expect := []float64{1, 2, 3, 4, 4.3, 6, 8, -100}
	if !testSlice(actual, expect) {
		t.Errorf("set dual: %v", actual)
	}
}

func TestQuat2Conjugate(t *testing.T) {
	actual := Quat2Conjugate(Quat2Create(), quat2A)
	expect := []float64{-1, -2, -3, 4, -2, -5, -6, -2}
	if !testSlice(actual, expect) {
		t.Errorf("conjugate: %v", actual)
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

func TestQuat2Lerp(t *testing.T) {
	actual := Quat2Lerp(Quat2Create(), quat2A, quat2B, 0.7)
	expect := []float64{3.8, 4.8, 5.8, 6.8, 6.9, 7.1, 6.0, -3.4}
	if !testSlice(actual, expect) {
		t.Errorf("lerp: %v", actual)
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

func TestQuat2Length(t *testing.T) {
	actual := Quat2Length(quat2A)
	expect := 5.477225
	if !equals(actual, expect) {
		t.Errorf("length: %v", actual)
	}

	actual = Quat2Len(quat2A)
	if !equals(actual, expect) {
		t.Errorf("len: %v", actual)
	}
}

func TestQuat2RotateX(t *testing.T) {
	quat2A := Quat2Normalize(Quat2Create(), quat2A)
	matrixA := Mat4FromQuat2(Mat4Create(), quat2A)

	actual := Quat2RotateX(Quat2Create(), quat2A, 5)
	matOut := Mat4RotateX(Mat4Create(), matrixA, 5)
	expect := Quat2FromMat4(Quat2Create(), matOut)
	if !equalsQuat2(actual, expect) {
		t.Errorf("rotate x: \n%v \n%v", actual, expect)
	}
}

func TestQuat2RotateY(t *testing.T) {
	quat2A := Quat2Normalize(Quat2Create(), quat2A)
	matrixA := Mat4FromQuat2(Mat4Create(), quat2A)

	actual := Quat2RotateY(Quat2Create(), quat2A, -2)
	matOut := Mat4RotateY(Mat4Create(), matrixA, -2)
	expect := Quat2FromMat4(Quat2Create(), matOut)
	if !equalsQuat2(actual, expect) {
		t.Errorf("rotate y: \n%v \n%v", actual, expect)
	}
}

func TestQuat2RotateZ(t *testing.T) {
	quat2A := Quat2Normalize(Quat2Create(), quat2A)
	matrixA := Mat4FromQuat2(Mat4Create(), quat2A)

	actual := Quat2RotateZ(Quat2Create(), quat2A, 1)
	matOut := Mat4RotateZ(Mat4Create(), matrixA, 1)
	expect := Quat2FromMat4(Quat2Create(), matOut)
	if !equalsQuat2(actual, expect) {
		t.Errorf("rotate z: \n%v \n%v", actual, expect)
	}
}

func TestQuat2FromRotation(t *testing.T) {
	actual := Quat2FromRotation(Quat2Create(), []float64{1, 2, 3, 4})
	expect := []float64{1, 2, 3, 4, 0, 0, 0, 0}
	if !testSlice(actual, expect) {
		t.Errorf("from rotation: %v", actual)
	}
}

func TestQuat2FromRotationTranslationValues(t *testing.T) {
	actual := Quat2FromRotationTranslationValues(1, 2, 3, 4, 1, 2, 3)
	expect := []float64{1, 2, 3, 4, 2, 4, 6, -7}
	if !testSlice(actual, expect) {
		t.Errorf("from rotation toranslation values: %v", actual)
	}
}

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
