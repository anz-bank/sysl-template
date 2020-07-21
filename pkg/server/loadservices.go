package server

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-template/gen/jsonplaceholder"
	"github.com/anz-bank/sysl-template/gen/simple"
	"github.com/go-chi/chi"
)

func LoadServices(ctx context.Context) error {
	router := chi.NewRouter()

	// simpleServiceInterface is the struct which is composed of our functions we wrote in `methods.go`
	// Struct embedding is used for the Service interface (yes, not interfaces)
	simpleServiceInterface := simple.ServiceInterface{
		GetFoobarList: GetFoobarList,
		Get:           Get,
	}

	// Default callback behaviour
	genCallbacks := common.DefaultCallback()

	serviceHandler := simple.NewServiceHandler(
		genCallbacks,
		&simpleServiceInterface,
		jsonplaceholder.NewClient(http.DefaultClient, "http://jsonplaceholder.typicode.com"))

	// Service Router
	serviceRouter := simple.NewServiceRouter(genCallbacks, serviceHandler)
	serviceRouter.WireRoutes(ctx, router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	serverAddress := ":" + port

	log.Println("Starting Server on " + serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, router))
	return nil
}
