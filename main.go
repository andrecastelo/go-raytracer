package main

import (
	"fmt"
	"log"
	"os"
)

const (
	IMAGE_HEIGHT = 512
	IMAGE_WIDTH  = 512
)

func main() {
	f, err := os.Create("test.ppm")

	if err != nil {
		log.Fatal("Unable to create file: test.ppm")
	}

	defer f.Close()

	f.WriteString("P3\n")
	f.WriteString(fmt.Sprintf("%d %d\n", IMAGE_WIDTH, IMAGE_HEIGHT))
	f.WriteString("255\n")

	for i := 0; i < IMAGE_HEIGHT; i++ {
		fmt.Printf("\rScanlines remaining: %d", IMAGE_HEIGHT-i)
		for j := 0; j < IMAGE_WIDTH; j++ {
			r := float64(i) / (IMAGE_WIDTH - 1)
			g := float64(j) / (IMAGE_HEIGHT - 1)
			b := 0.6

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			f.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}
}
