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

	query := url.Values{}
	query.Set("street", input.Street1)

	request, _ := http.NewRequest(http.MethodGet, "/street-address?" + query.Encode(), nil)
	s.client.Do(request)
	return AddressOutput{}
}
