package processor_test

import (
	"strings"
	"testing"

	"github.com/gojustforfun/processor"
	"github.com/stretchr/testify/suite"
)

func TestVerifyHandler(t *testing.T) {
	suite.Run(t, new(VerifyHandlerTestSuite))
}

type VerifyHandlerTestSuite struct {
	suite.Suite

	in          chan *processor.Envelope
	out         chan *processor.Envelope
	application *FakeVerifier
	handler     *processor.VerifyHandler
}

func (s *VerifyHandlerTestSuite) SetupTest() {
	s.in = make(chan *processor.Envelope, 10)
	s.out = make(chan *processor.Envelope, 10)
	s.application = NewFakeVerifier()
	s.handler = processor.NewVerifierHandler(s.in, s.out, s.application)
}

func (s *VerifyHandlerTestSuite) TestVerifierReceivesInput() {

	envelope := s.enqueueEnvelope("street")
	close(s.in)

	s.handler.Handle()

	s.Same(envelope, <-s.out)
	s.Equal("STREET", envelope.Output.DeliveryLine1)
}

func (s *VerifyHandlerTestSuite) enqueueEnvelope(inputStreet string) *processor.Envelope {
	envelope := &processor.Envelope{
		Input: processor.AddressInput{
			Street1: inputStreet,
		},
	}
	s.in <- envelope
	return envelope
}

func (s *VerifyHandlerTestSuite) TestInputQueueDrained() {

	envelope1 := s.enqueueEnvelope("41")
	envelope2 := s.enqueueEnvelope("42")
	envelope3 := s.enqueueEnvelope("43")
	close(s.in)

	s.handler.Handle()

	s.Same(envelope1, <-s.out)
	s.Same(envelope2, <-s.out)
	s.Same(envelope3, <-s.out)
}

type FakeVerifier struct {
	input processor.AddressInput
}

func NewFakeVerifier() *FakeVerifier {
	return &FakeVerifier{}
}

func (f *FakeVerifier) Verify(i processor.AddressInput) processor.AddressOutput {
	f.input = i
	return processor.AddressOutput{DeliveryLine1: strings.ToUpper(i.Street1)}
}
