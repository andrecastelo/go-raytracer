package canvas

import (
	"andrecastelo/raytracer/internal/tuple"
)

type Canvas struct {
    pixels [][]tuple.Tuple
} 

func MakeCanvas(size int) *Canvas {
    canvas := &Canvas{}

    canvas.pixels = make([][]tuple.Tuple, size)
    for i := 0; i < size; i++ {
        canvas.pixels[i] = make([]tuple.Tuple, size)
    }

    return canvas
}
