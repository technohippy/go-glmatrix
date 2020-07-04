package glmatrix

import (
	"math"
	"testing"
)

/*
var quatA := []float64{1,2,3,4}
var quatB := []float64{5,6,7,8}
*/
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
	actual := QuatRotationTo(QuatCreate(), []float64{0, 1, 0}, []float64{0, 1, 0})
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
