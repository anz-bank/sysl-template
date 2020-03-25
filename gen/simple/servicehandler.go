// Code generated by sysl DO NOT EDIT.
package simple

import (
	"net/http"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"
	"github.com/anz-bank/sysl-template/gen/jsonplaceholder"
)

// Handler interface for simple
type Handler interface {
	GetFoobarListHandler(w http.ResponseWriter, r *http.Request)
}

// ServiceHandler for simple API
type ServiceHandler struct {
	genCallback                           GenCallback
	serviceInterface                      *ServiceInterface
	jsonplaceholderjsonplaceholderService jsonplaceholder.Service
}

// NewServiceHandler for simple
func NewServiceHandler(genCallback GenCallback, serviceInterface *ServiceInterface, jsonplaceholderjsonplaceholderService jsonplaceholder.Service) *ServiceHandler {
	return &ServiceHandler{genCallback, serviceInterface, jsonplaceholderjsonplaceholderService}
}

// GetFoobarListHandler ...
func (s *ServiceHandler) GetFoobarListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetFoobarList == nil {
		s.genCallback.HandleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetFoobarListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	client := GetFoobarListClient{
		GetTodos: s.jsonplaceholderjsonplaceholderService.GetTodos,
	}

	todosresponse, err := s.serviceInterface.GetFoobarList(ctx, &req, client)
	if err != nil {
		s.genCallback.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, todosresponse)
}
