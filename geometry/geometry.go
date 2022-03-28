package geometry

import "math"

type Geometry interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	len, bdt float64
}

// Perimeter method defined on type rectangle to calculate perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	if r.len <= 0 || r.bdt <= 0 {
		return 0
	}

	return 2 * (r.len + r.bdt)
}

// Area method defined on type rectangle to calculate area of a rectangle
func (r Rectangle) Area() float64 {
	if r.len <= 0 || r.bdt <= 0 {
		return 0
	}

	return r.len * r.bdt
}

type Circle struct {
	rad float64
}

// Perimeter method defined on type circle to calculate circumference of a circle
func (c Circle) Perimeter() float64 {
	if c.rad <= 0 {
		return 0
	}

	return 2 * math.Pi * c.rad
}

// Area method defined on type circle to calculate area of a circle
func (c Circle) Area() float64 {
	if c.rad <= 0 {
		return 0
	}

	return math.Pi * c.rad * c.rad
}

type Square struct {
	side float64
}

// Perimeter method defined on type square to calculate perimeter of a square
func (s Square) Perimeter() float64 {
	if s.side <= 0 {
		return 0
	}

	return 4 * s.side
}

// Area method defined on type square to calculate area of a square
func (s Square) Area() float64 {
	if s.side <= 0 {
		return 0
	}

	return s.side * s.side
}
