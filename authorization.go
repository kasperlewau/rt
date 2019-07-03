package rt

import "net/http"

// UserAgent returns a Tripper that sets the Authorization header of any outgoing request
func Authorization(auth string) Tripper {
	return func(next http.RoundTripper) http.RoundTripper {
		return TripperFunc(func(r *http.Request) (*http.Response, error) {
			r.Header.Set("Authorization", auth)
			return next.RoundTrip(r)
		})
	}
}
