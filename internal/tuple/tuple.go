package tuple

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
