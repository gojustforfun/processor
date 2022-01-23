package processor_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestVerifierTestSuite(t *testing.T) {
	suite.Run(t, new(VerifierTestSuite))
}

type VerifierTestSuite struct {
	suite.Suite
}

func (s *VerifierTestSuite) SetupSuite() {

}

func (s *VerifierTestSuite) TestNothing() {

}
