package glmatrix

import (
	"math"
	"testing"
)

var mat2dA = []float64{
	1, 2,
	3, 4,
	5, 6,
}

var mat2dB = []float64{
	7, 8,
	9, 10,
	11, 12,
}

var out2d = []float64{
	0, 0,
	0, 0,
	0, 0,
}

var identity2d = []float64{
	1, 0,
	0, 1,
	0, 0,
}

func TestMat2dCreate(t *testing.T) {
	actual := Mat2dCreate()
	if !testSlice(actual, identity2d) {
		t.Errorf("create: %v", actual)
	}
}

func TestMat2dClone(t *testing.T) {
	actual := Mat2dClone(mat2dA)
	expect := mat2dA
	if !testSlice(actual, expect) {
		t.Errorf("clone: %v", actual)
	}
}

func TestMat2dCopy(t *testing.T) {
	actual := Mat2dCreate()
	Mat2dCopy(actual, mat2dA)
	expect := mat2dA
	if !testSlice(actual, expect) {
		t.Errorf("copy: %v", actual)
	}
}

func TestMat2dIdentity(t *testing.T) {
	actual := Mat2dCreate()
	Mat2dIdentity(actual)
	expect := identity2d
	if !testSlice(actual, expect) {
		t.Errorf("identity: %v", actual)
	}
}

func TestMat2dInvert(t *testing.T) {
	actual := Mat2dCreate()
	Mat2dInvert(actual, mat2dA)
	expect := []float64{
		-2, 1,
		1.5, -0.5,
		1, -2,
	}
	if !testSlice(actual, expect) {
		t.Errorf("invert: %v", actual)
	}
}

func TestMat2dDeterminant(t *testing.T) {
	actual := Mat2dDeterminant(mat2dA)
	expect := -2.
	if actual != expect {
		t.Errorf("determinant: %v", actual)
	}
}

func TestMat2dMultiply(t *testing.T) {
	actual := Mat2dCreate()
	Mat2dMultiply(actual, mat2dA, mat2dB)
	expect := []float64{
		31, 46,
		39, 58,
		52, 76,
	}
	if !testSlice(actual, expect) {
		t.Errorf("multiply: %v", actual)
	}
}

func TestMat2dTranslate(t *testing.T) {
	actual := Mat2dCreate()
	Mat2dTranslate(actual, mat2dA, []float64{2, 3})
	expect := []float64{
		1, 2,
		3, 4,
		16, 22,
	}
	if !testSlice(actual, expect) {
		t.Errorf("translate: %v", actual)
	}
}

func TestMat2dScale(t *testing.T) {
	actual := Mat2dCreate()
	Mat2dScale(actual, mat2dA, []float64{2, 3})
	expect := []float64{
		2, 4,
		9, 12,
		5, 6,
	}
	if !testSlice(actual, expect) {
		t.Errorf("scale: %v", actual)
	}
}

func TestMat2dRotate(t *testing.T) {
	rad := math.Pi * 0.5
	actual := Mat2dCreate()
	Mat2dRotate(actual, mat2dA, rad)
	expect := []float64{
		3, 4,
		-1, -2,
		5, 6,
	}
	if !testSlice(actual, expect) {
		t.Errorf("rotate: %v", actual)
	}
}

func TestMat2dStr(t *testing.T) {
	actual := Mat2dStr(mat2dA)
	expect := "mat2d(1, 2, 3, 4, 5, 6)"
	if actual != expect {
		t.Errorf("str: %v", actual)
	}
}

func TestMat2dFrob(t *testing.T) {
	actual := Mat2dFrob(mat2dA)
	expect := math.Sqrt(math.Pow(1, 2) +
		math.Pow(2, 2) +
		math.Pow(3, 2) +
		math.Pow(4, 2) +
		math.Pow(5, 2) +
		math.Pow(6, 2) + 1)
	if actual != expect {
		t.Errorf("frob: %v", actual)
	}
}

var Mat2dOp1 = []float64{
	1, 2, 3, 4,
	5, 6, 7, 8,
	9, 10, 11, 12,
	13, 14, 15, 16,
}

var Mat2dOp2 = []float64{
	17, 18, 19, 20,
	21, 22, 23, 24,
	25, 26, 27, 28,
	29, 30, 31, 32,
}

func TestMat2dAdd(t *testing.T) {
	actual := Mat2dAdd(Mat2dCreate(), mat2dA, mat2dB)
	expect := []float64{
		8, 10,
		12, 14,
		16, 18,
	}
	if !testSlice(actual, expect) {
		t.Errorf("add: %v", actual)
	}
}

func TestMat2dSubtract(t *testing.T) {
	actual := Mat2dSubtract(Mat2dCreate(), mat2dA, mat2dB)
	expect := []float64{
		-6, -6,
		-6, -6,
		-6, -6,
	}
	if !testSlice(actual, expect) {
		t.Errorf("subtract: %v", actual)
	}
}

func TestMat2dFromValues(t *testing.T) {
	actual := Mat2dFromValues(1, 2, 3, 4, 5, 6)
	expect := []float64{
		1, 2,
		3, 4,
		5, 6,
	}
	if !testSlice(actual, expect) {
		t.Errorf("from values: %v", actual)
	}
}

func TestMat2dSet(t *testing.T) {
	actual := Mat2dCreate()
	Mat2dSet(actual, 1, 2, 3, 4, 5, 6)
	expect := []float64{
		1, 2,
		3, 4,
		5, 6,
	}
	if !testSlice(actual, expect) {
		t.Errorf("set: %v", actual)
	}
}

func TestMat2dMultiplyScalar(t *testing.T) {
	actual := Mat2dMultiplyScalar(Mat2dCreate(), Mat2dOp1, 2)
	expect := []float64{
		2, 4,
		6, 8,
		10, 12,
	}
	if !testSlice(actual, expect) {
		t.Errorf("multiply scalar: %v", actual)
	}
}

func TestMat2dMultiplyScalarAndAdd(t *testing.T) {
	actual := Mat2dMultiplyScalarAndAdd(Mat2dCreate(), Mat2dOp1, Mat2dOp2, 0.5)
	expect := []float64{
		9.5, 11,
		12.5, 14,
		15.5, 17,
	}
	if !testSlice(actual, expect) {
		t.Errorf("multiply scalar and add: %v", actual)
	}
}

func TestMat2dExactEquals(t *testing.T) {
	mat2dA := []float64{
		1, 1,
		1, 1,
		1, 1,
	}
	mat2dB := []float64{
		1, 1,
		1, 1,
		1, 1,
	}
	matC := []float64{
		1, 1,
		1, 1,
		1, 1 + 1e-10,
	}
	if !Mat2dExactEquals(mat2dA, mat2dB) {
		t.Errorf("exact equal")
	}
	if Mat2dExactEquals(mat2dA, matC) {
		t.Errorf("exact equal")
	}
}

func TestMat2dEquals(t *testing.T) {
	mat2dA := []float64{
		1, 1,
		1, 1,
		1, 1,
	}
	mat2dB := []float64{
		1, 1,
		1, 1,
		1, 1,
	}
	matC := []float64{
		1, 1,
		1, 1,
		1, 1 + 1e-10,
	}
	if !Mat2dEquals(mat2dA, mat2dB) {
		t.Errorf("equal")
	}
	if !Mat2dEquals(mat2dA, matC) {
		t.Errorf("equal")
	}
}
