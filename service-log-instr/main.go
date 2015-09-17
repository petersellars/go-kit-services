package main

import (
	"net/http"
	"os"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	ctx := context.Background()
    logger := log.NewLogfmtLogger(os.Stderr)
	
    var svc StringService	
	svc = stringService{}

	// Transport Logging
	var uppercase endpoint.Endpoint
	uppercase = makeUppercaseEndpoint(svc)
	uppercase = loggingMiddleware(log.NewContext(logger).With("method", "uppercase")) (uppercase)

	var count endpoint.Endpoint
	count = makeCountEndpoint(svc)
	count = loggingMiddleware(log.NewContext(logger).With("method", "count")) (count)

	uppercaseHandler := httptransport.NewServer(
		ctx,
		// makeUppercaseEndpoint(svc),
		uppercase,
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		ctx,
		count,
		// makeCountEndpoint(svc)
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	_ = logger.Log("err", http.ListenAndServe(":8080", nil))
}
