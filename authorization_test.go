package rt

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthorization(t *testing.T) {
	auth := "test"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != auth {
			t.Fatalf("got unexpected User-Agent. want = %s, got = %s", auth, got)
		}
	}))
	defer srv.Close()
	chain := New(Authorization(auth))
	client := chain.Wrap(&http.Client{})
	req, err := http.NewRequest("GET", srv.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := client.Do(req); err != nil {
		t.Fatal(err)
	}
}
