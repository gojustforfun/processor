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

	return s.translateCandidate(output)
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
	if response != nil {
		response.Body.Close()
		json.NewDecoder(response.Body).Decode(&output)
	}
	return output
}

func (s *SmartyVerifier) translateCandidate(candidates []Candidate) AddressOutput {
	if len(candidates) == 0 {
		return AddressOutput{Status: "Invalid API Response"}
	}
	candidate := candidates[0]

	return AddressOutput{
		DeliveryLine1: candidate.DeliveryLine1,
		LastLine:      candidate.LastLine,
		City:          candidate.Components.City,
		State:         candidate.Components.State,
		ZIPCode:       candidate.Components.ZIPCode,
	}
}

type Candidate struct {
	DeliveryLine1 string `json:"delivery_line_1"`
	LastLine      string `json:"last_line"`
	Components    struct {
		City    string `json:"city_name"`
		State   string `json:"state_abbreviation"`
		ZIPCode string `json:"zipcode"`
	} `json:"components"`
}
