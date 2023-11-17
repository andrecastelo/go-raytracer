package tuple

import (
	"andrecastelo/raytracer/internal/compare"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TupleSuite struct {
	suite.Suite
}

func (suite *TupleSuite) EqualTuples(t1 *Tuple, t2 *Tuple) {
	suite.True(compare.Equal(t1.X, t2.X), fmt.Sprintf("%.12f is not equal to %.12f", t1.X, t2.X))
	suite.True(compare.Equal(t1.Y, t2.Y), fmt.Sprintf("%.12f is not equal to %.12f", t1.Y, t2.Y))
	suite.True(compare.Equal(t1.Z, t2.Z), fmt.Sprintf("%.12f is not equal to %.12f", t1.Z, t2.Z))
	suite.Equal(t1.W, t2.W)
}

func (suite *TupleSuite) TestPointIsCreatedCorrectly() {
	point := Point(1, 2, 3)

	suite.Equal(point.W, 1)
	suite.Equal(point.X, 1.0)
	suite.Equal(point.Y, 2.0)
	suite.Equal(point.Z, 3.0)
}

func (suite *TupleSuite) TestVectorIsCreatedCorrectly() {
	vector := Vector(4, 5, 6)

	suite.Equal(vector.W, 0)
	suite.Equal(vector.X, 4.0)
	suite.Equal(vector.Y, 5.0)
	suite.Equal(vector.Z, 6.0)
}

func (suite *TupleSuite) TestTupleAddition() {
	point := Point(3, -2, 5)
	vector := Vector(-2, 3, 1)

	newPoint := point.Add(vector)

	suite.Equal(newPoint.X, 1.0)
	suite.Equal(newPoint.Y, 1.0)
	suite.Equal(newPoint.Z, 6.0)
	suite.Equal(newPoint.W, 1)
}

func (suite *TupleSuite) TestSubtractingTwoPoints() {
	pointA := Point(3, 2, 1)
	pointB := Point(5, 6, 7)

	vector := pointA.Subtract(pointB)

	suite.Equal(vector.X, -2.0)
	suite.Equal(vector.Y, -4.0)
	suite.Equal(vector.Z, -6.0)
	suite.Equal(vector.W, 0)
}

func (suite *TupleSuite) TestSubtractingAVectorFromAPoint() {
	point := Point(3, 2, 1)
	vector := Vector(5, 6, 7)

	newVector := point.Subtract(vector)

	suite.Equal(newVector.X, -2.0)
	suite.Equal(newVector.Y, -4.0)
	suite.Equal(newVector.Z, -6.0)
	suite.Equal(newVector.W, 1)
}

func (suite *TupleSuite) TestSubtractingTwoVectors() {
	vectorA := Vector(3, 2, 1)
	vectorB := Vector(5, 6, 7)

	newVector := vectorA.Subtract(vectorB)

	suite.Equal(newVector.X, -2.0)
	suite.Equal(newVector.Y, -4.0)
	suite.Equal(newVector.Z, -6.0)
	suite.Equal(newVector.W, 0)
}

func (suite *TupleSuite) TestNegatingATuple() {
	tuple := MakeTuple(1, -2, 3, -4)

	newTuple := tuple.Negate()

	suite.Equal(newTuple.X, -1.0)
	suite.Equal(newTuple.Y, 2.0)
	suite.Equal(newTuple.Z, -3.0)
	suite.Equal(newTuple.W, 4)
}

func (suite *TupleSuite) TestMultiplyingTupleByScalar() {
	tuple := MakeTuple(1, -2, 3, -4)

	newTuple := tuple.Multiply(3.5)

	suite.Equal(newTuple.X, 3.5)
	suite.Equal(newTuple.Y, -7.0)
	suite.Equal(newTuple.Z, 10.5)
	suite.Equal(newTuple.W, -14)
}

func (suite *TupleSuite) TestMultiplyingTupleByFraction() {
	tuple := MakeTuple(1, -2, 3, -4)

	newTuple := tuple.Multiply(0.5)

	suite.Equal(newTuple.X, 0.5)
	suite.Equal(newTuple.Y, -1.0)
	suite.Equal(newTuple.Z, 1.5)
	suite.Equal(newTuple.W, -2)
}

func (suite *TupleSuite) TestDividingTupleByScalar() {
	tuple := MakeTuple(1, -2, 3, -4)

	newTuple := tuple.Divide(2)

	suite.Equal(newTuple.X, 0.5)
	suite.Equal(newTuple.Y, -1.0)
	suite.Equal(newTuple.Z, 1.5)
	suite.Equal(newTuple.W, -2)
}

func (suite *TupleSuite) TestVectorMagnitude() {
	suite.Equal(Vector(1, 0, 0).Magnitude(), 1.0)
	suite.Equal(Vector(0, 1, 0).Magnitude(), 1.0)
	suite.Equal(Vector(0, 0, 1).Magnitude(), 1.0)

	expectedMagnitude := math.Sqrt(14)
	suite.True(compare.Equal(Vector(1, 2, 3).Magnitude(), expectedMagnitude))
	suite.True(compare.Equal(Vector(-1, -2, -3).Magnitude(), expectedMagnitude))
}

func (suite *TupleSuite) TestVectorNormalization() {
	suite.EqualTuples(Vector(5, 0, 0).Normalize(), Vector(1, 0, 0))
	suite.EqualTuples(Vector(0, 5, 0).Normalize(), Vector(0, 1, 0))
	suite.EqualTuples(Vector(0, 0, 5).Normalize(), Vector(0, 0, 1))

	normalizedVector := Vector(1, 2, 3).Normalize()

	suite.EqualTuples(normalizedVector, Vector(0.267261241912, 0.534522483825, 0.801783725737))
}

func (suite *TupleSuite) TestVectorDotProduct() {
	vectorA := Vector(1, 2, 3)
	vectorB := Vector(2, 3, 4)

	dotProduct := vectorA.Dot(vectorB)

	suite.Equal(dotProduct, 20.0)

	suite.Equal(Vector(1, 0, 0).Dot(Vector(1, 0, 0)), 1.0)
	suite.Equal(Vector(1, 0, 0).Dot(Vector(-1, 0, 0)), -1.0)

	suite.Equal(Vector(1, 0, 0).Dot(Vector(0, 1, 0)), 0.0)
	suite.Equal(Vector(1, 0, 0).Dot(Vector(0, -1, 0)), 0.0)
}

func (suite *TupleSuite) TestVectorCrossProduct() {
	vectorA := Vector(1, 2, 3)
	vectorB := Vector(2, 3, 4)

	suite.EqualTuples(vectorA.Cross(vectorB), Vector(-1, 2, -1))
	suite.EqualTuples(vectorB.Cross(vectorA), Vector(1, -2, 1))

	x := Vector(1, 0, 0)
	y := Vector(0, 1, 0)
	z := Vector(0, 0, 1)

	suite.EqualTuples(x.Cross(y), z)
	suite.EqualTuples(y.Cross(z), x)
	suite.EqualTuples(x.Cross(y), z)

	suite.EqualTuples(z.Cross(y), x.Negate())
	suite.EqualTuples(y.Cross(x), z.Negate())
	suite.EqualTuples(x.Cross(z), y.Negate())
}

func (suite *TupleSuite) TestHadamardProduct() {
	colorA := Color(1.0, 0.2, 0.4)
	colorB := Color(0.9, 1.0, 0.1)

	blendedColor := colorA.Hadamard(colorB)

	suite.EqualTuples(blendedColor, Color(0.9, 0.2, 0.04))
}

func (suite *TupleSuite) TestArray() {
	tupleA := MakeTuple(13, 12, 11, 10)

	suite.Equal(tupleA.Array(), [...]float64{13.0, 12.0, 11.0, 10.0})
}

func TestTupleSuite(t *testing.T) {
	suite.Run(t, new(TupleSuite))
}
