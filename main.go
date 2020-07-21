// main.go
package main

import (
	"context"

	"github.com/anz-bank/sysl-template/pkg/server"
)

func main() {
	server.LoadServices(context.Background())
}
