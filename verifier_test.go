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
		Street1: "Street1",
		City:    "City",
		State:   "State",
		ZIPCode: "ZIPcode",
	}

	s.verifier.Verify(input)

	s.Equal(http.MethodGet, s.client.request.Method)
	s.Equal("/street-address", s.client.request.URL.Path)
	s.EqualQueryStringValue("street", "Street1")
	s.EqualQueryStringValue("city", "City")
	s.EqualQueryStringValue("state", "State")
	s.EqualQueryStringValue("zipcode", "ZIPcode")
}

func (s *VerifierTestSuite) EqualQueryStringValue(key string, val string) {
	query := s.client.request.URL.Query()
	s.Equal(val, query.Get(key))
}

func NewFakeHTTPClient() *FakeHTTPClient {
	return &FakeHTTPClient{}
}

type FakeHTTPClient struct {
	request *http.Request
}

func (f *FakeHTTPClient) Do(request *http.Request) (*http.Response, error) {
	f.request = request
	return nil, nil
}
