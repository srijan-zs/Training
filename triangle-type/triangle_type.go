package triangle

// checkTriangle function checks for the sides length and returns the respective type for the given sides of triangle
func checkTriangle(sides [3]int) string {
	if sides[0] <= 0 || sides[1] <= 0 || sides[2] <= 0 {
		return ""
	}

	if sides[0]+sides[1] <= sides[2] || sides[1]+sides[2] <= sides[0] || sides[2]+sides[0] <= sides[1] {
		return ""
	}

	if sides[0] == sides[1] && sides[1] == sides[2] {
		return "Equilateral"
	} else if sides[0] == sides[1] || sides[1] == sides[2] || sides[2] == sides[0] {
		return "Isosceles"
	}

	return "Scalene"
}
