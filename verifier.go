package processor

import (
	"encoding/json"
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

	request := s.buildRequest(input)

	response, _ := s.client.Do(request)

	output := s.decodeResponse(response)

	return s.translateCandidate(output[0])
}

func (s *SmartyVerifier) buildRequest(input AddressInput) *http.Request {
	query := url.Values{}
	query.Set("street", input.Street1)
	query.Set("city", input.City)
	query.Set("state", input.State)
	query.Set("zipcode", input.ZIPCode)

	request, _ := http.NewRequest(http.MethodGet, "/street-address?"+query.Encode(), nil)
	return request
}

func (s *SmartyVerifier) decodeResponse(response *http.Response) (output []Candidate) {
	json.NewDecoder(response.Body).Decode(&output)
	return output
}

func (s *SmartyVerifier) translateCandidate(candidate Candidate) AddressOutput {
	return AddressOutput{
		DeliveryLine1: candidate.DeliveryLine1,
	}
}

type Candidate struct {
	DeliveryLine1 string `json:"delivery_line_1"`
}
