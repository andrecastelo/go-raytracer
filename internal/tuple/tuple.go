package tuple

import "math"

type Tuple struct {
	x, y, z float64
	w       int
}

func MakeTuple(x float64, y float64, z float64, w int) *Tuple {
	return &Tuple{x: x, y: y, z: z, w: w}
}

func Point(x float64, y float64, z float64) *Tuple {
	return MakeTuple(x, y, z, 1)
}

func Vector(x float64, y float64, z float64) *Tuple {
	return MakeTuple(x, y, z, 0)
}

func Color(r float64, g float64, b float64) *Tuple {
	return MakeTuple(r, g, b, 0)
}

func (t *Tuple) Add(t2 *Tuple) *Tuple {
	return &Tuple{
		x: t.x + t2.x,
		y: t.y + t2.y,
		z: t.z + t2.z,
		w: t.w + t2.w,
	}
}

func (t *Tuple) Subtract(t2 *Tuple) *Tuple {
	return &Tuple{
		x: t.x - t2.x,
		y: t.y - t2.y,
		z: t.z - t2.z,
		w: t.w - t2.w,
	}
}

func (t *Tuple) Negate() *Tuple {
	return &Tuple{
		x: t.x * -1,
		y: t.y * -1,
		z: t.z * -1,
		w: t.w * -1,
	}
}

func (t *Tuple) Multiply(scalar float64) *Tuple {
	return &Tuple{
		x: t.x * scalar,
		y: t.y * scalar,
		z: t.z * scalar,
		w: int(float64(t.w) * scalar),
	}
}

func (t *Tuple) Divide(scalar float64) *Tuple {
	return &Tuple{
		x: t.x / scalar,
		y: t.y / scalar,
		z: t.z / scalar,
		w: int(float64(t.w) / scalar),
	}
}

func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.x, 2) + math.Pow(t.y, 2) + math.Pow(t.z, 2))
}

func (t *Tuple) Normalize() *Tuple {
	magnitude := t.Magnitude()

	return &Tuple{
		x: t.x / magnitude,
		y: t.y / magnitude,
		z: t.z / magnitude,
		w: 0,
	}
}

func (t *Tuple) Dot(t2 *Tuple) float64 {
	return t.x*t2.x + t.y*t2.y + t.z*t2.z + float64(t.w*t2.w)
}

func (t *Tuple) Cross(t2 *Tuple) *Tuple {
	return Vector(
		(t.y*t2.z)-(t.z*t2.y),
		(t.z*t2.x)-(t.x*t2.z),
		(t.x*t2.y)-(t.y*t2.x),
	)
}

// Hadamard Product (or Schur product) is a method of blending two colors
func (t *Tuple) Hadamard(t2 *Tuple) *Tuple {
	return Color(
		t.x*t2.x,
		t.y*t2.y,
		t.z*t2.z,
	)
}
