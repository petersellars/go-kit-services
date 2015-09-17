package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentationMiddleware struct {
	requestCount	metrics.Counter
	requestLatency	metrics.TimeHistogram
	countResult		metrics.Histogram
	StringService
}

func (imw instrumentationMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "uppercase"}
		errorField := metrics.Field{Key: "error", Value: fmt.Sprintf("%v", err)}
		imw.requestCount.With(methodField).With(errorField).Add(1)
		imw.requestLatency.With(methodField).With(errorField).Observe(time.Since(begin))
	}(time.Now())

	output, err = imw.StringService.Uppercase(s)
	return
}

func (imw instrumentationMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "count"}
		errorField := metrics.Field{Key: "error", Value: fmt.Sprintf("%v", error(nil))}
		imw.requestCount.With(methodField).With(errorField).Add(1)
		imw.requestLatency.With(methodField).With(errorField).Observe(time.Since(begin))
		imw.countResult.Observe(int64(n))
	}(time.Now())

	n = imw.StringService.Count(s)
	return
}
