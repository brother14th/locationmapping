package user

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingUserMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func (mw instrumentingUserMiddleware) SetPreferredLocation(username string, location string) (status bool, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "SetPreferredLocation", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	status, err = mw.Service.SetPreferredLocation(username, location)
	return
}

func (mw instrumentingUserMiddleware) GetPreferredLocation(username string) (location string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetPreferredLocation", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	location, err = mw.Service.GetPreferredLocation(username)
	return
}

// NewInstrumentingUserMiddleware creates new instrumenting middleware.
func NewInstrumentingUserMiddleware(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingUserMiddleware{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}
