package handlers

import (
	"context"

	util "github.com/anz-bank/sysl-template/internal"
	"github.com/anz-bank/sysl-template/internal/gen/petstore"
)

// GetPetsListRead reads photos from petstore
func GetPetsListRead(ctx context.Context,
	getPetsListRequest *petstore.GetPetListRequest,
	client petstore.GetPetListClient) (*petstore.Pet, error) {
	util.SetJSONResponseContentType(ctx)

	// create the dummy list of pets
	dummyPet := petstore.Pet{
		ID:   1,
		Name: "dummy pet dog",
		Tag:  util.NewString("dog"),
	}

	// return the result
	return &dummyPet, nil
}
