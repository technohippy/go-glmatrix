package glmatrix

import (
	"math"
	"testing"
)

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

func TestVec3Distance(t *testing.T) {
	vecA := []float64{1, 2, 3}
	vecB := []float64{4, 5, 6}
	actual := Vec3Distance(vecA, vecB)
	if !equals(actual, 5.196152) {
		t.Errorf("dist: %v", actual)
	}
}
