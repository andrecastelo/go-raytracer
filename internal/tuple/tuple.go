package tuple

type Tuple struct {
	x, y, z float64
	w       int
}

func Point(x float64, y float64, z float64) *Tuple {
	return &Tuple{x: x, y: y, z: z, w: 0}
}
