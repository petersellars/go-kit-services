package main

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

// Application Logging Middleware
type appLoggingMiddleware struct {
	logger log.Logger
	StringService
}

func (almw appLoggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		almw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = almw.StringService.Uppercase(s)
	return
}

func (almw appLoggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		almw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = almw.StringService.Count(s)
	return
}

// Transport Logging Middleware
func loggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}
