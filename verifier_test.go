package processor_test

import (
	"bytes"
	"io"
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
	s.EqualQueryStringValue("street", input.Street1)
	s.EqualQueryStringValue("city", input.City)
	s.EqualQueryStringValue("state", input.State)
	s.EqualQueryStringValue("zipcode", input.ZIPCode)
}

func (s *VerifierTestSuite) EqualQueryStringValue(key string, val string) {
	query := s.client.request.URL.Query()
	s.Equal(val, query.Get(key))
}

func NewFakeHTTPClient() *FakeHTTPClient {
	return &FakeHTTPClient{}
}

type FakeHTTPClient struct {
	request  *http.Request
	response *http.Response
	err      error
}

func (f *FakeHTTPClient) ConfigureResponseInfo(responseText string, statusCode int, err error) {
	f.response = &http.Response{
		Body:       io.NopCloser(bytes.NewBufferString(responseText)),
		StatusCode: statusCode,
	}
	f.err = err
}

func (f *FakeHTTPClient) Do(request *http.Request) (*http.Response, error) {
	f.request = request
	return f.response, f.err
}
