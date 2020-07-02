package gomatrix

import (
	"math"
	"testing"
)

var matA = []float64{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
	1, 2, 3, 1,
}

var matB = []float64{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
	4, 5, 6, 1,
}

var out = []float64{
	0, 0, 0, 0,
	0, 0, 0, 0,
	0, 0, 0, 0,
	0, 0, 0, 0,
}

var identity = []float64{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
	0, 0, 0, 1,
}

func testSlice(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, ai := range a {
		//if ai != b[i] {
		if !equals(ai, b[i]) {
			return false
		}
	}
	return true
}

func TestCreate(t *testing.T) {
	actual := Mat4Create()
	if !testSlice(actual, identity) {
		t.Errorf("create: %v", actual)
	}
}

func TestTranspose(t *testing.T) {
	actual := Mat4Create()
	Mat4Transpose(actual, matA)
	expect := []float64{
		1, 0, 0, 1,
		0, 1, 0, 2,
		0, 0, 1, 3,
		0, 0, 0, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("transpose: %v", actual)
	}
}

func TestInvert(t *testing.T) {
	actual := Mat4Create()
	Mat4Invert(actual, matA)
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

func TestMultiply(t *testing.T) {
	actual := Mat4Create()
	Mat4Multiply(actual, matA, matB)
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

func TestTranslate(t *testing.T) {
	actual := Mat4Create()
	Mat4Translate(actual, matA, []float64{4, 5, 6})
	expect := []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		5, 7, 9, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("translate: %v", actual)
	}
}

func TestScale(t *testing.T) {
	actual := Mat4Create()
	Mat4Scale(actual, matA, []float64{4, 5, 6})
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

func TestRotate(t *testing.T) {
	rad := math.Pi * 0.5
	axis := []float64{1, 0, 0}
	actual := Mat4Create()
	Mat4Rotate(actual, matA, rad, axis)
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

func TestPerspective(t *testing.T) {
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

func TestOrtho(t *testing.T) {
	actual := Mat4Create()
	Mat4Ortho(actual, -1, 1, -1, 1, -1, 1)
	expect := []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, -1, 0,
		0, 0, 0, 1,
	}
	if !testSlice(actual, expect) {
		t.Errorf("should place values into out: %v", actual)
	}
}

func TestLookAt(t *testing.T) {
	eye := []float64{0, 0, 0}
	//center := []float64{0, 0, -1}
	view := []float64{0, -1, 0}
	up := []float64{0, 0, -1}
	//right := []float64{1, 0, 0}
	out := Mat4Create()
	Mat4LookAt(out, eye, view, up)
	actual := Vec3TransformMat4([]float64{0, 0, 0}, view, out)
	if !testSlice(actual, []float64{0, 0, -1}) {
		t.Errorf("looking down: %v", actual)
	}

	Mat4LookAt(out, []float64{0, 2, 0}, []float64{0, 0.6, 0}, []float64{0, 0, -1})
	actual = Vec3TransformMat4([]float64{0, 0, 0}, []float64{0, 2, -1}, out)
	if !testSlice(actual, []float64{0, 1, 0}) {
		t.Errorf("#74: %v", actual)
	}
}
