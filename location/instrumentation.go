package location

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/brother14th/locationmapping/db"

)

type instrumentingLocationReportMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func (mw instrumentingLocationReportMiddleware) GetLocationReport(location string) (locationReport db.LocationReport, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetLocationReport", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	locationReport, err = mw.Service.GetLocationReport(location)
	return
}

// NewInstrumentingLocationReportMiddleware creates new instrumenting middleware.
func NewInstrumentingLocationReportMiddleware(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingLocationReportMiddleware{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}
