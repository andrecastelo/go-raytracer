package compare

import "math"

// Values with differences smaller than Epsilon are considered equal
const Epsilon = 1e-9

func Equal(a float64, b float64) bool {
	return math.Abs(a-b) < Epsilon
}
