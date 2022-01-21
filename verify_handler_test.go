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

	in          chan interface{}
	out         chan interface{}
	application *FakeVerifier
	handler     *VerifyHandler
}

func (s *VerifyHandlerTestSuite) SetupSuite() {
	s.in = make(chan interface{}, 10)
	s.out = make(chan interface{}, 10)
	s.application = NewFakeVerifier()
	s.handler = NewVerifierHandler(s.in, s.out, s.application)
}

func (s *VerifyHandlerTestSuite) TestVerifierReceivesInput() {
	s.in <- 2
	close(s.in)

	s.handler.Handle()

	s.Equal(2, <-s.out)
	s.Equal(2, s.application.input)
}

type FakeVerifier struct {
	input interface{}
}

func NewFakeVerifier() *FakeVerifier {
	return &FakeVerifier{}
}

func (f *FakeVerifier) Verify(i interface{}) {
	f.input = i
}