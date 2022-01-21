package processor

type VerifyHandler struct {
	in  chan interface{}
	out chan interface{}
}

func NewVerifier(in, out chan interface{}) *VerifyHandler {
	return &VerifyHandler{
		in:  in,
		out: out,
	}
}

func (v *VerifyHandler) Listen() {
	v.out <- 1
}
