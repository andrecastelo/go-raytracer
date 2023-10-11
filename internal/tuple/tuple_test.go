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

	suite.Equal(point.w, 0)
}

func TestTupleSuite(t *testing.T) {
	suite.Run(t, new(TupleSuite))
}
