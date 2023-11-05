package canvas

import (
	// "andrecastelo/raytracer/internal/tuple"
	"andrecastelo/raytracer/internal/tuple"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CanvasSuite struct {
    suite.Suite
}

func (suite *CanvasSuite) TestCanvasCreation() {
    canvas := MakeCanvas(32, 16)

    suite.Equal(canvas.Width, 32)
    suite.Equal(canvas.Height, 16)
    suite.Len(canvas.Pixels, 32)

    for _, line := range canvas.Pixels {
        suite.Len(line, 16)
    } 

    firstPixel := canvas.Pixels[0][0]
	suite.Equal(firstPixel.Array(), [...]float64{0.0, 0.0, 0.0})
}

func (suite *CanvasSuite) TestWritingPixelToCanvas() {
    canvas := MakeCanvas(10, 10)
    red := tuple.Color(1.0, 0.0, 0.0)

    canvas.WritePixel(2, 3, red)

    suite.Equal(canvas.Pixels[3][2].Array(), [...]float64{1.0, 0.0, 0.0})
}

func (suite *CanvasSuite) TestCanvasPpmArray() {
    canvas := MakeCanvas(5, 3)
    c1 := tuple.Color(1.5, 0.0, 0.0)
    c2 := tuple.Color(0.0, 0.5, 0.0)
    c3 := tuple.Color(-0.5, 0.0, 1.0)
    canvas.WritePixel(0, 0, c1)
    canvas.WritePixel(1, 2, c2)
    canvas.WritePixel(2, 4, c3)
    ppm := canvas.PpmArray()
    
    suite.Equal("P3", ppm[0])
    suite.Equal("5 3", ppm[1])
    suite.Equal("255", ppm[2])
    suite.Equal("255 0 0", ppm[3])
    suite.Equal("0 127 0", ppm[10])
    suite.Equal("0 0 255", ppm[17])
}

func (suite *CanvasSuite) TestCanvasPpmString() {
    canvas := MakeCanvas(3, 3)
    c1 := tuple.Color(1.0, 0.0, 0.0)
    c2 := tuple.Color(0.0, 0.5, 0.0)
    c3 := tuple.Color(-0.5, 0.0, 1.0)
    canvas.WritePixel(0, 0, c1)
    canvas.WritePixel(0, 1, c2)
    canvas.WritePixel(0, 2, c3)
    ppm := canvas.Ppm()

    expected := `P3
3 3
255
255 0 0
0 0 0
0 0 0
0 127 0
0 0 0
0 0 0
0 0 255
0 0 0
0 0 0
`
    suite.Equal(expected, ppm)
}

func TestCanvasSuite(t *testing.T) {
    suite.Run(t, new(CanvasSuite))
}
