package rt

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserAgent(t *testing.T) {
	ua := "test"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("User-Agent"); got != ua {
			t.Fatalf("got unexpected User-Agent. want = %s, got = %s", ua, got)
		}
	}))
	defer srv.Close()
	chain := New(UserAgent(ua))
	client := chain.Wrap(&http.Client{})
	req, err := http.NewRequest("GET", srv.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := client.Do(req); err != nil {
		t.Fatal(err)
	}
}
