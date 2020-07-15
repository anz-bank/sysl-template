// Code generated by sysl DO NOT EDIT.
package petdemo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"
)

// Service interface for Petdemo
type Service interface {
	GetRandomPetPicList(ctx context.Context, req *GetRandomPetPicListRequest) (*PetResponse, error)
}

// Client for Petdemo API
type Client struct {
	client *http.Client
	url    string
}

// NewClient for Petdemo
func NewClient(client *http.Client, serviceURL string) *Client {
	return &Client{client, serviceURL}
}

// GetRandomPetPicList ...
func (s *Client) GetRandomPetPicList(ctx context.Context, req *GetRandomPetPicListRequest) (*PetResponse, error) {
	required := []string{}
	var okResponse PetResponse
	var errorResponse GenericError
	u, err := url.Parse(fmt.Sprintf("%s/random-pet-pic", s.url))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest(ctx, s.client, "GET", u.String(), nil, required, &okResponse, &errorResponse)
	if err != nil {
		response, ok := err.(*restlib.HTTPResult)
		if !ok {
			return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Petdemo <- GET "+u.String(), err)
		}
		return nil, common.CreateDownstreamError(ctx, common.DownstreamResponseError, response.HTTPResponse, response.Body, &errorResponse)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkPetResponseResponse, ok := result.Response.(*PetResponse)
	if ok {
		valErr := validator.Validate(OkPetResponseResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkPetResponseResponse, nil
	}

	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}
