package glmatrix

import (
	"math"
	"testing"
)

var vec3A = []float64{1, 2, 3}
var vec3B = []float64{4, 5, 6}

func TestVec3RotateX(t *testing.T) {
	actual := Vec3Create()
	vec3A := []float64{0, 1, 0}
	vec3B := []float64{0, 0, 0}
	Vec3RotateX(actual, vec3A, vec3B, math.Pi)
	if !testSlice(actual, []float64{0, -1, 0}) {
		t.Errorf("rotateX: %v", actual)
	}
}

func TestVec3TransformMat4(t *testing.T) {
	matr := Mat4LookAt(Mat4Create(), []float64{5, 6, 7}, []float64{2, 6, 7}, []float64{0, 1, 0})
	vec3A := []float64{1, 2, 3}
	actual := Vec3Create()
	Vec3TransformMat4(actual, vec3A, matr)
	if !testSlice(actual, []float64{4, -4, -4}) {
		t.Errorf("lookAt should rotate and translate the input: %v", actual)
	}
}

func TestVec3TransformMat3(t *testing.T) {
	actual := Vec3Create()
	vec3A := []float64{1, 2, 3}
	matr := []float64{1, 0, 0, 0, 1, 0, 0, 0, 1}
	Vec3TransformMat3(actual, vec3A, matr)
	if !testSlice(actual, vec3A) {
		t.Errorf("transform with an identity: %v", actual)
	}

	actual = Vec3Create()
	vec3A = []float64{0, 1, 0}
	matr = []float64{1, 0, 0, 0, 0, 1, 0, -1, 0}
	Vec3TransformMat3(actual, vec3A, matr)
	if !testSlice(actual, []float64{0, 0, 1}) {
		t.Errorf("transform with 90deg about X: %v", actual)
	}

	actual = Vec3Create()
	vec3A = []float64{1, 0, 0}
	matr = []float64{0, 0, -1, 0, 1, 0, 1, 0, 0}
	Vec3TransformMat3(actual, vec3A, matr)
	if !testSlice(actual, []float64{0, 0, -1}) {
		t.Errorf("transform with 90deg about Y: %v", actual)
	}

	actual = Vec3Create()
	vec3A = []float64{1, 0, 0}
	matr = Mat4LookAt(Mat4Create(), []float64{5, 6, 7}, []float64{2, 6, 7}, []float64{0, 1, 0})
	n := Mat3Create()
	matr = Mat3Transpose(n, Mat3Invert(n, Mat3FromMat4(n, matr)))
	Vec3TransformMat3(actual, vec3A, matr)
	if !testSlice(actual, []float64{0, 0, 1}) {
		t.Errorf("transform with a lookAt normal matrix: %v", actual)
	}
}

func TestVec3TransformQuat(t *testing.T) {
	actual := Vec3TransformQuat(Vec3Create(), vec3A, []float64{0.18257418567011074, 0.3651483713402215, 0.5477225570103322, 0.730296742680443})
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
	actual := Vec3Clone(vec3A)
	expect := vec3A
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
	Vec3Copy(actual, vec3A)
	expect := vec3A
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
	actual := Vec3Add(Vec3Create(), vec3A, vec3B)
	expect := []float64{5, 7, 9}
	if !testSlice(actual, expect) {
		t.Errorf("add: %v", actual)
	}
}

func TestVec3Subtract(t *testing.T) {
	actual := Vec3Subtract(Vec3Create(), vec3A, vec3B)
	expect := []float64{-3, -3, -3}
	if !testSlice(actual, expect) {
		t.Errorf("subtract: %v", actual)
	}
}

func TestVec3Multiply(t *testing.T) {
	actual := Vec3Multiply(Vec3Create(), vec3A, vec3B)
	expect := []float64{4, 10, 18}
	if !testSlice(actual, expect) {
		t.Errorf("multiply: %v", actual)
	}
}

func TestVec3Divide(t *testing.T) {
	actual := Vec3Divide(Vec3Create(), vec3A, vec3B)
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
	vec3A := []float64{1, 3, 1}
	vec3B := []float64{3, 1, 3}
	actual := Vec3Min(Vec3Create(), vec3A, vec3B)
	expect := []float64{1, 1, 1}
	if !testSlice(actual, expect) {
		t.Errorf("min: %v", actual)
	}
}

func TestVec3Max(t *testing.T) {
	vec3A := []float64{1, 3, 1}
	vec3B := []float64{3, 1, 3}
	actual := Vec3Max(Vec3Create(), vec3A, vec3B)
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
	actual := Vec3Scale(Vec3Create(), vec3A, 2)
	expect := []float64{2, 4, 6}
	if !testSlice(actual, expect) {
		t.Errorf("scale: %v", actual)
	}
}

func TestVec3ScaleAndAdd(t *testing.T) {
	actual := Vec3ScaleAndAdd(Vec3Create(), vec3A, vec3B, 0.5)
	expect := []float64{3, 4.5, 6}
	if !testSlice(actual, expect) {
		t.Errorf("scale and add: %v", actual)
	}
}

func TestVec3Distance(t *testing.T) {
	actual := Vec3Distance(vec3A, vec3B)
	expect := 5.196152
	if !equals(actual, expect) {
		t.Errorf("dist: %v", actual)
	}
}

func TestVec3SquaredDistance(t *testing.T) {
	actual := Vec3SquaredDistance(vec3A, vec3B)
	expect := 27.
	if !equals(actual, expect) {
		t.Errorf("squared dist: %v", actual)
	}
}

func TestVec3Length(t *testing.T) {
	actual := Vec3Length(vec3A)
	expect := 3.741657
	if !equals(actual, expect) {
		t.Errorf("length: %v", actual)
	}
	actual = Vec3Len(vec3A)
	if !equals(actual, expect) {
		t.Errorf("len: %v", actual)
	}
}

func TestVec3SquaredLength(t *testing.T) {
	actual := Vec3SquaredLength(vec3A)
	expect := 14.
	if !equals(actual, expect) {
		t.Errorf("squared length: %v", actual)
	}
	actual = Vec3SqrLen(vec3A)
	if !equals(actual, expect) {
		t.Errorf("sqrlen: %v", actual)
	}
}

func TestVec3Negate(t *testing.T) {
	actual := Vec3Negate(Vec3Create(), vec3A)
	expect := []float64{-1, -2, -3}
	if !testSlice(actual, expect) {
		t.Errorf("negate: %v", actual)
	}
}

func TestVec3Normalize(t *testing.T) {
	vec3A := []float64{5, 0, 0}
	actual := Vec3Normalize(Vec3Create(), vec3A)
	expect := []float64{1, 0, 0}
	if !testSlice(actual, expect) {
		t.Errorf("normalize: %v", actual)
	}
}

func TestVec3Dot(t *testing.T) {
	actual := Vec3Dot(vec3A, vec3B)
	expect := 32.
	if !equals(actual, expect) {
		t.Errorf("dot: %v", actual)
	}
}

func TestVec3Cross(t *testing.T) {
	actual := Vec3Cross(Vec3Create(), vec3A, vec3B)
	expect := []float64{-3, 6, -3}
	if !testSlice(actual, expect) {
		t.Errorf("cross: %v", actual)
	}
}

func TestVec3Lerp(t *testing.T) {
	actual := Vec3Lerp(Vec3Create(), vec3A, vec3B, 0.5)
	expect := []float64{2.5, 3.5, 4.5}
	if !testSlice(actual, expect) {
		t.Errorf("lerp: %v", actual)
	}
}

func TestVec3Slerp(t *testing.T) {
	actual := Vec3Slerp(Vec3Create(), []float64{1, 0, 0}, []float64{0, 1, 0}, 0)
	expect := []float64{1, 0, 0}
	if !testSlice(actual, expect) {
		t.Errorf("slerp: %v", actual)
	}
}

func TestVec3Random(t *testing.T) {
	actual := Vec3Random(Vec3Create(), 1.)
	expect := 1.
	if !equals(Vec3Len(actual), expect) {
		t.Errorf("random: %v", Vec3Len(actual))
	}

	actual = Vec3Random(Vec3Create(), 5.)
	expect = 5.
	if !equals(Vec3Len(actual), expect) {
		t.Errorf("random: %v", Vec3Len(actual))
	}
}

func TestVec3ForEach(t *testing.T) {
	vec3Array := []float64{
		1, 2, 3,
		4, 5, 6,
		0, 0, 0,
	}
	fn := func(a []float64, b []float64, c []float64) {
		Vec3Normalize(a, b)
	}
	actual := Vec3ForEach(vec3Array, 0., 0., 0., fn, []float64{})
	expect := []float64{
		0.267261, 0.534522, 0.801783,
		0.455842, 0.569802, 0.683763,
		0, 0, 0,
	}
	if !testSlice(actual, expect) {
		t.Errorf("forEach: %v", actual)
	}
}

func TestVec3Angle(t *testing.T) {
	actual := Vec3Angle(vec3A, vec3B)
	expect := 0.225726
	if !equals(actual, expect) {
		t.Errorf("angle: %v", actual)
	}
}

func TestVec3ExactEquals(t *testing.T) {
	vec3A := []float64{0, 1, 2}
	vec3B := []float64{0, 1, 2}
	vecC := []float64{0, 1, 2 + 1e-10}
	if !Vec3ExactEquals(vec3A, vec3B) {
		t.Errorf("exact equal")
	}
	if Vec3ExactEquals(vec3A, vecC) {
		t.Errorf("exact equal")
	}
}

func TestVec3Equals(t *testing.T) {
	vec3A := []float64{0, 1, 2}
	vec3B := []float64{0, 1, 2}
	vecC := []float64{0, 1, 2 + 1e-10}
	if !Vec3Equals(vec3A, vec3B) {
		t.Errorf("exact equal")
	}
	if !Vec3Equals(vec3A, vecC) {
		t.Errorf("exact equal")
	}
}

func TestVec3Zero(t *testing.T) {
	actual := []float64{1, 2, 3}
	Vec3Zero(actual)
	expect := []float64{0, 0, 0}
	if !testSlice(actual, expect) {
		t.Errorf("zero: %v", actual)
	}
}
