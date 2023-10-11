package tuple

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TupleSuite struct {
	suite.Suite
}

func (suite *TupleSuite) TestPointIsCreatedCorrectly() {
	point := Point(1, 2, 3)

	suite.Equal(point.w, 1)
	suite.Equal(point.x, 1.0)
	suite.Equal(point.y, 2.0)
	suite.Equal(point.z, 3.0)
}

func (suite *TupleSuite) TestVectorIsCreatedCorrectly() {
	vector := Vector(4, 5, 6)

	suite.Equal(vector.w, 0)
	suite.Equal(vector.x, 4.0)
	suite.Equal(vector.y, 5.0)
	suite.Equal(vector.z, 6.0)
}

func (suite *TupleSuite) TestTupleAddition() {
	point := Point(3, -2, 5)
	vector := Vector(-2, 3, 1)

	newPoint := point.Add(vector)

	suite.Equal(newPoint.w, 1)
	suite.Equal(newPoint.x, 1.0)
	suite.Equal(newPoint.y, 1.0)
	suite.Equal(newPoint.z, 6.0)
}

func (suite *TupleSuite) TestSubtractingTwoPoints() {
	pointA := Point(3, 2, 1)
	pointB := Point(5, 6, 7)

	vector := pointA.Subtract(pointB)

	suite.Equal(vector.w, 0)
	suite.Equal(vector.x, -2.0)
	suite.Equal(vector.y, -4.0)
	suite.Equal(vector.z, -6.0)
}

func (suite *TupleSuite) TestSubtractingAVectorFromAPoint() {
	point := Point(3, 2, 1)
	vector := Vector(5, 6, 7)

	newVector := point.Subtract(vector)

	suite.Equal(newVector.w, 1)
	suite.Equal(newVector.x, -2.0)
	suite.Equal(newVector.y, -4.0)
	suite.Equal(newVector.z, -6.0)
}

func TestTupleSuite(t *testing.T) {
	suite.Run(t, new(TupleSuite))
}
