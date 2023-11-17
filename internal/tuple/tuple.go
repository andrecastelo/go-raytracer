package tuple

import "math"

type Tuple struct {
	X, Y, Z float64
	W       int
}

// Instantiates a tuple
func MakeTuple(x float64, y float64, z float64, w int) *Tuple {
	return &Tuple{X: x, Y: y, Z: z, W: w}
}

// Returns a *Tuple with the last element being 1, making it a point
func Point(x float64, y float64, z float64) *Tuple {
	return MakeTuple(x, y, z, 1)
}

// Returns a *Tuple with the last element being 0.
func Vector(x float64, y float64, z float64) *Tuple {
	return MakeTuple(x, y, z, 0)
}

// Color returns a *Tuple with the last element being 0, since it will not
// be used in color calculations. Here, x = r, y = g, z = b.
func Color(r float64, g float64, b float64) *Tuple {
	return MakeTuple(r, g, b, 0)
}

// Returns the sum of two vectors
func (t *Tuple) Add(t2 *Tuple) *Tuple {
	return &Tuple{
		X: t.X + t2.X,
		Y: t.Y + t2.Y,
		Z: t.Z + t2.Z,
		W: t.W + t2.W,
	}
}

// Subtracts a vector from another
func (t *Tuple) Subtract(t2 *Tuple) *Tuple {
	return &Tuple{
		X: t.X - t2.X,
		Y: t.Y - t2.Y,
		Z: t.Z - t2.Z,
		W: t.W - t2.W,
	}
}

// Inverts the vector
func (t *Tuple) Negate() *Tuple {
	return &Tuple{
		X: t.X * -1,
		Y: t.Y * -1,
		Z: t.Z * -1,
		W: t.W * -1,
	}
}

// Multiplies a vector by a scalar value.
func (t *Tuple) Multiply(scalar float64) *Tuple {
	return &Tuple{
		X: t.X * scalar,
		Y: t.Y * scalar,
		Z: t.Z * scalar,
		W: int(float64(t.W) * scalar),
	}
}

// Divides a vector by a scalar and returns a new vector.
func (t *Tuple) Divide(scalar float64) *Tuple {
	return &Tuple{
		X: t.X / scalar,
		Y: t.Y / scalar,
		Z: t.Z / scalar,
		W: int(float64(t.W) / scalar),
	}
}

// Magnitude should be pretty self explanatory. Returns the "size" of
// the vector.
func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2))
}

// Normalize will, well, return the normalized vector
func (t *Tuple) Normalize() *Tuple {
	magnitude := t.Magnitude()

	return &Tuple{
		X: t.X / magnitude,
		Y: t.Y / magnitude,
		Z: t.Z / magnitude,
		W: 0,
	}
}

// Dot calculates the dot product between two vectors
func (t *Tuple) Dot(t2 *Tuple) float64 {
	return t.X*t2.X + t.Y*t2.Y + t.Z*t2.Z + float64(t.W*t2.W)
}

// Cross will calculate the cross product between two vectors and return a
// new vector
func (t *Tuple) Cross(t2 *Tuple) *Tuple {
	return Vector(
		(t.Y*t2.Z)-(t.Z*t2.Y),
		(t.Z*t2.X)-(t.X*t2.Z),
		(t.X*t2.Y)-(t.Y*t2.X),
	)
}

// Hadamard Product (or Schur product) is a method of blending two colors
func (t *Tuple) Hadamard(t2 *Tuple) *Tuple {
	return Color(
		t.X*t2.X,
		t.Y*t2.Y,
		t.Z*t2.Z,
	)
}

// Useful when we want to break the tuple into RGB. We discard the W
// component here
func (t *Tuple) Array() [4]float64 {
	return [...]float64{t.X, t.Y, t.Z, float64(t.W)}
}
