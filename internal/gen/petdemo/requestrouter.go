// Code generated by sysl DO NOT EDIT.
package petdemo

import (
	"context"
	"net/http"

	"github.com/anz-bank/sysl-go/core"
	"github.com/anz-bank/sysl-go/handlerinitialiser"
	"github.com/go-chi/chi"
)

// Router interface for Petdemo
type Router interface {
	Route(router *chi.Mux)
}

// ServiceRouter for Petdemo API
type ServiceRouter struct {
	gc               core.RestGenCallback
	svcHandler       *ServiceHandler
	basePathFromSpec string
}

// swaggerFile is a struct to store the swagger file system
type swaggerFile struct {
	file http.FileSystem
}

// swagger will receive the embedded swagger file if it is generated by the resource application
var swagger = swaggerFile{}

// NewServiceRouter creates a new service router for Petdemo
func NewServiceRouter(gc core.RestGenCallback, svcHandler *ServiceHandler) handlerinitialiser.HandlerInitialiser {
	return &ServiceRouter{gc, svcHandler, ""}
}

// WireRoutes ...
//nolint:funlen
func (s *ServiceRouter) WireRoutes(ctx context.Context, r chi.Router) {
	r.Route(core.SelectBasePath(s.basePathFromSpec, s.gc.BasePath()), func(r chi.Router) {
		s.gc.AddMiddleware(ctx, r)
		r.Get("/random-pet-pic", s.svcHandler.GetRandomPetPicListHandler)
	})
}

// Config ...
func (s *ServiceRouter) Config() interface{} {
	return s.gc.Config()
}

// Name ...
func (s *ServiceRouter) Name() string {
	return "Petdemo"
}
