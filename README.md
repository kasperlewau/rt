# rt [![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/kasperlewau/rt) [![Build Status](https://travis-ci.org/kasperlewau/rt.svg?branch=master)](https://travis-ci.org/kasperlewau/rt) [![Go Report Card](https://goreportcard.com/badge/github.com/kasperlewau/rt)](https://goreportcard.com/report/github.com/kasperlewau/rt)

package rt provides a set of `http.RoundTripper` middleware for common tasks

## install
```sh
go get github.com/kasperlewau/rt
```

## usage
```go
import (
	"net/http"
	"time"

	"github.com/kasperlewau/rt"
)

func main() {
	chain := rt.New(
		rt.UserAgent("my_ua"),
	)

	client := chain.Wrap(&http.Client{
		Timeout: 10 * time.Second,
	})

	req, _ := http.NewRequest("GET", "https://www.whatsmyua.info/api/v1/ua", nil)
	resp, _ := client.Do(req)
}
```

## license
MIT
