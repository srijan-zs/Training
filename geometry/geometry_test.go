package geometry

import "testing"

// TestPerimeter function contains the test cases for perimeter method in Geometry interface of the geometry package
func TestPerimeter(t *testing.T) {
	cases := []struct {
		desc   string
		input  Geometry
		output float64
	}{
		{"negative input", Rectangle{-2, -1}, 0},
		{"negative input", Circle{-2}, 0},
		{"negative input", Square{-3}, 0},
		{"zero input", Rectangle{0, 0}, 0},
		{"zero input", Circle{0}, 0},
		{"zero input", Square{0}, 0},
		{"case rectangle", Rectangle{3, 5}, 16},
		{"case circle", Circle{7}, 43.982297150257104},
		{"case square", Square{8}, 32},
	}

	for i, tc := range cases {
		output := tc.input.Perimeter()
		if output != tc.output {
			t.Errorf("TEST[%d], failed. %s\nExpected: %v\nGot: %v", i+1, tc.desc, tc.output, output)
		}
	}
}

// TestArea function contains the test cases for Area method in Geometry interface of the geometry package
func TestArea(t *testing.T) {
	cases := []struct {
		desc   string
		input  Geometry
		output float64
	}{
		{"negative input", Rectangle{-3, -1}, 0},
		{"negative input", Circle{-4}, 0},
		{"negative input", Square{-2}, 0},
		{"zero input", Rectangle{0, 0}, 0},
		{"zero input", Circle{0}, 0},
		{"zero input", Square{0}, 0},
		{"case rectangle", Rectangle{3, 4}, 12},
		{"case circle", Circle{7}, 153.93804002589985},
		{"case square", Square{5}, 25},
	}

	for i, tc := range cases {
		output := tc.input.Area()
		if output != tc.output {
			t.Errorf("TEST[%d], failed. %s\nExpected: %v\nGot: %v", i+1, tc.desc, tc.output, output)
		}
	}
}

// BenchmarkPerimeter function runs the benchmark tests for the Perimeter method in Geometry interface of the geometry package
func BenchmarkPerimeter(b *testing.B) {
	cases := []struct {
		desc   string
		input  Geometry
		output float64
	}{
		{"negative input", Rectangle{-2, -1}, 0},
		{"negative input", Circle{-2}, 0},
		{"negative input", Square{-3}, 0},
		{"zero input", Rectangle{0, 0}, 0},
		{"zero input", Circle{0}, 0},
		{"zero input", Square{0}, 0},
		{"case rectangle", Rectangle{3, 5}, 16},
		{"case circle", Circle{7}, 43.982297150257104},
		{"case square", Square{8}, 32},
	}

	for _, tc := range cases {
		for i := 0; i < b.N; i++ {
			tc.input.Perimeter()
		}
	}
}

// BenchmarkArea function runs the benchmark tests for the Area method in Geometry interface of the geometry package
func BenchmarkArea(b *testing.B) {
	cases := []struct {
		desc   string
		input  Geometry
		output float64
	}{
		{"negative input", Rectangle{-3, -1}, 0},
		{"negative input", Circle{-4}, 0},
		{"negative input", Square{-2}, 0},
		{"zero input", Rectangle{0, 0}, 0},
		{"zero input", Circle{0}, 0},
		{"zero input", Square{0}, 0},
		{"case rectangle", Rectangle{3, 4}, 12},
		{"case circle", Circle{7}, 153.93804002589985},
		{"case square", Square{5}, 25},
	}

	for _, tc := range cases {
		for i := 0; i < b.N; i++ {
			tc.input.Area()
		}
	}
}
