package glmatrix

import (
	"math"
	"testing"
)

var mat3A = []float64{
	1, 0, 0,
	0, 1, 0,
	1, 2, 1,
}

var mat3B = []float64{
	1, 0, 0,
	0, 1, 0,
	3, 4, 1,
}

var out3 = []float64{
	0, 0, 0,
	0, 0, 0,
	0, 0, 0,
}

var identity3 = []float64{
	1, 0, 0,
	0, 1, 0,
	0, 0, 1,
}

func TestMat3Create(t *testing.T) {
	actual := Mat3Create()
	if !testSlice(actual, identity3) {
		t.Errorf("create: %v", actual)
	}
}

func TestMat3Clone(t *testing.T) {
	actual := Mat3Clone(mat3A)
	expect := mat3A
	if !testSlice(actual, expect) {
		t.Errorf("clone: %v", actual)
	}
}

func TestMat3Copy(t *testing.T) {
	actual := Mat3Create()
	Mat3Copy(actual, mat3A)
	expect := mat3A
	if !testSlice(actual, expect) {
		t.Errorf("copy: %v", actual)
	}
}

func TestMat3Identity(t *testing.T) {
	actual := Mat3Create()
	Mat3Identity(actual)
	expect := identity3
	if !testSlice(actual, expect) {
		t.Errorf("identity: %v", actual)
	}
}

func TestMat3Transpose(t *testing.T) {
	actual := Mat3Create()
	Mat3Transpose(actual, mat3A)
	expect := []float64{
		1, 0, 1,
		0, 1, 2,
		0, 0, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("transpose: %v", actual)
	}

	actual = []float64{
		1, 0, 1,
		0, 1, 2,
		0, 0, 1,
	}
	Mat3Transpose(actual, actual)
	expect = []float64{
		1, 0, 0,
		0, 1, 0,
		1, 2, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("transpose: %v", actual)
	}
}

func TestMat3Invert(t *testing.T) {
	actual := Mat3Create()
	Mat3Invert(actual, mat3A)
	expect := []float64{
		1, 0, 0,
		0, 1, 0,
		-1, -2, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("invert: %v", actual)
	}
}

func TestMat3Adjoint(t *testing.T) {
	actual := Mat3Adjoint(Mat3Create(), mat3A)
	expect := []float64{
		1, 0, 0,
		0, 1, 0,
		-1, -2, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("adjoint: %v", actual)
	}
}

func TestMat3Determinant(t *testing.T) {
	actual := Mat3Determinant(mat3A)
	expect := 1.
	if actual != expect {
		t.Errorf("determinant: %v", actual)
	}
}

func TestMat3Multiply(t *testing.T) {
	actual := Mat3Create()
	Mat3Multiply(actual, mat3A, mat3B)
	expect := []float64{
		1, 0, 0,
		0, 1, 0,
		4, 6, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("multiply: %v", actual)
	}
}

func TestMat3Translate(t *testing.T) {
	actual := Mat3Create()
	Mat3Translate(actual, mat3A, []float64{4, 5, 6})
	expect := []float64{
		1, 0, 0,
		0, 1, 0,
		5, 7, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("translate: %v", actual)
	}
}

func TestMat3Scale(t *testing.T) {
	actual := Mat3Create()
	Mat3Scale(actual, mat3A, []float64{2, 2})
	expect := []float64{
		2, 0, 0,
		0, 2, 0,
		1, 2, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("scale: %v", actual)
	}
}

func TestMat3Rotate(t *testing.T) {
	rad := math.Pi * 0.5
	actual := Mat3Create()
	Mat3Rotate(actual, mat3A, rad)
	expect := []float64{
		0, 1, 0,
		-1, 0, 0,
		1, 2, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("rotate: %v", actual)
	}
}

func TestMat3FromTranslation(t *testing.T) {
	actual := Mat3FromTranslation(Mat3Create(), []float64{2, 3})
	expect := []float64{
		1, 0, 0,
		0, 1, 0,
		2, 3, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("from translation: %v", actual)
	}
}

func TestMat3FromRotation(t *testing.T) {
	actual := Mat3FromRotation(Mat3Create(), math.Pi/4)
	s := math.Sin(math.Pi / 4)
	c := math.Cos(math.Pi / 4)
	expect := []float64{
		c, s, 0,
		-s, c, 0,
		0, 0, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("from rotation: %v", actual)
	}
}

func TestMat3FromScaling(t *testing.T) {
	actual := Mat3FromScaling(Mat3Create(), []float64{2, 3})
	expect := []float64{
		2, 0, 0,
		0, 3, 0,
		0, 0, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("from scaling: %v", actual)
	}
}

func TestMat3FromMat2d(t *testing.T) {
	actual := Mat3FromMat2d(Mat3Create(), []float64{1, 2, 3, 4, 5, 6})
	expect := []float64{
		1, 2, 0,
		3, 4, 0,
		5, 6, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("from mat2d: %v", actual)
	}
}

func TestMat3Str(t *testing.T) {
	actual := Mat3Str(mat3A)
	expect := "mat3(1, 0, 0, 0, 1, 0, 1, 2, 1)"
	if actual != expect {
		t.Errorf("str: %v", actual)
	}
}

func TestMat3Frob(t *testing.T) {
	actual := Mat3Frob(mat3A)
	expect := math.Sqrt(
		math.Pow(1, 2) + math.Pow(0, 2) + math.Pow(0, 2) +
			math.Pow(0, 2) + math.Pow(1, 2) + math.Pow(0, 2) +
			math.Pow(1, 2) + math.Pow(2, 2) + math.Pow(1, 2),
	)

	if actual != expect {
		t.Errorf("frob: %v %v", actual, expect)
	}
}

func TestMat3Add(t *testing.T) {
	mat3A := []float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	mat3B := []float64{
		10, 11, 12,
		13, 14, 15,
		16, 17, 18,
	}
	actual := Mat3Add(Mat3Create(), mat3A, mat3B)
	expect := []float64{
		11, 13, 15,
		17, 19, 21,
		23, 25, 27,
	}
	if !testSlice(actual, expect) {
		t.Errorf("add: %v", actual)
	}
}

func TestMat3Subtract(t *testing.T) {
	mat3A := []float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	mat3B := []float64{
		10, 11, 12,
		13, 14, 15,
		16, 17, 18,
	}
	actual := Mat3Subtract(Mat3Create(), mat3A, mat3B)
	expect := []float64{
		-9, -9, -9,
		-9, -9, -9,
		-9, -9, -9,
	}
	if !testSlice(actual, expect) {
		t.Errorf("subtract: %v", actual)
	}
}

func TestMat3FromValues(t *testing.T) {
	actual := Mat3FromValues(1, 2, 3, 4, 5, 6, 7, 8, 9)
	expect := []float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	if !testSlice(actual, expect) {
		t.Errorf("from values: %v", actual)
	}
}

func TestMat3Set(t *testing.T) {
	actual := Mat3Create()
	Mat3Set(actual, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	expect := []float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	if !testSlice(actual, expect) {
		t.Errorf("set: %v", actual)
	}
}

func TestMat3MultiplyScalar(t *testing.T) {
	mat3A := []float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	actual := Mat3MultiplyScalar(Mat3Create(), mat3A, 2)
	expect := []float64{
		2, 4, 6,
		8, 10, 12,
		14, 16, 18,
	}
	if !testSlice(actual, expect) {
		t.Errorf("multiply scalar: %v", actual)
	}
}

func TestMat3MultiplyScalarAndAdd(t *testing.T) {
	mat3A := []float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	mat3B := []float64{
		10, 11, 12,
		13, 14, 15,
		16, 17, 18,
	}
	actual := Mat3MultiplyScalarAndAdd(Mat3Create(), mat3A, mat3B, 0.5)
	expect := []float64{
		6, 7.5, 9,
		10.5, 12, 13.5,
		15, 16.5, 18,
	}
	if !testSlice(actual, expect) {
		t.Errorf("multiply scalar and add: %v", actual)
	}
}

func TestMat3ExactEquals(t *testing.T) {
	mat3A := []float64{
		1, 1, 1,
		1, 1, 1,
		1, 1, 1,
	}
	mat3B := []float64{
		1, 1, 1,
		1, 1, 1,
		1, 1, 1,
	}
	matC := []float64{
		1, 1, 1,
		1, 1, 1,
		1, 1, 1 + 1e-10,
	}
	if !Mat3ExactEquals(mat3A, mat3B) {
		t.Errorf("exact equal")
	}
	if Mat3ExactEquals(mat3A, matC) {
		t.Errorf("exact equal")
	}
}

func TestMat3Equals(t *testing.T) {
	mat3A := []float64{
		1, 1, 1,
		1, 1, 1,
		1, 1, 1,
	}
	mat3B := []float64{
		1, 1, 1,
		1, 1, 1,
		1, 1, 1,
	}
	matC := []float64{
		1, 1, 1,
		1, 1, 1,
		1, 1, 1 + 1e-10,
	}
	if !Mat3Equals(mat3A, mat3B) {
		t.Errorf("equal")
	}
	if !Mat3Equals(mat3A, matC) {
		t.Errorf("equal")
	}
}
