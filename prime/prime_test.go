package prime

import (
	"reflect"
	"testing"
)

// TestFindPrime function contains the test cases which tests the findPrime function of prime package
func TestFindPrime(t *testing.T) {
	cases := []struct {
		desc   string
		input  int
		output []int
	}{
		{"small positive integer", 15, []int{2, 3, 5, 7, 11, 13}},
		{"zero input", 0, nil},
		{"negative integer", -5, nil},
		{"large positive integer", 50, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}},
	}

	for i, tc := range cases {
		output := findPrime(tc.input)

		if !reflect.DeepEqual(output, tc.output) {
			t.Errorf("TEST[%d], failed. %s\nExpected: %v\nGot: %v", i+1, tc.desc, tc.output, output)
		}
	}
}

// BenchmarkFindPrime function runs the benchmark tests for the findPrime function of prime package
func BenchmarkFindPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findPrime(1000)
	}
}
