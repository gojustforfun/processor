package processor

type VerifyHandler struct {
	in  chan interface{}
	out chan interface{}
	application interface{}
}

func NewVerifier(in, out chan interface{}, application interface{}) *VerifyHandler {
	return &VerifyHandler{
		in:  in,
		out: out,
		application: application,
	}
}

func (v *VerifyHandler) Handle() {
	v.out <- <-v.in
}
