package processor

type VerifyHandler struct {
	input    chan *Envelope
	output   chan *Envelope
	verifier Verifier
}

type Verifier interface {
	Verify(AddressInput) AddressOutput
}

func NewVerifierHandler(input, output chan *Envelope, verifier Verifier) *VerifyHandler {
	return &VerifyHandler{
		input:    input,
		output:   output,
		verifier: verifier,
	}
}

func (v *VerifyHandler) Handle() {
	for envelope := range v.input {
		envelope.Output = v.verifier.Verify(envelope.Input)
		v.output <- envelope
	}
}
