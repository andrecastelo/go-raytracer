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
	suite.True(compare.Equal(t1.x, t2.x), fmt.Sprintf("%.12f is not equal to %.12f", t1.x, t2.x))
	suite.True(compare.Equal(t1.y, t2.y), fmt.Sprintf("%.12f is not equal to %.12f", t1.y, t2.y))
	suite.True(compare.Equal(t1.z, t2.z), fmt.Sprintf("%.12f is not equal to %.12f", t1.z, t2.z))
	suite.Equal(t1.w, t2.w)
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

func TestTupleSuite(t *testing.T) {
	suite.Run(t, new(TupleSuite))
}
