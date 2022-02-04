package processor_test

import (
	"bytes"
	"errors"
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
	s.client.ConfigureResponseInfo(`[{}]`, http.StatusOK, nil)

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

func (s *VerifierTestSuite) TestResponseParsed() {
	s.client.ConfigureResponseInfo(rawJSONInput, http.StatusOK, nil)

	addressOutput := s.verifier.Verify(processor.AddressInput{})

	s.Equal("1 Santa Claus Ln", addressOutput.DeliveryLine1)
	s.Equal("North Pole AK 99705-9901", addressOutput.LastLine)
	s.Equal("North Pole", addressOutput.City)
	s.Equal("AK", addressOutput.State)
	s.Equal("99705", addressOutput.ZIPCode)
}

const rawJSONInput = `[
	{
		"delivery_line_1":"1 Santa Claus Ln",
		"last_line":"North Pole AK 99705-9901",
		"components":{
			"city_name": "North Pole",
			"state_abbreviation": "AK",
			"zipcode": "99705"
		}
	}
]`

func (s *VerifierTestSuite) TestMalformedJSONHandled() {
	s.client.ConfigureResponseInfo(malformedRawJSONOutput, http.StatusOK, nil)
	result := s.verifier.Verify(processor.AddressInput{})
	s.Equal("Invalid API Response", result.Status)
}

const malformedRawJSONOutput = `I am not JSON`

func (s *VerifierTestSuite) TestHTTPErrorHandled() {
	s.client.ConfigureResponseInfo("", 0, errors.New("Gophers!"))
	result := s.verifier.Verify(processor.AddressInput{})
	s.Equal("Invalid API Response", result.Status)
}

func (s *VerifierTestSuite) TestHTTPResponseBodyClosed() {
	s.client.ConfigureResponseInfo(rawJSONInput, http.StatusOK, nil)
	s.verifier.Verify(processor.AddressInput{})
	body, ok := s.client.response.Body.(*SpyBuffer)
	s.Equal(true, ok)
	s.Equal(1, body.Closed)
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
	if err == nil {
		f.response = &http.Response{
			Body:       NewSpyBuffer(responseText),
			StatusCode: statusCode,
		}
	}
	f.err = err
}

func (f *FakeHTTPClient) Do(request *http.Request) (*http.Response, error) {
	f.request = request
	return f.response, f.err
}

func NewSpyBuffer(s string) *SpyBuffer {
	return &SpyBuffer{
		Buffer: *bytes.NewBufferString(s),
	}
}

type SpyBuffer struct {
	bytes.Buffer
	Closed int
}

func (s *SpyBuffer) Close() error {
	s.Closed++
	return nil
}
