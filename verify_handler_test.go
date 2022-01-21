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

	in          chan *Envelope
	out         chan *Envelope
	application *FakeVerifier
	handler     *VerifyHandler
}

func (s *VerifyHandlerTestSuite) SetupSuite() {
	s.in = make(chan *Envelope, 10)
	s.out = make(chan *Envelope, 10)
	s.application = NewFakeVerifier()
	s.handler = NewVerifierHandler(s.in, s.out, s.application)
}

func (s *VerifyHandlerTestSuite) TestVerifierReceivesInput() {
	envelope := &Envelope{}
	s.in <- envelope
	close(s.in)

	s.handler.Handle()

	s.Same(envelope, <-s.out)
	s.Same(envelope, s.application.input)
}

type FakeVerifier struct {
	input *Envelope
}

func NewFakeVerifier() *FakeVerifier {
	return &FakeVerifier{}
}

func (f *FakeVerifier) Verify(i *Envelope) {
	f.input = i
}
