package glmatrix

import (
	"math"
	"testing"
)

var vec4A = []float64{1, 2, 3, 4}
var vec4B = []float64{5, 6, 7, 8}

func TestVec4Create(t *testing.T) {
	actual := Vec4Create()
	expect := []float64{0, 0, 0, 0}
	if !testSlice(actual, expect) {
		t.Errorf("create: %v", actual)
	}
}

func TestVec4Clone(t *testing.T) {
	actual := Vec4Clone(vec4A)
	expect := vec4A
	if !testSlice(actual, expect) {
		t.Errorf("clone: %v", actual)
	}
	if &actual == &expect {
		t.Errorf("clone: %v", actual)
	}
}

func TestVec4FromValues(t *testing.T) {
	actual := Vec4FromValues(1, 2, 3, 4)
	expect := []float64{1, 2, 3, 4}
	if !testSlice(actual, expect) {
		t.Errorf("from values: %v", actual)
	}
}

func TestVec4Copy(t *testing.T) {
	actual := Vec4Create()
	Vec4Copy(actual, vec4A)
	expect := vec4A
	if !testSlice(actual, expect) {
		t.Errorf("copy: %v", actual)
	}
	if &actual == &expect {
		t.Errorf("copy: %v", actual)
	}
}

func TestVec4Set(t *testing.T) {
	actual := Vec4Set(Vec4Create(), 1, 2, 3, 4)
	expect := []float64{1, 2, 3, 4}
	if !testSlice(actual, expect) {
		t.Errorf("set: %v", actual)
	}
}

func TestVec4Add(t *testing.T) {
	actual := Vec4Add(Vec4Create(), vec4A, vec4B)
	expect := []float64{6, 8, 10, 12}
	if !testSlice(actual, expect) {
		t.Errorf("add: %v", actual)
	}
}

func TestVec4Subtract(t *testing.T) {
	actual := Vec4Subtract(Vec4Create(), vec4A, vec4B)
	expect := []float64{-4, -4, -4, -4}
	if !testSlice(actual, expect) {
		t.Errorf("subtract: %v", actual)
	}
}

func TestVec4Multiply(t *testing.T) {
	actual := Vec4Multiply(Vec4Create(), vec4A, vec4B)
	expect := []float64{5, 12, 21, 32}
	if !testSlice(actual, expect) {
		t.Errorf("multiply: %v", actual)
	}
}

func TestVec4Divide(t *testing.T) {
	actual := Vec4Divide(Vec4Create(), vec4A, vec4B)
	expect := []float64{0.2, 0.333333, 0.428571, 0.5}
	if !testSlice(actual, expect) {
		t.Errorf("divide: %v", actual)
	}
}

func TestVec4Ceil(t *testing.T) {
	actual := Vec4Ceil(Vec4Create(), []float64{math.E, math.Pi, math.Sqrt2, math.SqrtPi})
	expect := []float64{3, 4, 2, 2}
	if !testSlice(actual, expect) {
		t.Errorf("ceil: %v", actual)
	}
}

func TestVec4Floor(t *testing.T) {
	actual := Vec4Floor(Vec4Create(), []float64{math.E, math.Pi, math.Sqrt2, math.SqrtPi})
	expect := []float64{2, 3, 1, 1}
	if !testSlice(actual, expect) {
		t.Errorf("floor: %v", actual)
	}
}

func TestVec4Min(t *testing.T) {
	vec4A := []float64{1, 3, 1, 3}
	vec4B := []float64{3, 1, 3, 1}
	actual := Vec4Min(Vec4Create(), vec4A, vec4B)
	expect := []float64{1, 1, 1, 1}
	if !testSlice(actual, expect) {
		t.Errorf("min: %v", actual)
	}
}

func TestVec4Max(t *testing.T) {
	vec4A := []float64{1, 3, 1, 3}
	vec4B := []float64{3, 1, 3, 1}
	actual := Vec4Max(Vec4Create(), vec4A, vec4B)
	expect := []float64{3, 3, 3, 3}
	if !testSlice(actual, expect) {
		t.Errorf("max: %v", actual)
	}
}

func TestVec4Round(t *testing.T) {
	actual := Vec4Round(Vec4Create(), []float64{math.E, math.Pi, math.Sqrt2, math.SqrtPi})
	expect := []float64{3, 3, 1, 2}
	if !testSlice(actual, expect) {
		t.Errorf("round: %v", actual)
	}
}

func TestVec4Scale(t *testing.T) {
	actual := Vec4Scale(Vec4Create(), vec4A, 2)
	expect := []float64{2, 4, 6, 8}
	if !testSlice(actual, expect) {
		t.Errorf("scale: %v", actual)
	}
}

func TestVec4ScaleAndAdd(t *testing.T) {
	actual := Vec4ScaleAndAdd(Vec4Create(), vec4A, vec4B, 0.5)
	expect := []float64{3.5, 5, 6.5, 8}
	if !testSlice(actual, expect) {
		t.Errorf("scale and add: %v", actual)
	}
}

func TestVec4Distance(t *testing.T) {
	actual := Vec4Distance(vec4A, vec4B)
	expect := 8.
	if !equals(actual, expect) {
		t.Errorf("dist: %v", actual)
	}
}

func TestVec4SquaredDistance(t *testing.T) {
	actual := Vec4SquaredDistance(vec4A, vec4B)
	expect := 64.
	if !equals(actual, expect) {
		t.Errorf("squared dist: %v", actual)
	}
}

func TestVec4Length(t *testing.T) {
	actual := Vec4Length(vec4A)
	expect := 5.477225
	if !equals(actual, expect) {
		t.Errorf("length: %v", actual)
	}
	actual = Vec4Len(vec4A)
	if !equals(actual, expect) {
		t.Errorf("len: %v", actual)
	}
}

func TestVec4SquaredLength(t *testing.T) {
	actual := Vec4SquaredLength(vec4A)
	expect := 30.
	if !equals(actual, expect) {
		t.Errorf("squared length: %v", actual)
	}
	actual = Vec4SqrLen(vec4A)
	if !equals(actual, expect) {
		t.Errorf("sqrlen: %v", actual)
	}
}

func TestVec4Negate(t *testing.T) {
	actual := Vec4Negate(Vec4Create(), vec4A)
	expect := []float64{-1, -2, -3, -4}
	if !testSlice(actual, expect) {
		t.Errorf("negate: %v", actual)
	}
}

func TestVec4Normalize(t *testing.T) {
	vec4A := []float64{5, 0, 0, 0}
	actual := Vec4Normalize(Vec4Create(), vec4A)
	expect := []float64{1, 0, 0, 0}
	if !testSlice(actual, expect) {
		t.Errorf("normalize: %v", actual)
	}
}

func TestVec4Dot(t *testing.T) {
	actual := Vec4Dot(vec4A, vec4B)
	expect := 70.
	if !equals(actual, expect) {
		t.Errorf("dot: %v", actual)
	}
}

func TestVec4Cross(t *testing.T) {
	vec4A := []float64{1, 0, 0, 0}
	vec4B := []float64{0, 1, 0, 0}
	vec4C := []float64{0, 0, 1, 0}
	actual := Vec4Cross(Vec4Create(), vec4A, vec4B, vec4C)
	expect := []float64{0, 0, 0, -1}
	if !testSlice(actual, expect) {
		t.Errorf("cross: %v", actual)
	}
}

func TestVec4Lerp(t *testing.T) {
	actual := Vec4Lerp(Vec4Create(), vec4A, vec4B, 0.5)
	expect := []float64{3, 4, 5, 6}
	if !testSlice(actual, expect) {
		t.Errorf("lerp: %v", actual)
	}
}

func TestVec4Random(t *testing.T) {
	actual := Vec4Random(Vec4Create(), 1.)
	expect := 1.
	if !equals(Vec4Len(actual), expect) {
		t.Errorf("random: %v %v", Vec4Len(actual), actual)
	}

	actual = Vec4Random(Vec4Create(), 5.)
	expect = 5.
	if !equals(Vec4Len(actual), expect) {
		t.Errorf("random: %v %v", Vec4Len(actual), actual)
	}
}

func TestVec4ForEach(t *testing.T) {
	vec4Array := []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		0, 0, 0, 0,
	}
	fn := func(a []float64, b []float64, c []float64) {
		Vec4Normalize(a, b)
	}
	actual := Vec4ForEach(vec4Array, 0, 0, 0, fn, []float64{})
	expect := []float64{
		0.182574, 0.365148, 0.547722, 0.730296,
		0.379049, 0.454858, 0.530668, 0.606478,
		0, 0, 0, 0,
	}
	if !testSlice(actual, expect) {
		t.Errorf("forEach: %v", actual)
	}
}

func TestVec4ExactEquals(t *testing.T) {
	vec4A := []float64{0, 1, 2, 3}
	vec4B := []float64{0, 1, 2, 3}
	vecC := []float64{0, 1, 2, 3 + 1e-10}
	if !Vec4ExactEquals(vec4A, vec4B) {
		t.Errorf("exact equal")
	}
	if Vec4ExactEquals(vec4A, vecC) {
		t.Errorf("exact equal")
	}
}

func TestVec4Equals(t *testing.T) {
	vec4A := []float64{0, 1, 2, 3}
	vec4B := []float64{0, 1, 2, 3}
	vecC := []float64{0, 1, 2, 3 + 1e-10}
	if !Vec4Equals(vec4A, vec4B) {
		t.Errorf("exact equal")
	}
	if !Vec4Equals(vec4A, vecC) {
		t.Errorf("exact equal")
	}
}

func TestVec4Zero(t *testing.T) {
	actual := []float64{1, 2, 3, 4}
	Vec4Zero(actual)
	expect := []float64{0, 0, 0, 0}
	if !testSlice(actual, expect) {
		t.Errorf("zero: %v", actual)
	}
}
