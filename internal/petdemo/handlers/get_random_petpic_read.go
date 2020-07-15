package handlers

import (
	"context"

	"github.com/anz-bank/sysl-template/internal/gen/flickr"
	"github.com/anz-bank/sysl-template/internal/gen/petdemo"
	"github.com/anz-bank/sysl-template/internal/gen/petstore"
)

// Handler for fetching random pet pic
type GetRandomPetPicListRead struct{}

// GetRandomPetPicListRead reads random pic from downstream
func (h GetRandomPetPicListRead) GetRandomPetPicListRead(ctx context.Context,
	getRandomPetPicListRequest *petdemo.GetRandomPetPicListRequest,
	client petdemo.GetRandomPetPicListClient) (*petdemo.PetResponse, error) {
	setJSONResponseContentType(ctx)
	reqPetstore := petstore.GetPetListRequest{}
	// get the pet list
	pet, err := client.GetPetList(ctx, &reqPetstore)
	if err != nil {
		return nil, err
	}
	reqFlickr := flickr.GetRestListRequest{
		Method: newString("flickr.photos.search"),
		Tags:   newString(*pet.Tag),
	}
	photo, err := client.GetRestList(ctx, &reqFlickr)
	if err != nil {
		return nil, err
	}

	// return the result
	return &petdemo.PetResponse{
		Name: *(photo.Photos[0].ID),
		URI:  *photo.Photos[0].Urls.URL[0].Type,
	}, nil
}
