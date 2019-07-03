package rt

import "net/http"

// Headers returns a Tripper that sets the given headers of any outgoing request
func Headers(h http.Header) Tripper {
	return func(next http.RoundTripper) http.RoundTripper {
		return TripperFunc(func(r *http.Request) (*http.Response, error) {
			for k, vv := range h {
				for _, v := range vv {
					r.Header.Add(k, v)
				}
			}
			return next.RoundTrip(r)
		})
	}
}
