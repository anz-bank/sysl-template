package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/anz-bank/sysl-go/common"
)

// badRequestError returns an error that represents an HTTP Bad Request (400) with the given message.
func badRequestError(message string) common.CustomError {
	return map[string]string{"name": "BadRequestError", "http_code": "1001", "http_message": fmt.Sprintf("Bad request: %s", message), "http_status": "400"}
}

// setJSONResponseContentType sets the content type of the response to the appropriate application/json value.
func setJSONResponseContentType(ctx context.Context) {
	headers, _ := common.RespHeaderAndStatusFromContext(ctx)
	headers["Content-Type"] = []string{"application/json; charset=utf-8"}
}

func newInt(i int64) *int64 {
	return &i
}

//  return a string ptr
func newString(s string) *string {
	return &s
}

// create a new request context
func newRequestContext() context.Context {
	headers := http.Header{}
	ctx := context.Background()
	ctx = common.RequestHeaderToContext(ctx, headers)
	return ctx
}
