package triangle

import (
	"reflect"
	"testing"
)

// TestCheckTriangle function contains the test cases which tests the checkTriangle function of triangle package
func TestCheckTriangle(t *testing.T) {
	cases := []struct {
		desc   string
		input  [3]int
		output string
	}{
		{"all sides equal", [3]int{5, 5, 5}, "Equilateral"},
		{"two sides equal", [3]int{4, 4, 6}, "Isosceles"},
		{"all sides different", [3]int{3, 4, 5}, "Scalene"},
		{"negative input", [3]int{-5, 3, -5}, ""},
		{"zero input", [3]int{0, 4, 5}, ""},
		{"invalid triangle sides", [3]int{1, 2, 10}, ""},
	}

	for i, tc := range cases {
		output := checkTriangle(tc.input)

		if !reflect.DeepEqual(output, tc.output) {
			t.Errorf("TEST[%d] failed. %s\nExpected: %v\nGot: %v", i+1, tc.desc, tc.output, output)
		}
	}
}

// BenchmarkCheckTriangle runs the benchmark tests for the checkTriangle function of triangle package
func BenchmarkCheckTriangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkTriangle([3]int{2, 3, 4})
	}
}
