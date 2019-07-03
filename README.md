# rt

package rt provides a set of `http.RoundTripper` middleware for common tasks

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
