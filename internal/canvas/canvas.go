package canvas

import (
	"andrecastelo/raytracer/internal/tuple"
	"fmt"
	"strings"
)

// Canvas is a struct that will hold a matrix of tuple.Tuple.
type Canvas struct {
    Width  int
    Height int
    Pixels [][]tuple.Tuple
} 

func MakeCanvas(width int, height int) *Canvas {
    canvas := &Canvas{}

    canvas.Width = width
    canvas.Height = height
    canvas.Pixels = make([][]tuple.Tuple, height)
    for i := 0; i < height; i++ {
        canvas.Pixels[i] = make([]tuple.Tuple, width)
    }

    return canvas
}

func clamp(value int, max_value int, min_value int) int {
    if value < min_value {
        return min_value
    } else if value > max_value {
        return max_value
    } else {
        return value
    }
}

func (canvas *Canvas) WritePixel(x int, y int, color *tuple.Tuple) *Canvas {
    canvas.Pixels[y][x] = *color

    return canvas
}

func (canvas *Canvas) PpmArray() []string  {
    var ppm []string

    ppm = append(ppm, "P3")
    ppm = append(ppm, fmt.Sprintf("%d %d", canvas.Width, canvas.Height))
    ppm = append(ppm, "255")

    for _, line := range canvas.Pixels {
        for _, pixel := range line {
            array := pixel.Array()
			ir := clamp(int(255.999 * array[0]), 255, 0)
			ig := clamp(int(255.999 * array[1]), 255, 0)
			ib := clamp(int(255.999 * array[2]), 255, 0)

            line := fmt.Sprintf("%d %d %d", ir, ig, ib)
            ppm = append(ppm, line)
        }
    }

    return ppm
}

// Ppm will convert the canvas to the PPM file format. 
func (canvas *Canvas) Ppm() string {
    return strings.Join(canvas.PpmArray(), "\n") + "\n"
}
