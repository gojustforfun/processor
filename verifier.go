package processor

import (
	"net/http"
	"net/url"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type SmartyVerifier struct {
	client HTTPClient
}

func NewSmartyVerifier(client HTTPClient) *SmartyVerifier {
	return &SmartyVerifier{
		client: client,
	}
}

func (s *SmartyVerifier) Verify(input AddressInput) AddressOutput {

	values := url.Values{}
	values.Set("street", input.Street1)

	s.client.Do(&http.Request{URL: &url.URL{
		RawQuery: values.Encode(),
	}})
	return AddressOutput{}
}