package glmatrix

import (
	"math"
	"testing"
)

func testSlice(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, ai := range a {
		if !equals(ai, b[i]) {
			return false
		}
	}
	return true
}

func TestToRadian(t *testing.T) {
	actual := ToRadian(90)
	expect := math.Pi / 2
	if !equals(actual, expect) {
		t.Errorf("deg:90 rad:%v", actual)
	}
}
