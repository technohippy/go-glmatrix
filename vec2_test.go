package glmatrix

import (
	"math"
	"testing"
)

var vec2A = []float64{1, 2}
var vec2B = []float64{3, 4}

func TestVec2Create(t *testing.T) {
	actual := Vec2Create()
	expect := []float64{0, 0}
	if !testSlice(actual, expect) {
		t.Errorf("create: %v", actual)
	}
}

func TestVec2Clone(t *testing.T) {
	actual := Vec2Clone(vec2A)
	expect := vec2A
	if !testSlice(actual, expect) {
		t.Errorf("clone: %v", actual)
	}
	if &actual == &expect {
		t.Errorf("clone: %v", actual)
	}
}

func TestVec2FromValues(t *testing.T) {
	actual := Vec2FromValues(1, 2)
	expect := []float64{1, 2}
	if !testSlice(actual, expect) {
		t.Errorf("from values: %v", actual)
	}
}

func TestVec2Copy(t *testing.T) {
	actual := Vec2Create()
	Vec2Copy(actual, vec2A)
	expect := vec2A
	if !testSlice(actual, expect) {
		t.Errorf("copy: %v", actual)
	}
	if &actual == &expect {
		t.Errorf("copy: %v", actual)
	}
}

func TestVec2Set(t *testing.T) {
	actual := Vec2Set(Vec2Create(), 1, 2)
	expect := []float64{1, 2}
	if !testSlice(actual, expect) {
		t.Errorf("set: %v", actual)
	}
}

func TestVec2Add(t *testing.T) {
	actual := Vec2Add(Vec2Create(), vec2A, vec2B)
	expect := []float64{4, 6}
	if !testSlice(actual, expect) {
		t.Errorf("add: %v", actual)
	}
}

func TestVec2Subtract(t *testing.T) {
	actual := Vec2Subtract(Vec2Create(), vec2A, vec2B)
	expect := []float64{-2, -2}
	if !testSlice(actual, expect) {
		t.Errorf("subtract: %v", actual)
	}
}

func TestVec2Multiply(t *testing.T) {
	actual := Vec2Multiply(Vec2Create(), vec2A, vec2B)
	expect := []float64{3, 8}
	if !testSlice(actual, expect) {
		t.Errorf("multiply: %v", actual)
	}
}

func TestVec2Divide(t *testing.T) {
	actual := Vec2Divide(Vec2Create(), vec2A, vec2B)
	expect := []float64{0.333333, 0.5}
	if !testSlice(actual, expect) {
		t.Errorf("divide: %v", actual)
	}
}

func TestVec2Ceil(t *testing.T) {
	actual := Vec2Ceil(Vec2Create(), []float64{math.E, math.Pi})
	expect := []float64{3, 4}
	if !testSlice(actual, expect) {
		t.Errorf("ceil: %v", actual)
	}
}

func TestVec2Floor(t *testing.T) {
	actual := Vec2Floor(Vec2Create(), []float64{math.E, math.Pi})
	expect := []float64{2, 3}
	if !testSlice(actual, expect) {
		t.Errorf("floor: %v", actual)
	}
}

func TestVec2Min(t *testing.T) {
	vec2A := []float64{1, 4}
	vec2B := []float64{3, 2}
	actual := Vec2Min(Vec2Create(), vec2A, vec2B)
	expect := []float64{1, 2}
	if !testSlice(actual, expect) {
		t.Errorf("min: %v", actual)
	}
}

func TestVec2Max(t *testing.T) {
	vec2A := []float64{1, 4}
	vec2B := []float64{3, 2}
	actual := Vec2Max(Vec2Create(), vec2A, vec2B)
	expect := []float64{3, 4}
	if !testSlice(actual, expect) {
		t.Errorf("max: %v", actual)
	}
}

func TestVec2Round(t *testing.T) {
	actual := Vec2Round(Vec2Create(), []float64{math.E, math.Pi})
	expect := []float64{3, 3}
	if !testSlice(actual, expect) {
		t.Errorf("round: %v", actual)
	}
}

func TestVec2Scale(t *testing.T) {
	actual := Vec2Scale(Vec2Create(), vec2A, 2)
	expect := []float64{2, 4}
	if !testSlice(actual, expect) {
		t.Errorf("scale: %v", actual)
	}
}

func TestVec2ScaleAndAdd(t *testing.T) {
	actual := Vec2ScaleAndAdd(Vec2Create(), vec2A, vec2B, 0.5)
	expect := []float64{2.5, 4}
	if !testSlice(actual, expect) {
		t.Errorf("scale and add: %v", actual)
	}
}

func TestVec2Distance(t *testing.T) {
	actual := Vec2Distance(vec2A, vec2B)
	expect := 2.82842712
	if !equals(actual, expect) {
		t.Errorf("dist: %v", actual)
	}
}

func TestVec2SquaredDistance(t *testing.T) {
	actual := Vec2SquaredDistance(vec2A, vec2B)
	expect := 8.
	if !equals(actual, expect) {
		t.Errorf("squared dist: %v", actual)
	}
}

func TestVec2Length(t *testing.T) {
	actual := Vec2Length(vec2A)
	expect := 2.236067
	if !equals(actual, expect) {
		t.Errorf("length: %v", actual)
	}
	actual = Vec2Len(vec2A)
	if !equals(actual, expect) {
		t.Errorf("len: %v", actual)
	}
}

func TestVec2SquaredLength(t *testing.T) {
	actual := Vec2SquaredLength(vec2A)
	expect := 5.
	if !equals(actual, expect) {
		t.Errorf("squared length: %v", actual)
	}
	actual = Vec2SqrLen(vec2A)
	if !equals(actual, expect) {
		t.Errorf("sqrlen: %v", actual)
	}
}

func TestVec2Negate(t *testing.T) {
	actual := Vec2Negate(Vec2Create(), vec2A)
	expect := []float64{-1, -2}
	if !testSlice(actual, expect) {
		t.Errorf("negate: %v", actual)
	}
}

func TestVec2Normalize(t *testing.T) {
	vec2A := []float64{5, 0}
	actual := Vec2Normalize(Vec2Create(), vec2A)
	expect := []float64{1, 0}
	if !testSlice(actual, expect) {
		t.Errorf("normalize: %v", actual)
	}
}

func TestVec2Dot(t *testing.T) {
	actual := Vec2Dot(vec2A, vec2B)
	expect := 11.
	if !equals(actual, expect) {
		t.Errorf("dot: %v", actual)
	}
}

func TestVec2Cross(t *testing.T) {
	actual := Vec2Cross(Vec3Create(), vec2A, vec2B)
	expect := []float64{0, 0, -2}
	if !testSlice(actual, expect) {
		t.Errorf("cross: %v", actual)
	}
}

func TestVec2Lerp(t *testing.T) {
	actual := Vec2Lerp(Vec2Create(), vec2A, vec2B, 0.5)
	expect := []float64{2, 3}
	if !testSlice(actual, expect) {
		t.Errorf("lerp: %v %v", actual, expect)
	}
}

func TestVec2Random(t *testing.T) {
	actual := Vec2Random(Vec2Create(), 1.)
	expect := 1.
	if !equals(Vec2Len(actual), expect) {
		t.Errorf("random: %v", Vec2Len(actual))
	}

	actual = Vec2Random(Vec2Create(), 5.)
	expect = 5.
	if !equals(Vec2Len(actual), expect) {
		t.Errorf("random: %v", Vec2Len(actual))
	}
}

func TestVec2TransformMat2(t *testing.T) {
	matA := []float64{
		1, 2,
		3, 4,
	}
	actual := Vec2TransformMat2(Vec2Create(), vec2A, matA)
	expect := []float64{7, 10}
	if !testSlice(actual, expect) {
		t.Errorf("transform mat2: %v %v", actual, expect)
	}
}

func TestVec2TransformMat2d(t *testing.T) {
	matA := []float64{
		1, 2,
		3, 4,
		5, 6,
	}
	actual := Vec2TransformMat2d(Vec2Create(), vec2A, matA)
	expect := []float64{12, 16}
	if !testSlice(actual, expect) {
		t.Errorf("transform mat2: %v %v", actual, expect)
	}
}

func TestVec2TransformMat3(t *testing.T) {
	matA := []float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	actual := Vec2TransformMat3(Vec2Create(), vec2A, matA)
	expect := []float64{14, 18}
	if !testSlice(actual, expect) {
		t.Errorf("transform mat3: %v %v", actual, expect)
	}
}

func TestVec2TransformMat4(t *testing.T) {
	matA := []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}
	actual := Vec2TransformMat4(Vec2Create(), vec2A, matA)
	expect := []float64{20, 24}
	if !testSlice(actual, expect) {
		t.Errorf("transform mat3: %v %v", actual, expect)
	}
}

func TestVec2Rotate(t *testing.T) {
	actual := Vec2Rotate(Vec2Create(), []float64{2, 0}, []float64{1, 0}, math.Pi/2)
	expect := []float64{1, 1}
	if !testSlice(actual, expect) {
		t.Errorf("rotate: %v %v", actual, expect)
	}
}

func TestVec2ForEach(t *testing.T) {
	vec2Array := []float64{
		1, 2,
		3, 4,
		0, 0,
	}
	fn := func(a []float64, b []float64, c []float64) {
		Vec2Normalize(a, b)
	}
	actual := Vec2ForEach(vec2Array, 0, 0, 0, fn, []float64{})
	expect := []float64{
		0.447214, 0.894427,
		0.6, 0.8,
		0, 0,
	}
	if !testSlice(actual, expect) {
		t.Errorf("forEach: %v", actual)
	}
}

func TestVec2Angle(t *testing.T) {
	vec2A := []float64{1, 0}
	vec2B := []float64{1, 2}
	actual := Vec2Angle(vec2A, vec2B)
	expect := 1.107148
	if !equals(actual, expect) {
		t.Errorf("angle: %v", actual)
	}
}

func TestVec2ExactEquals(t *testing.T) {
	vec2A := []float64{0, 1}
	vec2B := []float64{0, 1}
	vecC := []float64{0, 1 + 1e-10}
	if !Vec2ExactEquals(vec2A, vec2B) {
		t.Errorf("exact equal")
	}
	if Vec2ExactEquals(vec2A, vecC) {
		t.Errorf("exact equal")
	}
}

func TestVec2Equals(t *testing.T) {
	vec2A := []float64{0, 1}
	vec2B := []float64{0, 1}
	vecC := []float64{0, 1 + 1e-10}
	if !Vec2Equals(vec2A, vec2B) {
		t.Errorf("exact equal")
	}
	if !Vec2Equals(vec2A, vecC) {
		t.Errorf("exact equal")
	}
}

func TestVec2Zero(t *testing.T) {
	actual := []float64{1, 2}
	Vec2Zero(actual)
	expect := []float64{0, 0}
	if !testSlice(actual, expect) {
		t.Errorf("zero: %v", actual)
	}
}
