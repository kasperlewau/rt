// package rt provides RoundTripper middleware for an *http.Client
package rt

import "net/http"

// Tripper wraps a RoundTripper and returns a new
type Tripper func(http.RoundTripper) http.RoundTripper

// TripperFunc does what an http.HandlerFunc does
type TripperFunc func(*http.Request) (*http.Response, error)

func (f TripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

// Chain is a list of Trippers
// The chain is an effective no-op. Users must call the Wrap method
// on an *http.Client to utilize the Chain
type Chain []Tripper

// New constructs a Chain of Trippers
func New(trippers ...Tripper) Chain {
	c := make([]Tripper, len(trippers))
	for i := 0; i < len(trippers); i++ {
		c[i] = trippers[i]
	}
	return c
}

// Wrap returns a clone of the given *http.Client with its Transport
// set to the chain of Trippers in reverse order
func (c Chain) Wrap(client *http.Client) *http.Client {
	if len(c) == 0 {
		return client
	}

	tr := client.Transport
	if tr == nil {
		tr = http.DefaultTransport
	}

	for i := len(c) - 1; i >= 0; i-- {
		tr = c[i](tr)
	}

	clone := *client
	clone.Transport = tr
	return &clone
}
