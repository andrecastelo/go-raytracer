package matrix

import "errors"

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
