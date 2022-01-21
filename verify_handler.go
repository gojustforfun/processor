package processor

type Verifier struct {
	in  chan interface{}
	out chan interface{}
}

func NewVerifier(in, out chan interface{}) *Verifier {
	return &Verifier{
		in:  in,
		out: out,
	}
}

func (v *Verifier) Listen() {
	v.out <- 1
}
