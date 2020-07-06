package glmatrix

import (
	"math"
	"testing"
)

var mat4A = []float64{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
	1, 2, 3, 1,
}

var mat4B = []float64{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
	4, 5, 6, 1,
}

var out4 = []float64{
	0, 0, 0, 0,
	0, 0, 0, 0,
	0, 0, 0, 0,
	0, 0, 0, 0,
}

var identity4 = []float64{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
	0, 0, 0, 1,
}

func TestMat4Create(t *testing.T) {
	actual := Mat4Create()
	if !testSlice(actual, identity4) {
		t.Errorf("create: %v", actual)
	}
}

func TestMat4Clone(t *testing.T) {
	actual := Mat4Clone(mat4A)
	expect := mat4A
	if !testSlice(actual, expect) {
		t.Errorf("clone: %v", actual)
	}
}

func TestMat4Copy(t *testing.T) {
	actual := Mat4Create()
	Mat4Copy(actual, mat4A)
	expect := mat4A
	if !testSlice(actual, expect) {
		t.Errorf("copy: %v", actual)
	}
}

func TestMat4Identity(t *testing.T) {
	actual := Mat4Create()
	Mat4Identity(actual)
	expect := identity4
	if !testSlice(actual, expect) {
		t.Errorf("identity: %v", actual)
	}
}

func TestMat4Transpose(t *testing.T) {
	actual := Mat4Create()
	Mat4Transpose(actual, mat4A)
	expect := []float64{
		1, 0, 0, 1,
		0, 1, 0, 2,
		0, 0, 1, 3,
		0, 0, 0, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("transpose: %v", actual)
	}

	actual = []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		1, 2, 3, 1,
	}
	Mat4Transpose(actual, actual)
	if !testSlice(actual, expect) {
		t.Errorf("transpose: %v", actual)
	}
}

func TestMat4Invert(t *testing.T) {
	actual := Mat4Create()
	Mat4Invert(actual, mat4A)
	expect := []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		-1, -2, -3, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("invert: %v", actual)
	}
}

func TestMat4Adjoint(t *testing.T) {
	actual := Mat4Adjoint(Mat4Create(), mat4A)
	expect := []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		-1, -2, -3, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("adjoint: %v", actual)
	}
}

func TestMat4Determinant(t *testing.T) {
	actual := Mat4Determinant(mat4A)
	expect := 1.
	if actual != expect {
		t.Errorf("determinant: %v", actual)
	}
}

func TestMat4Multiply(t *testing.T) {
	actual := Mat4Create()
	Mat4Multiply(actual, mat4A, mat4B)
	expect := []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		5, 7, 9, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("multiply: %v", actual)
	}
}

func TestMat4Translate(t *testing.T) {
	actual := Mat4Create()
	Mat4Translate(actual, mat4A, []float64{4, 5, 6})
	expect := []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		5, 7, 9, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("translate: %v", actual)
	}

	actual = []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		1, 2, 3, 1,
	}
	Mat4Translate(actual, actual, []float64{4, 5, 6})
	if !testSlice(actual, expect) {
		t.Errorf("translate: %v", actual)
	}
}

func TestMat4Scale(t *testing.T) {
	actual := Mat4Create()
	Mat4Scale(actual, mat4A, []float64{4, 5, 6})
	expect := []float64{
		4, 0, 0, 0,
		0, 5, 0, 0,
		0, 0, 6, 0,
		1, 2, 3, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("scale: %v", actual)
	}
}

func TestMat4Rotate(t *testing.T) {
	rad := math.Pi * 0.5
	axis := []float64{1, 0, 0}
	actual := Mat4Create()
	Mat4Rotate(actual, mat4A, rad, axis)
	expect := []float64{
		1, 0, 0, 0,
		0, math.Cos(rad), math.Sin(rad), 0,
		0, -math.Sin(rad), math.Cos(rad), 0,
		1, 2, 3, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("rotate: %v", actual)
	}
}

func TestMat4RotateX(t *testing.T) {
	rad := math.Pi * 0.5
	actual := Mat4RotateX(Mat4Create(), mat4A, rad)
	expect := []float64{
		1, 0, 0, 0,
		0, math.Cos(rad), math.Sin(rad), 0,
		0, -math.Sin(rad), math.Cos(rad), 0,
		1, 2, 3, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("rotate x: %v", actual)
	}
}

func TestMat4RotateY(t *testing.T) {
	rad := math.Pi * 0.5
	actual := Mat4RotateY(Mat4Create(), mat4A, rad)
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

func TestMat4RotateZ(t *testing.T) {
	rad := math.Pi * 0.5
	actual := Mat4RotateZ(Mat4Create(), mat4A, rad)
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

func TestMat4GetTranslation(t *testing.T) {
	actual := Mat4GetTranslation(Vec3Create(), mat4A)
	expect := []float64{1, 2, 3}
	if !testSlice(actual, expect) {
		t.Errorf("get translation: %v", actual)
	}
}

func TestMat4GetScaling(t *testing.T) {
	actual := Mat4GetScaling(Vec3Create(), mat4B)
	expect := []float64{1, 1, 1}
	if !testSlice(actual, expect) {
		t.Errorf("get scaling: %v", actual)
	}
}

func TestMat4GetRotation(t *testing.T) {
	actual := Mat4GetRotation(QuatCreate(), identity4)
	expect := QuatIdentity(QuatCreate())
	if !testSlice(actual, expect) {
		t.Errorf("get rotation: %v", actual)
	}
}

func TestMat4Frustum(t *testing.T) {
	actual := Mat4Frustum(NewMat4(), -1, 1, -1, 1, -1, 1)
	expect := []float64{
		-1, 0, 0, 0,
		0, -1, 0, 0,
		0, 0, 0, -1,
		0, 0, 1, 0,
	}
	if !testSlice(actual, expect) {
		t.Errorf("flustum: %v", actual)
	}
}

func TestMat4Perspective(t *testing.T) {
	actual := Mat4Create()
	Mat4Perspective(actual, 45*math.Pi/180., 640./480, 0.1, 200.)
	expect := []float64{
		1.81066, 0, 0, 0,
		0, 2.414213, 0, 0,
		0, 0, -1.001, -1,
		0, 0, -0.2001, 0,
	}
	if !testSlice(actual, expect) {
		t.Errorf("with nonzero near, 45deg fovy, and realistic aspect ratio: %v", actual)
	}
}

func TestMat4Ortho(t *testing.T) {
	actual := Mat4Create()
	Mat4Ortho(actual, -1, 1, -1, 1, -1, 1)
	expect := []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, -1, 0,
		0, 0, 0, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("should place values into out4: %v", actual)
	}
}

func TestMat4LookAt(t *testing.T) {
	eye := []float64{0, 0, 0}
	//center := []float64{0, 0, -1}
	view := []float64{0, -1, 0}
	up := []float64{0, 0, -1}
	//right := []float64{1, 0, 0}
	out4 := Mat4Create()
	Mat4LookAt(out4, eye, view, up)
	actual := Vec3TransformMat4([]float64{0, 0, 0}, view, out4)
	expect := []float64{0, 0, -1}
	if !testSlice(actual, expect) {
		t.Errorf("looking down: %v", actual)
	}

	Mat4LookAt(out4, []float64{0, 2, 0}, []float64{0, 0.6, 0}, []float64{0, 0, -1})
	actual = Vec3TransformMat4([]float64{0, 0, 0}, []float64{0, 2, -1}, out4)
	expect = []float64{0, 1, 0}
	if !testSlice(actual, expect) {
		t.Errorf("#74: %v", actual)
	}
}

func TestMat4TargetTo(t *testing.T) {
	view := []float64{0, -1, 0}
	up := []float64{0, 0, -1}
	out4 := Mat4TargetTo(Mat4Create(), []float64{0, 0, 0}, view, up)
	actual := Vec3TransformMat4([]float64{0, 0, 0}, view, out4)
	expect := []float64{0, 0, 1}
	if !testSlice(actual, expect) {
		t.Errorf("target to: %v", actual)
	}
}

func TestMat4FromTranslation(t *testing.T) {
	actual := Mat4FromTranslation(Mat4Create(), []float64{2, 3, 4})
	expect := []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		2, 3, 4, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("from translation: %v", actual)
	}
}

func TestMat4FromScaling(t *testing.T) {
	actual := Mat4FromScaling(Mat4Create(), []float64{2, 3, 4})
	expect := []float64{
		2, 0, 0, 0,
		0, 3, 0, 0,
		0, 0, 4, 0,
		0, 0, 0, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("from scaling: %v", actual)
	}
}

func TestMat4FromRotation(t *testing.T) {
	actual := Mat4FromRotation(Mat4Create(), math.Pi/4, []float64{0, 1, 0})
	s := math.Sin(math.Pi / 4)
	c := math.Cos(math.Pi / 4)
	expect := []float64{
		c, 0, -s, 0,
		0, 1, 0, 0,
		s, 0, c, 0,
		0, 0, 0, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("from rotation: \n%v \n%v", actual, expect)
	}
}

func TestMat4FromXRotation(t *testing.T) {
	actual := Mat4FromXRotation(Mat4Create(), math.Pi/4)
	s := math.Sin(math.Pi / 4)
	c := math.Cos(math.Pi / 4)
	expect := []float64{
		1, 0, 0, 0,
		0, c, s, 0,
		0, -s, c, 0,
		0, 0, 0, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("from rotation: \n%v \n%v", actual, expect)
	}
}

func TestMat4FromYRotation(t *testing.T) {
	actual := Mat4FromYRotation(Mat4Create(), math.Pi/4)
	s := math.Sin(math.Pi / 4)
	c := math.Cos(math.Pi / 4)
	expect := []float64{
		c, 0, -s, 0,
		0, 1, 0, 0,
		s, 0, c, 0,
		0, 0, 0, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("from rotation: \n%v \n%v", actual, expect)
	}
}

func TestMat4FromZRotation(t *testing.T) {
	actual := Mat4FromZRotation(Mat4Create(), math.Pi/4)
	s := math.Sin(math.Pi / 4)
	c := math.Cos(math.Pi / 4)
	expect := []float64{
		c, s, 0, 0,
		-s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("from rotation: \n%v \n%v", actual, expect)
	}
}

func TestMat4FromRotationTranslationScale(t *testing.T) {
	q := QuatFromValues(1, 0, 0, 0)
	v := Vec3FromValues(1, 2, 3)
	s := []float64{4, 5, 6}

	actual := Mat4FromRotationTranslationScale(Mat4Create(), q, v, s)

	transMat := Mat4Create()
	Mat4Identity(transMat)
	Mat4Translate(transMat, transMat, v)
	rotateMat := Mat4Create()
	rotateMat = Mat4FromQuat(rotateMat, q)
	expect := Mat4Multiply(Mat4Create(), transMat, rotateMat)
	Mat4Scale(expect, expect, s)

	if !testSlice(actual, expect) {
		t.Errorf("from rotation, translation and scale: \n%v \n%v", actual, expect)
	}
}

func TestMat4Str(t *testing.T) {
	actual := Mat4Str(mat4A)
	expect := "mat4(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 2, 3, 1)"
	if actual != expect {
		t.Errorf("str: %v", actual)
	}
}

func TestMat4Frob(t *testing.T) {
	actual := Mat4Frob(mat4A)
	expect := math.Sqrt(math.Pow(1, 2) + math.Pow(1, 2) + math.Pow(1, 2) + math.Pow(1, 2) + math.Pow(1, 2) + math.Pow(2, 2) + math.Pow(3, 2))
	if actual != expect {
		t.Errorf("frob: %v", actual)
	}
}

var mat4Op1 = []float64{
	1, 2, 3, 4,
	5, 6, 7, 8,
	9, 10, 11, 12,
	13, 14, 15, 16,
}

var mat4Op2 = []float64{
	17, 18, 19, 20,
	21, 22, 23, 24,
	25, 26, 27, 28,
	29, 30, 31, 32,
}

func TestMat4Add(t *testing.T) {
	actual := Mat4Add(Mat4Create(), mat4Op1, mat4Op2)
	expect := []float64{
		18, 20, 22, 24,
		26, 28, 30, 32,
		34, 36, 38, 40,
		42, 44, 46, 48,
	}
	if !testSlice(actual, expect) {
		t.Errorf("add: %v", actual)
	}
}

func TestMat4Subtract(t *testing.T) {
	actual := Mat4Subtract(Mat4Create(), mat4Op1, mat4Op2)
	expect := []float64{
		-16, -16, -16, -16,
		-16, -16, -16, -16,
		-16, -16, -16, -16,
		-16, -16, -16, -16,
	}
	if !testSlice(actual, expect) {
		t.Errorf("subtract: %v", actual)
	}
}

func TestMat4FromValues(t *testing.T) {
	actual := Mat4FromValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	expect := []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}
	if !testSlice(actual, expect) {
		t.Errorf("from values: %v", actual)
	}
}

func TestMat4Set(t *testing.T) {
	actual := Mat4Create()
	Mat4Set(actual, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	expect := []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}
	if !testSlice(actual, expect) {
		t.Errorf("set: %v", actual)
	}
}

func TestMat4MultiplyScalar(t *testing.T) {
	actual := Mat4MultiplyScalar(Mat4Create(), mat4Op1, 2)
	expect := []float64{
		2, 4, 6, 8,
		10, 12, 14, 16,
		18, 20, 22, 24,
		26, 28, 30, 32,
	}
	if !testSlice(actual, expect) {
		t.Errorf("multiply scalar: %v", actual)
	}
}

func TestMat4MultiplyScalarAndAdd(t *testing.T) {
	actual := Mat4MultiplyScalarAndAdd(Mat4Create(), mat4Op1, mat4Op2, 0.5)
	expect := []float64{
		9.5, 11, 12.5, 14,
		15.5, 17, 18.5, 20,
		21.5, 23, 24.5,
		26, 27.5, 29, 30.5, 32,
	}
	if !testSlice(actual, expect) {
		t.Errorf("multiply scalar and add: %v", actual)
	}
}

func TestMat4ExactEquals(t *testing.T) {
	mat4A := []float64{
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
	}
	mat4B := []float64{
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
	}
	matC := []float64{
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1 + 1e-10,
	}
	if !Mat4ExactEquals(mat4A, mat4B) {
		t.Errorf("exact equal")
	}
	if Mat4ExactEquals(mat4A, matC) {
		t.Errorf("exact equal")
	}
}

func TestMat4Equals(t *testing.T) {
	mat4A := []float64{
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
	}
	mat4B := []float64{
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
	}
	matC := []float64{
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1 + 1e-10,
	}
	if !Mat4Equals(mat4A, mat4B) {
		t.Errorf("equal")
	}
	if !Mat4Equals(mat4A, matC) {
		t.Errorf("equal")
	}
}
