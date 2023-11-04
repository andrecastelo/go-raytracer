package main

import (
	"andrecastelo/raytracer/internal/canvas"
	"andrecastelo/raytracer/internal/tuple"
	"log"
	"os"
)

const (
	IMAGE_HEIGHT = 512
	IMAGE_WIDTH  = 512
)

func main() {
	f, err := os.Create("gradient.ppm")

	if err != nil {
		log.Fatal("Unable to create file: gradient.ppm")
	}

	defer f.Close()

    c := canvas.MakeCanvas(IMAGE_WIDTH, IMAGE_HEIGHT)

    for x, line := range c.Pixels {
        for y, _ := range line {
			r := float64(x) / (IMAGE_WIDTH - 1)
			g := float64(y) / (IMAGE_HEIGHT - 1)
			b := 1.0 - (float64(x) / (IMAGE_HEIGHT - 1))

            c.WritePixel(x, y, tuple.Color(r, g, b))
        } 
    }

    ppmString := c.Ppm()
    f.WriteString(ppmString)
}
