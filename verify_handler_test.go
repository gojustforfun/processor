package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifierReceivesInput(t *testing.T) {
	handler, out := setup()
	handler.Listen()
	assert.Equal(t, 1, <-out)
}

func setup() (handler *VerifyHandler, out chan interface{}) {
	in := make(chan interface{}, 10)
	out = make(chan interface{}, 10)
	handler = NewVerifier(in, out)
	
	in <- 1
	close(in)
	return 
}
