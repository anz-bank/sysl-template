package handlers

import (
	"context"

	"github.com/anz-bank/sysl-template/internal/gen/flickr"
	"github.com/anz-bank/sysl-template/internal/gen/petdemo"
	"github.com/anz-bank/sysl-template/internal/gen/petstore"
)

/* - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 * get random pet pic
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - */

type GetRandomPetPicListRead struct {
	petClient petdemo.GetRandomPetPicListClient
	ctx       context.Context
}

func (h GetRandomPetPicListRead) GetRandomPetPicListRead(ctx context.Context,
	getRandomPetPicListRequest *petdemo.GetRandomPetPicListRequest,
	client petdemo.GetRandomPetPicListClient) (*petdemo.PetResponse, error) {
	setJSONResponseContentType(ctx)
	reqPetstore := petstore.GetPetsListRequest{
		Limit: newInt(10),
	}
	// get the pet list
	pet, err := h.petClient.GetPetsList(ctx, &reqPetstore)
	if err != nil {
		return nil, err
	}
	reqFlickr := flickr.GetRestListRequest{
		Method: newString("flickr%2Ephotos%2Esearch"),
		Tags:   newString(pet.Name),
	}
	photo, err := h.petClient.GetRestList(ctx, &reqFlickr)
	if err != nil {
		return nil, err
	}

	// return the result
	return &petdemo.PetResponse{
		Name: *(photo.Photos[1].ID),
		URI:  *photo.Photos[1].Urls.URL[0].Type,
	}, nil
}
