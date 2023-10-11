package tuple

import (
	"andrecastelo/raytracer/internal/compare"
	"math"
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

	suite.Equal(newPoint.x, 1.0)
	suite.Equal(newPoint.y, 1.0)
	suite.Equal(newPoint.z, 6.0)
	suite.Equal(newPoint.w, 1)
}

func (suite *TupleSuite) TestSubtractingTwoPoints() {
	pointA := Point(3, 2, 1)
	pointB := Point(5, 6, 7)

	vector := pointA.Subtract(pointB)

	suite.Equal(vector.x, -2.0)
	suite.Equal(vector.y, -4.0)
	suite.Equal(vector.z, -6.0)
	suite.Equal(vector.w, 0)
}

func (suite *TupleSuite) TestSubtractingAVectorFromAPoint() {
	point := Point(3, 2, 1)
	vector := Vector(5, 6, 7)

	newVector := point.Subtract(vector)

	suite.Equal(newVector.x, -2.0)
	suite.Equal(newVector.y, -4.0)
	suite.Equal(newVector.z, -6.0)
	suite.Equal(newVector.w, 1)
}

func (suite *TupleSuite) TestSubtractingTwoVectors() {
	vectorA := Vector(3, 2, 1)
	vectorB := Vector(5, 6, 7)

	newVector := vectorA.Subtract(vectorB)

	suite.Equal(newVector.x, -2.0)
	suite.Equal(newVector.y, -4.0)
	suite.Equal(newVector.z, -6.0)
	suite.Equal(newVector.w, 0)
}

func (suite *TupleSuite) TestNegatingATuple() {
	tuple := MakeTuple(1, -2, 3, -4)

	newTuple := tuple.Negate()

	suite.Equal(newTuple.x, -1.0)
	suite.Equal(newTuple.y, 2.0)
	suite.Equal(newTuple.z, -3.0)
	suite.Equal(newTuple.w, 4)
}

func (suite *TupleSuite) TestMultiplyingTupleByScalar() {
	tuple := MakeTuple(1, -2, 3, -4)

	newTuple := tuple.Multiply(3.5)

	suite.Equal(newTuple.x, 3.5)
	suite.Equal(newTuple.y, -7.0)
	suite.Equal(newTuple.z, 10.5)
	suite.Equal(newTuple.w, -14)
}

func (suite *TupleSuite) TestMultiplyingTupleByFraction() {
	tuple := MakeTuple(1, -2, 3, -4)

	newTuple := tuple.Multiply(0.5)

	suite.Equal(newTuple.x, 0.5)
	suite.Equal(newTuple.y, -1.0)
	suite.Equal(newTuple.z, 1.5)
	suite.Equal(newTuple.w, -2)
}

func (suite *TupleSuite) TestDividingTupleByScalar() {
	tuple := MakeTuple(1, -2, 3, -4)

	newTuple := tuple.Divide(2)

	suite.Equal(newTuple.x, 0.5)
	suite.Equal(newTuple.y, -1.0)
	suite.Equal(newTuple.z, 1.5)
	suite.Equal(newTuple.w, -2)
}

func (suite *TupleSuite) TestVectorMagnitude() {
	suite.Equal(Vector(1, 0, 0).Magnitude(), 1.0)
	suite.Equal(Vector(0, 1, 0).Magnitude(), 1.0)
	suite.Equal(Vector(0, 0, 1).Magnitude(), 1.0)

	expectedMagnitude := math.Sqrt(14)
	suite.True(compare.Equal(Vector(1, 2, 3).Magnitude(), expectedMagnitude))
	suite.True(compare.Equal(Vector(-1, -2, -3).Magnitude(), expectedMagnitude))
}

func TestTupleSuite(t *testing.T) {
	suite.Run(t, new(TupleSuite))
}
