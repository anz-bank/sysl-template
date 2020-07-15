package handlers

import (
	"context"

	"github.com/anz-bank/sysl-template/internal/gen/petstore"
)

// Handler for reading photos from petstore
type GetPetsListRead struct{}

// GetPetsListRead reads photos from petstore
func (h GetPetsListRead) GetPetsListRead(ctx context.Context,
	getPetsListRequest *petstore.GetPetListRequest,
	client petstore.GetPetListClient) (*petstore.Pet, error) {
	setJSONResponseContentType(ctx)

	// create the dummy list of pets
	dummyPet := petstore.Pet{
		ID:   1,
		Name: "dummy pet dog",
		Tag:  newString("dog"),
	}

	// return the result
	return &dummyPet, nil
}
