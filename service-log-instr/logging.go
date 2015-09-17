package main

import (
    "golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/endpoint"
)

// Transport Logging
func loggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}
