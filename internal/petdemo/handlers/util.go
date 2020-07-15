package handlers

import (
	"context"
	"net/http"

	"github.com/anz-bank/sysl-go/common"
)

// setJSONResponseContentType sets the content type of the response to the appropriate application/json value.
func setJSONResponseContentType(ctx context.Context) {
	headers := common.RequestHeaderFromContext(ctx)
	headers["Content-Type"] = []string{"application/json; charset=utf-8"}
}

//  newString creates a string ptr
func newString(s string) *string {
	return &s
}

//  newFloat64 creates a float ptr
func newFloat64(f float64) *float64 {
	return &f
}

// newRequestContext creates a new request context
func newRequestContext() context.Context {
	headers := http.Header{}
	ctx := context.Background()
	ctx = common.RequestHeaderToContext(ctx, headers)
	return ctx
}
