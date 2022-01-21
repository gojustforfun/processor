package processor

type VerifyHandler struct {
	in          chan interface{}
	out         chan interface{}
	application Verifier
}

type Verifier interface {
	Verify(interface{})
}

func NewVerifierHandler(in, out chan interface{}, application Verifier) *VerifyHandler {
	return &VerifyHandler{
		in:          in,
		out:         out,
		application: application,
	}
}

func (v *VerifyHandler) Handle() {
	received := <-v.in

	v.application.Verify(received)
	
	v.out <- received
}