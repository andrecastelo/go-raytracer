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

// Multiply will multiply two 4x4 matrices and return a new one
func Multiply(a Matrix, b Matrix) Matrix {
	result := make([][]float64, 4)
	for line := 0; line < 4; line++ {
		result[line] = make([]float64, 4)
		for col := 0; col < 4; col++ {
			result[line][col] = a[line][0] * b[0][col]
			result[line][col] += a[line][1] * b[1][col]
			result[line][col] += a[line][2] * b[2][col]
			result[line][col] += a[line][3] * b[3][col]
		}
	}

	return result
}
