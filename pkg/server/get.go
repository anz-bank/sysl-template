package server

import (
	"context"

	"github.com/anz-bank/sysl-template/gen/simple"
)

func Get(ctx context.Context, req *simple.GetRequest, client simple.GetClient) (*simple.Welcome, error) {
	return &simple.Welcome{Content: "Hello, World"}, nil
}
