package rt

import "net/http"

// UserAgent returns a Tripper that sets the User-Agent header of any
// outgoing request
func UserAgent(ua string) Tripper {
	return func(next http.RoundTripper) http.RoundTripper {
		return TripperFunc(func(r *http.Request) (*http.Response, error) {
			r.Header.Set("User-Agent", ua)
			return next.RoundTrip(r)
		})
	}
}
