package glmatrix

import (
	"math"
	"testing"
)

var vecA = []float64{1, 2, 3}
var vecB = []float64{4, 5, 6}

func TestVec3RotateX(t *testing.T) {
	actual := Vec3Create()
	vecA := []float64{0, 1, 0}
	vecB := []float64{0, 0, 0}
	Vec3RotateX(actual, vecA, vecB, math.Pi)
	if !testSlice(actual, []float64{0, -1, 0}) {
		t.Errorf("rotateX: %v", actual)
	}
}

func TestVec3TransformMat4(t *testing.T) {
	matr := Mat4LookAt(Mat4Create(), []float64{5, 6, 7}, []float64{2, 6, 7}, []float64{0, 1, 0})
	vecA := []float64{1, 2, 3}
	actual := Vec3Create()
	Vec3TransformMat4(actual, vecA, matr)
	if !testSlice(actual, []float64{4, -4, -4}) {
		t.Errorf("lookAt should rotate and translate the input: %v", actual)
	}
}

func TestVec3TransformMat3(t *testing.T) {
	actual := Vec3Create()
	vecA := []float64{1, 2, 3}
	matr := []float64{1, 0, 0, 0, 1, 0, 0, 0, 1}
	Vec3TransformMat3(actual, vecA, matr)
	if !testSlice(actual, vecA) {
		t.Errorf("transform with an identity: %v", actual)
	}

	actual = Vec3Create()
	vecA = []float64{0, 1, 0}
	matr = []float64{1, 0, 0, 0, 0, 1, 0, -1, 0}
	Vec3TransformMat3(actual, vecA, matr)
	if !testSlice(actual, []float64{0, 0, 1}) {
		t.Errorf("transform with 90deg about X: %v", actual)
	}

	actual = Vec3Create()
	vecA = []float64{1, 0, 0}
	matr = []float64{0, 0, -1, 0, 1, 0, 1, 0, 0}
	Vec3TransformMat3(actual, vecA, matr)
	if !testSlice(actual, []float64{0, 0, -1}) {
		t.Errorf("transform with 90deg about Y: %v", actual)
	}

	actual = Vec3Create()
	vecA = []float64{1, 0, 0}
	matr = Mat4LookAt(Mat4Create(), []float64{5, 6, 7}, []float64{2, 6, 7}, []float64{0, 1, 0})
	n := Mat3Create()
	matr = Mat3Transpose(n, Mat3Invert(n, Mat3FromMat4(n, matr)))
	Vec3TransformMat3(actual, vecA, matr)
	if !testSlice(actual, []float64{0, 0, 1}) {
		t.Errorf("transform with a lookAt normal matrix: %v", actual)
	}
}

func TestVec3TransformQuat(t *testing.T) {
	actual := Vec3TransformQuat(Vec3Create(), vecA, []float64{0.18257418567011074, 0.3651483713402215, 0.5477225570103322, 0.730296742680443})
	expect := []float64{1, 2, 3}
	if !testSlice(actual, expect) {
		t.Errorf("transform quat: %v", actual)
	}
}

func TestVec3Create(t *testing.T) {
	actual := Vec3Create()
	expect := []float64{0, 0, 0}
	if !testSlice(actual, expect) {
		t.Errorf("create: %v", actual)
	}
}

func TestVec3Clone(t *testing.T) {
	actual := Vec3Clone(vecA)
	expect := vecA
	if !testSlice(actual, expect) {
		t.Errorf("clone: %v", actual)
	}
	if &actual == &expect {
		t.Errorf("clone: %v", actual)
	}
}

func TestVec3FromValues(t *testing.T) {
	actual := Vec3FromValues(1, 2, 3)
	expect := []float64{1, 2, 3}
	if !testSlice(actual, expect) {
		t.Errorf("from values: %v", actual)
	}
}

func TestVec3Copy(t *testing.T) {
	actual := Vec3Create()
	Vec3Copy(actual, vecA)
	expect := vecA
	if !testSlice(actual, expect) {
		t.Errorf("copy: %v", actual)
	}
	if &actual == &expect {
		t.Errorf("copy: %v", actual)
	}
}

func TestVec3Set(t *testing.T) {
	actual := Vec3Set(Vec3Create(), 1, 2, 3)
	expect := []float64{1, 2, 3}
	if !testSlice(actual, expect) {
		t.Errorf("set: %v", actual)
	}
}

func TestVec3Add(t *testing.T) {
	actual := Vec3Add(Vec3Create(), vecA, vecB)
	expect := []float64{5, 7, 9}
	if !testSlice(actual, expect) {
		t.Errorf("add: %v", actual)
	}
}

func TestVec3Subtract(t *testing.T) {
	actual := Vec3Subtract(Vec3Create(), vecA, vecB)
	expect := []float64{-3, -3, -3}
	if !testSlice(actual, expect) {
		t.Errorf("subtract: %v", actual)
	}
}

func TestVec3Multiply(t *testing.T) {
	actual := Vec3Multiply(Vec3Create(), vecA, vecB)
	expect := []float64{4, 10, 18}
	if !testSlice(actual, expect) {
		t.Errorf("multiply: %v", actual)
	}
}

func TestVec3Divide(t *testing.T) {
	actual := Vec3Divide(Vec3Create(), vecA, vecB)
	expect := []float64{0.25, 0.4, 0.5}
	if !testSlice(actual, expect) {
		t.Errorf("divide: %v", actual)
	}
}

func TestVec3Ceil(t *testing.T) {
	actual := Vec3Ceil(Vec3Create(), []float64{math.E, math.Pi, math.Sqrt2})
	expect := []float64{3, 4, 2}
	if !testSlice(actual, expect) {
		t.Errorf("ceil: %v", actual)
	}
}

func TestVec3Floor(t *testing.T) {
	actual := Vec3Floor(Vec3Create(), []float64{math.E, math.Pi, math.Sqrt2})
	expect := []float64{2, 3, 1}
	if !testSlice(actual, expect) {
		t.Errorf("floor: %v", actual)
	}
}

func TestVec3Min(t *testing.T) {
	vecA := []float64{1, 3, 1}
	vecB := []float64{3, 1, 3}
	actual := Vec3Min(Vec3Create(), vecA, vecB)
	expect := []float64{1, 1, 1}
	if !testSlice(actual, expect) {
		t.Errorf("min: %v", actual)
	}
}

func TestVec3Max(t *testing.T) {
	vecA := []float64{1, 3, 1}
	vecB := []float64{3, 1, 3}
	actual := Vec3Max(Vec3Create(), vecA, vecB)
	expect := []float64{3, 3, 3}
	if !testSlice(actual, expect) {
		t.Errorf("max: %v", actual)
	}
}

func TestVec3Round(t *testing.T) {
	actual := Vec3Round(Vec3Create(), []float64{math.E, math.Pi, math.Sqrt2})
	expect := []float64{3, 3, 1}
	if !testSlice(actual, expect) {
		t.Errorf("round: %v", actual)
	}
}

func TestVec3Scale(t *testing.T) {
	actual := Vec3Scale(Vec3Create(), vecA, 2)
	expect := []float64{2, 4, 6}
	if !testSlice(actual, expect) {
		t.Errorf("scale: %v", actual)
	}
}

func TestVec3ScaleAndAdd(t *testing.T) {
	actual := Vec3ScaleAndAdd(Vec3Create(), vecA, vecB, 0.5)
	expect := []float64{3, 4.5, 6}
	if !testSlice(actual, expect) {
		t.Errorf("scale and add: %v", actual)
	}
}

func TestVec3Distance(t *testing.T) {
	actual := Vec3Distance(vecA, vecB)
	expect := 5.196152
	if !equals(actual, expect) {
		t.Errorf("dist: %v", actual)
	}
}

func TestVec3SquaredDistance(t *testing.T) {
	actual := Vec3SquaredDistance(vecA, vecB)
	expect := 27.
	if !equals(actual, expect) {
		t.Errorf("squared dist: %v", actual)
	}
}

func TestVec3Length(t *testing.T) {
	actual := Vec3Length(vecA)
	expect := 3.741657
	if !equals(actual, expect) {
		t.Errorf("length: %v", actual)
	}
	actual = Vec3Len(vecA)
	if !equals(actual, expect) {
		t.Errorf("len: %v", actual)
	}
}

func TestVec3SquaredLength(t *testing.T) {
	actual := Vec3SquaredLength(vecA)
	expect := 14.
	if !equals(actual, expect) {
		t.Errorf("squared length: %v", actual)
	}
	actual = Vec3SqrLen(vecA)
	if !equals(actual, expect) {
		t.Errorf("sqrlen: %v", actual)
	}
}

func TestVec3Negate(t *testing.T) {
	actual := Vec3Negate(Vec3Create(), vecA)
	expect := []float64{-1, -2, -3}
	if !testSlice(actual, expect) {
		t.Errorf("negate: %v", actual)
	}
}

func TestVec3Normalize(t *testing.T) {
	vecA := []float64{5, 0, 0}
	actual := Vec3Normalize(Vec3Create(), vecA)
	expect := []float64{1, 0, 0}
	if !testSlice(actual, expect) {
		t.Errorf("normalize: %v", actual)
	}
}

func TestVec3Dot(t *testing.T) {
	actual := Vec3Dot(vecA, vecB)
	expect := 32.
	if !equals(actual, expect) {
		t.Errorf("dot: %v", actual)
	}
}

func TestVec3Cross(t *testing.T) {
	actual := Vec3Cross(Vec3Create(), vecA, vecB)
	expect := []float64{-3, 6, -3}
	if !testSlice(actual, expect) {
		t.Errorf("cross: %v", actual)
	}
}
