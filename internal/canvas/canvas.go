package canvas

import (
	"andrecastelo/raytracer/internal/tuple"
	"fmt"
)

// Canvas is a struct that will hold a matrix of tuple.Tuple.
type Canvas struct {
    width  int
    height int
    pixels [][]tuple.Tuple
} 

func MakeCanvas(width int, height int) *Canvas {
    canvas := &Canvas{}

    canvas.width = width
    canvas.height = height
    canvas.pixels = make([][]tuple.Tuple, width)
    for i := 0; i < width; i++ {
        canvas.pixels[i] = make([]tuple.Tuple, height)
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
    canvas.pixels[x][y] = *color

    return canvas
}

func (canvas *Canvas) PpmArray() []string  {
    var ppm []string

    ppm = append(ppm, "P3")
    ppm = append(ppm, fmt.Sprintf("%d %d", canvas.width, canvas.height))
    ppm = append(ppm, "255")

    for _, line := range canvas.pixels {
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

func (canvas *Canvas) Ppm() string {
    var ppmString string
    ppm := canvas.PpmArray()

    for _, line := range ppm {
        ppmString += line
        ppmString += "\n"
    }

    return ppmString
}
