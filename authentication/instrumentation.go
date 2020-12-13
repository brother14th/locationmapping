package authentication

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingAuthMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func (mw instrumentingAuthMiddleware) Authenticate(userName string, password string) (token string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "authenticate", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	token, err = mw.Service.Authenticate(userName, password)
	return
}

// NewInstrumentingAuthMiddleware creates new instrumenting Service.
func NewInstrumentingAuthMiddleware(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingAuthMiddleware{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}
