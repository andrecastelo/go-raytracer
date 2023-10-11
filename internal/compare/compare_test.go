package compare

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CompareSuite struct {
	suite.Suite
}

func (suite *CompareSuite) TestFloatComparisons() {
	suite.False(Equal(1.0, 0.0))
	suite.False(Equal(1.0, 0.9))
	suite.False(Equal(1.0, 0.99))
	suite.False(Equal(1.0, 0.999))
	suite.False(Equal(1.0, 0.9999))
	suite.False(Equal(1.0, 0.99999))
	suite.False(Equal(1.0, 0.999999))
	suite.False(Equal(1.0, 0.9999999))
	suite.False(Equal(1.0, 0.99999999))
	suite.True(Equal(1.0, 0.999999999))
	suite.True(Equal(1.0, 0.9999999999))
}

func TestCompareSuite(t *testing.T) {
	suite.Run(t, new(CompareSuite))
}
