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
	envelope := &Envelope{
		Input: AddressInput{
			Street1: "42",
		},
	}
	s.application.output = AddressOutput{
		DeliveryLine1: "DeliveryLine1",
	}

	s.in <- envelope
	close(s.in)

	s.handler.Handle()

	s.Same(envelope, <-s.out)
	s.Equal("42", s.application.input.Street1)
	s.Equal("DeliveryLine1", envelope.Output.DeliveryLine1)
}

type FakeVerifier struct {
	input AddressInput
	output AddressOutput
}

func NewFakeVerifier() *FakeVerifier {
	return &FakeVerifier{}
}

func (f *FakeVerifier) Verify(i AddressInput) AddressOutput {
	f.input = i
	return f.output
}
