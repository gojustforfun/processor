package processor_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestVerifierTestSuite(t *testing.T) {
	suite.Run(t, new(VerifierTestSuite))
}

type VerifierTestSuite struct {
	suite.Suite

	client *FakeHTTPClient
}

func (s *VerifierTestSuite) SetupSuite() {
	s.client = NewFakeHTTPClient()
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
