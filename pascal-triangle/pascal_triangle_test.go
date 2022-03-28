package pascal

import (
	"reflect"
	"testing"
)

// TestPascal function contains the test cases which tests the pascal function of pascal package
func TestPascal(t *testing.T) {
	cases := []struct {
		desc   string
		input  int
		output [][]int
	}{
		{"negative input", -4, nil},
		{"zero input", 0, nil},
		{"small integer", 3, [][]int{{1}, {1, 1}, {1, 2, 1}}},
		{"large integer", 7, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}, {1, 6, 15, 20, 15, 6, 1}}},
	}

	for i, tc := range cases {
		output := pascal(tc.input)

		if !reflect.DeepEqual(output, tc.output) {
			t.Errorf("TEST[%d], failed. %s\nExpected: %v\nGot: %v", i+1, tc.desc, tc.output, output)
		}
	}
}

// BenchmarkPascal function runs the benchmark tests for the pascal function of pascal package
func BenchmarkPascal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pascal(9)
	}
}
