package rt

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeaders(t *testing.T) {
	auth := []string{"a", "b"}
	ua := []string{"c", "d"}
	h := map[string][]string{
		"Authorization": auth,
		"User-Agent":    ua,
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Header.Get("User-Agent"), ua[0]; got != want {
			t.Fatalf("got unexpected User-Agent. want = %s, got = %s", ua, got)
		}
		if got, want := r.Header.Get("Authorization"), auth[0]; got != want {
			t.Fatalf("got unexpected Authorization. want = %s, got = %s", auth, got)
		}
	}))
	defer srv.Close()
	chain := New(Headers(h))
	client := chain.Wrap(&http.Client{})
	req, err := http.NewRequest("GET", srv.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := client.Do(req); err != nil {
		t.Fatal(err)
	}
}
