package tuple

type Tuple struct {
	x, y, z float64
	w       int
}

func Point(x float64, y float64, z float64) *Tuple {
	return &Tuple{x: x, y: y, z: z, w: 1}
}

func Vector(x float64, y float64, z float64) *Tuple {
	return &Tuple{x: x, y: y, z: z, w: 0}
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
