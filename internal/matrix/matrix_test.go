package matrix

import (
	"andrecastelo/raytracer/internal/compare"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MatrixSuite struct {
	suite.Suite
}

func (suite *MatrixSuite) MatrixesAreEqual(a [][]float64, b [][]float64) {
	for i := range a {
		for j := range a[i] {
			suite.True(compare.Equal(a[i][j], b[i][j]))
		}
	}
}

func (suite *MatrixSuite) TestMatrixEquality() {
	a := [][]float64{
		{1.0, 2.0, 3.0, 4.0},
		{5.0, 6.0, 7.0, 8.0},
		{9.0, 10.0, 11.0, 12.0},
		{13.0, 14.0, 15.0, 16.0},
	}

	b := [][]float64{
		{1.0, 2.0, 3.0, 4.0},
		{5.0, 6.0, 7.0, 8.0},
		{9.0, 10.0, 11.0, 12.0},
		{13.0, 14.0, 15.0, 16.0},
	}

	c := [][]float64{
		{0.0, 2.0, 3.0, 4.0},
		{5.0, 6.0, 7.0, 8.0},
		{9.0, 10.0, 11.0, 12.0},
		{13.0, 14.0, 15.0, 16.0},
	}

	suite.True(Equal(a, b))
	suite.False(Equal(a, c))
}

func (suite *MatrixSuite) TestCanCreateMatrix() {
	twoByTwo := [][]float64{
		{-3.0, 5.0},
		{1.0, -2.0},
	}

	twoByTwoMatrix, err := MakeMatrix(twoByTwo)

	suite.NotNil(twoByTwoMatrix)
	suite.Nil(err)
	suite.MatrixesAreEqual(twoByTwo, twoByTwoMatrix)

	threeByThree := [][]float64{
		{-3.0, 5.0, 0.0},
		{1.0, -2.0, -7.0},
		{0.0, 1.0, 1.0},
	}

	threeByThreeMatrix, err := MakeMatrix(threeByThree)

	suite.NotNil(threeByThreeMatrix)
	suite.Nil(err)
	suite.MatrixesAreEqual(threeByThree, threeByThreeMatrix)

	fourByFour := [][]float64{
		{-3.0, 5.0, 0.0, 1.0},
		{1.0, -2.0, -7.0, 8.5},
		{0.0, 1.0, 1.0, 16.0},
		{13.0, 14.0, 15.0, 16.0},
	}

	fourByFourMatrix, err := MakeMatrix(fourByFour)

	suite.NotNil(fourByFourMatrix)
	suite.Nil(err)
	suite.MatrixesAreEqual(fourByFour, fourByFourMatrix)
}

func (suite *MatrixSuite) TestMatrixCreationFailure() {
	invalid := [][]float64{
		{1.0, 2.0, 3.0, 4.0},
		{5.0, 6.0, 7.0},
		{8.0, 9.0, 10.0, 11.0},
		{12.0, 13.0, 14.0, 15.0},
	}

	invalidMatrix, err := MakeMatrix(invalid)

	suite.Nil(invalidMatrix)
	suite.NotNil(err)
}

func TestMatrixSuite(t *testing.T) {
	suite.Run(t, new(MatrixSuite))
}
