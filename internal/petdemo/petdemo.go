package petdemo

import (
	"context"
	"net/http"
	"time"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/config"
	"github.com/anz-bank/sysl-go/core"
	"github.com/anz-bank/sysl-go/handlerinitialiser"
	"github.com/anz-bank/sysl-template/internal/gen/petdemo"
	"github.com/anz-bank/sysl-template/internal/petdemo/handlers"
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus"
)

// ListenAndServe func
func ListenAndServe(cfg *ConfigContainer) error {
	// construct the mapping from endpoint to handler
	si := petdemo.ServiceInterface{
		GetRandomPetPicList: handlers.GetRandomPetPicListRead,
	}

	// construct the rest configuration (aka. gen callback)
	restConfig := restGenCallback{
		UpstreamTimeout:   cfg.GenCode.Upstream.ContextTimeout,
		DownstreamTimeout: cfg.GenCode.Downstream.(*petdemo.DownstreamConfig).ContextTimeout,
		RouterBasePath:    "/",
		UpstreamConfig:    &cfg.GenCode.Upstream,
	}

	// construct the downstream clients
	clients, err := petdemo.BuildDownstreamClients(&cfg.DefaultConfig)
	if err != nil {
		return err
	}

	// construct the service router
	serviceRouter := petdemo.BuildRestHandlerInitialiser(si, restConfig, clients)

	// start the server
	restManager := newRestManager(cfg, serviceRouter)
	return core.NewServerParams(context.Background(), "petdemo",
		core.WithPkgLogger(),
		core.WithRestManager(restManager),
		core.WithPrometheusRegistry(prometheus.NewRegistry())).Start()
}

// restManager
type restManager struct {
	config              *ConfigContainer
	enabledHandlers     []handlerinitialiser.HandlerInitialiser
	enabledGrpcHandlers []handlerinitialiser.GrpcHandlerInitialiser
}

// newRestManager create new restManager
func newRestManager(cfg *ConfigContainer, router handlerinitialiser.HandlerInitialiser) *restManager {
	return &restManager{
		config:              cfg,
		enabledHandlers:     []handlerinitialiser.HandlerInitialiser{router},
		enabledGrpcHandlers: nil}
}

// EnabledHandlers func
func (h *restManager) EnabledHandlers() []handlerinitialiser.HandlerInitialiser {
	return h.enabledHandlers
}

// LibraryConfig func
func (h *restManager) LibraryConfig() *config.LibraryConfig {
	return &h.config.Library
}

// AdminServerConfig func
func (h *restManager) AdminServerConfig() *config.CommonHTTPServerConfig {
	return &h.config.AdminServer
}

// PublicServerConfig func
func (h *restManager) PublicServerConfig() *config.CommonHTTPServerConfig {
	return &h.config.GenCode.Upstream.HTTP
}

// EnabledGrpcHandlers func
func (h *restManager) EnabledGrpcHandlers() []handlerinitialiser.GrpcHandlerInitialiser {
	return h.enabledGrpcHandlers
}

// GrpcAdminServerConfig func
func (h *restManager) GrpcAdminServerConfig() *config.CommonServerConfig {
	return nil
}

// GrpcPublicServerConfig func
func (h *restManager) GrpcPublicServerConfig() *config.CommonServerConfig {
	return &h.config.GenCode.Upstream.GRPC
}

/* - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 * rest gen callback
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - */

type restGenCallback struct {
	UpstreamTimeout   time.Duration
	DownstreamTimeout time.Duration
	RouterBasePath    string
	UpstreamConfig    interface{}
}

// DownStreamTimeoutContext func
func (c restGenCallback) DownstreamTimeoutContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, c.DownstreamTimeout)
}

// Config func
func (c restGenCallback) Config() interface{} {
	return c.UpstreamConfig
}

// MapError func
func (c restGenCallback) MapError(context.Context, error) *common.HTTPError {
	return nil
}

// AddMiddleware func
func (c restGenCallback) AddMiddleware(ctx context.Context, r chi.Router) {
	r.Use(
		common.Timeout(ctx, c.UpstreamTimeout, http.HandlerFunc(c.timeoutHandler)),
	)
}

// BasePath func
func (c restGenCallback) BasePath() string {
	if c.RouterBasePath == "" {
		return "/"
	}
	return c.RouterBasePath
}

// Timeout handler
func (c restGenCallback) timeoutHandler(w http.ResponseWriter, r *http.Request) {
	common.HandleError(r.Context(), w, common.InternalError, "timeout expired while processing response", nil,
		func(ctx context.Context, err error) *common.HTTPError {
			return nil
		})
}
