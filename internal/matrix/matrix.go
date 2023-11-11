package matrix

import (
	"andrecastelo/raytracer/internal/compare"
	"errors"
)

type Matrix = [][]float64

func MakeMatrix(elements [][]float64) (Matrix, error) {
	height := len(elements)
	width := len(elements[0])

	for i := 0; i < height; i++ {
		if len(elements[i]) != width {
			return nil, errors.New("Matrix does not have the same width for all rows")
		}
	}

	return elements, nil
}

func Equal(a Matrix, b Matrix) bool {
	for i := range a {
		for j := range a[i] {
			equal := compare.Equal(a[i][j], b[i][j])
			if !equal {
				return false
			}
		}
	}

	return true
}
