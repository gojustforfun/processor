package processor

type VerifyHandler struct {
	in          chan *Envelope
	out         chan *Envelope
	application Verifier
}

type Verifier interface {
	Verify(AddressInput) AddressOutput
}

func NewVerifierHandler(in, out chan *Envelope, application Verifier) *VerifyHandler {
	return &VerifyHandler{
		in:          in,
		out:         out,
		application: application,
	}
}

func (v *VerifyHandler) Handle() {
	received := <-v.in

	v.application.Verify(received.Input)

	v.out <- received
}
