package processor

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestVerifyHandler(t *testing.T) {
	suite.Run(t, new(VerifyHandlerTestSuite))
}

type VerifyHandlerTestSuite struct {
	suite.Suite
	
	in      chan interface{}
	out     chan interface{}
	handler *VerifyHandler
}

func (s *VerifyHandlerTestSuite) SetupSuite() {
	s.in = make(chan interface{}, 10)
	s.out = make(chan interface{}, 10)
	s.handler = NewVerifier(s.in, s.out)
}

func (s *VerifyHandlerTestSuite) TestVerifierReceivesInput() {
	s.in <- 1
	close(s.in)

	s.handler.Listen()

	s.Equal(1, <-s.out)
}
