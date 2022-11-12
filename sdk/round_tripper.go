package sdk

import "net/http"

const authorizationHeaderName = "x-api-key"

func NewAuthHeaderRoundTripper(apiKey string, next http.RoundTripper) *AuthHeaderRoundTripper {
	return &AuthHeaderRoundTripper{apiKey: apiKey, next: next}
}

type AuthHeaderRoundTripper struct {
	apiKey string
	next   http.RoundTripper
}

func (a AuthHeaderRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Add(authorizationHeaderName, a.apiKey)
	request.Header.Add("Content-Type", "application/json")

	return a.next.RoundTrip(request)
}
