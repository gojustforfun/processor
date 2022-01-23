package processor

import (
	"strings"
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

	envelope := s.enqueueEnvelope("street")

	s.handler.Handle()

	s.Same(envelope, <-s.out)
	s.Equal("STREET", envelope.Output.DeliveryLine1)
}

func (s *VerifyHandlerTestSuite) enqueueEnvelope(inputStreet string) *Envelope {
	envelope := &Envelope{
		Input: AddressInput{
			Street1: inputStreet,
		},
	}
	s.in <- envelope
	return envelope
}

type FakeVerifier struct {
	input  AddressInput
}

func NewFakeVerifier() *FakeVerifier {
	return &FakeVerifier{}
}

func (f *FakeVerifier) Verify(i AddressInput) AddressOutput {
	f.input = i
	return AddressOutput{DeliveryLine1: strings.ToUpper(i.Street1)}
}
