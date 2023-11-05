package clamp

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ClampSuite struct {
    suite.Suite
}

func (suite *ClampSuite) TestIntegerClamping() {
    suite.Equal(50, Clamp(-22, 100, 50))
    suite.Equal(60, Clamp(134, 60, 0))
}

func (suite *ClampSuite) TestFloatClamping() {
    suite.Equal(5.5, Clamp(-13.112, 9.8, 5.5))
    suite.Equal(10.0, Clamp(91.23, 10.0, 0.0))
}

func TestClampSuite(t *testing.T) {
    suite.Run(t, new(ClampSuite))
}
