package handlers

import (
	"context"
	"net/http"

	"github.com/anz-bank/sysl-go/common"
)

//  return a string ptr
func newString(s string) *string {
	return &s
}

// setJSONResponseContentType sets the content type of the response to the appropriate application/json value.
func setJSONResponseContentType(ctx context.Context) {
	headers := common.RequestHeaderFromContext(ctx)
	headers["Content-Type"] = []string{"application/json; charset=utf-8"}
}

// create a new request context
func newRequestContext() context.Context {
	headers := http.Header{}
	ctx := context.Background()
	ctx = common.RequestHeaderToContext(ctx, headers)
	return ctx
}
