package canvas

import (
	// "andrecastelo/raytracer/internal/tuple"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CanvasSuite struct {
    suite.Suite
}

func (suite *CanvasSuite) TestCanvasCreation() {
    canvas := MakeCanvas(32)

    suite.Len(canvas.pixels, 32)

    for _, line := range canvas.pixels {
        suite.Len(line, 32)
    } 

    firstPixel := canvas.pixels[0][0]
	suite.Equal(firstPixel.Array(), [...]float64{0.0, 0.0, 0.0})
}

func TestCanvasSuite(t *testing.T) {
    suite.Run(t, new(CanvasSuite))
}
