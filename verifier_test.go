package processor_test

import (
	"net/http"
	"testing"

	"github.com/gojustforfun/processor"
	"github.com/stretchr/testify/suite"
)

func TestVerifierTestSuite(t *testing.T) {
	suite.Run(t, new(VerifierTestSuite))
}

type VerifierTestSuite struct {
	suite.Suite

	client   *FakeHTTPClient
	verifier *processor.SmartyVerifier
}

func (s *VerifierTestSuite) SetupSuite() {
	s.client = NewFakeHTTPClient()
	s.verifier = processor.NewSmartyVerifier(s.client)
}

func (s *VerifierTestSuite) TestRequestComposedProperly() {
	input := processor.AddressInput{
		Street1: "street1",
	}

	s.verifier.Verify(input)

	s.Equal("street=street1", s.client.request.URL.RawQuery)
}

func NewFakeHTTPClient() *FakeHTTPClient {
	return &FakeHTTPClient{}
}

type FakeHTTPClient struct {
	request *http.Request
}

func (f *FakeHTTPClient) Do(r *http.Request) (*http.Response, error) {
	return nil, nil
}
