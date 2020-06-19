package tracing

import (
	"fmt"
	"net/http"
	"time"
)

type h struct {
	next http.RoundTripper
}

func (h *h) RoundTrip(request *http.Request) (*http.Response, error) {
	start := time.Now().UnixNano()
	defer func() {
		d := time.Now().UnixNano() - start
		fmt.Printf("duration: %v\n", time.Duration(d))
	}()
	return h.next.RoundTrip(request)
}

// NewClient 创建统计数据的http client
func NewClient(r http.RoundTripper) *http.Client {
	if r == nil {
		//r = http.DefaultTransport
		r = &http.Transport{}
	}

	// metrics
	h := &h{next: r}
	c := &http.Client{
		Transport: h,
	}
	return c
}
