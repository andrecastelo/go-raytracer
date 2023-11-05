package main

import (
	"andrecastelo/raytracer/internal/canvas"
	. "andrecastelo/raytracer/internal/clamp"
	"andrecastelo/raytracer/internal/tuple"
	"log"
	"math"
	"os"
    "fmt"
)

const (
	IMAGE_HEIGHT = 512
	IMAGE_WIDTH  = 512
    FILENAME = "projectile"
)

// writePoint will write the pixel to the canvas, inverting Y and making sure
// they don't exceed the boundaries of IMAGE_WIDTH x IMAGE_HEIGHT
func writePoint(c *canvas.Canvas, point *tuple.Tuple, color *tuple.Tuple) {
    x := math.Round(Clamp(point.X, IMAGE_WIDTH - 1, 0));
    y := math.Round(Clamp(IMAGE_HEIGHT - point.Y, IMAGE_HEIGHT - 1, 0));

    c.WritePixel(int(x), int(y), color)
}

// This command will paint a projectile affected by gravity. lmao
func main() {
    filename := fmt.Sprintf("%s.ppm", FILENAME)
	f, err := os.Create(filename)

	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to create file: %s", filename))
	}

	defer f.Close()

    c := canvas.MakeCanvas(IMAGE_WIDTH, IMAGE_HEIGHT)

    white := tuple.Color(1, 1, 1)
    projectile := tuple.Point(0, 0, 0)
    force := tuple.Vector(10, 10, 0)
    gravity := tuple.Vector(0, 0.2, 0)

    pixelsDrawn := 0

    for projectile.X < IMAGE_WIDTH && projectile.X >= 0 && 
        projectile.Y < IMAGE_HEIGHT && projectile.Y >= 0 {
        writePoint(c, projectile, white)
        projectile = projectile.Add(force)
        force = force.Subtract(gravity)
        pixelsDrawn++
        fmt.Printf("\rPixels drawn: %d", pixelsDrawn)
    }

    ppmString := c.Ppm()
    f.WriteString(ppmString)
}

