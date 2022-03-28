package leapyear

import "testing"

// TestCheckLeap function contains test cases which tests the checkLeap function of leapyear package
func TestCheckLeap(t *testing.T) {
	cases := []struct {
		desc   string
		input  int
		output bool
	}{
		{"divisible by 4 & 100", 2000, true},
		{"divisible by 4 but not 100", 1900, false},
		{"divisible only by 4", 1992, true},
		{"not divisible by 4", 1993, false},
		{"empty input", 0, false},
		{"negative input", -5, false},
	}

	for i, tc := range cases {
		output := checkLeap(tc.input)

		if output != tc.output {
			t.Errorf("TEST[%d], failed. %s\nExpected: %v\nGot: %v", i+1, tc.desc, tc.output, output)
		}
	}
}

// BenchmarkCheckLeap function runs the benchmark tests for the checkLeap function of leapyear package
func BenchmarkCheckLeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkLeap(1234)
	}
}
