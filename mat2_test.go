package glmatrix

import (
	"math"
	"testing"
)

var mat2A = []float64{
	1, 2,
	3, 4,
}

var mat2B = []float64{
	5, 6,
	7, 8,
}

var out2 = []float64{
	0, 0,
	0, 0,
}

var identity2 = []float64{
	1, 0,
	0, 1,
}

func TestMat2Create(t *testing.T) {
	actual := Mat2Create()
	if !testSlice(actual, identity2) {
		t.Errorf("create: %v", actual)
	}
}

func TestMat2Clone(t *testing.T) {
	actual := Mat2Clone(mat2A)
	expect := mat2A
	if !testSlice(actual, expect) {
		t.Errorf("clone: %v", actual)
	}
}

func TestMat2Copy(t *testing.T) {
	actual := Mat2Create()
	Mat2Copy(actual, mat2A)
	expect := mat2A
	if !testSlice(actual, expect) {
		t.Errorf("copy: %v", actual)
	}
}

func TestMat2Identity(t *testing.T) {
	actual := Mat2Create()
	Mat2Identity(actual)
	expect := identity2
	if !testSlice(actual, expect) {
		t.Errorf("identity: %v", actual)
	}
}

func TestMat2Transpose(t *testing.T) {
	actual := Mat2Create()
	Mat2Transpose(actual, mat2A)
	expect := []float64{
		1, 3,
		2, 4,
	}
	if !testSlice(actual, expect) {
		t.Errorf("transpose: %v", actual)
	}
}

func TestMat2Invert(t *testing.T) {
	actual := Mat2Create()
	Mat2Invert(actual, mat2A)
	expect := []float64{
		-2, 1,
		1.5, -0.5,
	}
	if !testSlice(actual, expect) {
		t.Errorf("invert: %v", actual)
	}
}

func TestMat2Adjoint(t *testing.T) {
	actual := Mat2Adjoint(Mat2Create(), mat2A)
	expect := []float64{
		4, -2,
		-3, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("adjoint: %v", actual)
	}
}

func TestMat2Determinant(t *testing.T) {
	actual := Mat2Determinant(mat2A)
	expect := -2.
	if actual != expect {
		t.Errorf("determinant: %v", actual)
	}
}

func TestMat2Multiply(t *testing.T) {
	actual := Mat2Create()
	Mat2Multiply(actual, mat2A, mat2B)
	expect := []float64{
		23, 34,
		31, 46,
	}
	if !testSlice(actual, expect) {
		t.Errorf("multiply: %v", actual)
	}
}

func TestMat2Scale(t *testing.T) {
	actual := Mat2Create()
	Mat2Scale(actual, mat2A, []float64{2, 3})
	expect := []float64{
		2, 4,
		9, 12,
	}
	if !testSlice(actual, expect) {
		t.Errorf("scale: %v", actual)
	}
}

func TestMat2Rotate(t *testing.T) {
	rad := math.Pi * 0.5
	actual := Mat2Create()
	Mat2Rotate(actual, mat2A, rad)
	expect := []float64{
		3, 4,
		-1, -2,
	}
	if !testSlice(actual, expect) {
		t.Errorf("rotate: %v", actual)
	}
}

func TestMat2Str(t *testing.T) {
	actual := Mat2Str(mat2A)
	expect := "mat2(1, 2, 3, 4)"
	if actual != expect {
		t.Errorf("str: %v", actual)
	}
}

func TestMat2Frob(t *testing.T) {
	actual := Mat2Frob(mat2A)
	expect := math.Sqrt(math.Pow(1, 2) + math.Pow(2, 2) + math.Pow(3, 2) + math.Pow(4, 2))
	if actual != expect {
		t.Errorf("frob: %v", actual)
	}
}

func TestMat2LDU(t *testing.T) {
	L := Mat2Create()
	D := Mat2Create()
	U := Mat2Create()
	result := Mat2LDU(L, D, U, []float64{4, 3, 6, 3})
	resultL := Mat2Create()
	resultL[2] = 1.5
	resultD := Mat2Create()
	resultU := Mat2Create()
	resultU[0] = 4.
	resultU[1] = 3.
	resultU[3] = -1.5
	if !testSlice(result[0], resultL) {
		t.Errorf("ldu: %v", result[0])
	}
	if !testSlice(result[1], resultD) {
		t.Errorf("ldu: %v", result[1])
	}
	if !testSlice(result[2], resultU) {
		t.Errorf("ldu: %v", result[2])
	}
}

func TestMat2Add(t *testing.T) {
	actual := Mat2Add(Mat2Create(), mat2A, mat2B)
	expect := []float64{
		6, 8,
		10, 12,
	}
	if !testSlice(actual, expect) {
		t.Errorf("add: %v", actual)
	}
}

func TestMat2Subtract(t *testing.T) {
	actual := Mat2Subtract(Mat2Create(), mat2A, mat2B)
	expect := []float64{
		-4, -4,
		-4, -4,
	}
	if !testSlice(actual, expect) {
		t.Errorf("subtract: %v", actual)
	}
}

func TestMat2FromValues(t *testing.T) {
	actual := Mat2FromValues(1, 2, 3, 4)
	expect := []float64{
		1, 2, 3, 4,
	}
	if !testSlice(actual, expect) {
		t.Errorf("from values: %v", actual)
	}
}

func TestMat2Set(t *testing.T) {
	actual := Mat2Create()
	Mat2Set(actual, 1, 2, 3, 4)
	expect := []float64{
		1, 2, 3, 4,
	}
	if !testSlice(actual, expect) {
		t.Errorf("set: %v", actual)
	}
}

func TestMat2MultiplyScalar(t *testing.T) {
	actual := Mat2MultiplyScalar(Mat2Create(), mat2A, 2)
	expect := []float64{
		2, 4, 6, 8,
	}
	if !testSlice(actual, expect) {
		t.Errorf("multiply scalar: %v", actual)
	}
}

func TestMat2MultiplyScalarAndAdd(t *testing.T) {
	actual := Mat2MultiplyScalarAndAdd(Mat2Create(), mat2A, mat2B, 0.5)
	expect := []float64{
		3.5, 5, 6.5, 8,
	}
	if !testSlice(actual, expect) {
		t.Errorf("multiply scalar and add: %v", actual)
	}
}

func TestMat2ExactEquals(t *testing.T) {
	mat2A := []float64{
		1, 1,
		1, 1,
	}
	mat2B := []float64{
		1, 1,
		1, 1,
	}
	mat2C := []float64{
		1, 1,
		1, 1 + 1e-10,
	}
	if !Mat2ExactEquals(mat2A, mat2B) {
		t.Errorf("exact equal")
	}
	if Mat2ExactEquals(mat2A, mat2C) {
		t.Errorf("exact equal")
	}
}

func TestMat2Equals(t *testing.T) {
	mat2A := []float64{
		1, 1,
		1, 1,
	}
	mat2B := []float64{
		1, 1,
		1, 1,
	}
	mat2C := []float64{
		1, 1,
		1, 1 + 1e-10,
	}
	if !Mat2Equals(mat2A, mat2B) {
		t.Errorf("equal")
	}
	if !Mat2Equals(mat2A, mat2C) {
		t.Errorf("equal")
	}
}
