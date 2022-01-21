package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifierReceivesInput(t *testing.T) {
	in := make(chan interface{}, 10)
	out := make(chan interface{}, 10)
	verifier := NewVerifier(in, out)

	expected := 1
	in <- expected
	close(in)

	verifier.Listen()

	assert.Equal(t, expected, <-out)

}
