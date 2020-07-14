package handlers

import (
	"context"

	"github.com/anz-bank/sysl-template/internal/gen/petstore"
)

/* - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 * get petlist read
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - */

type GetPetsListRead struct{}

func (h GetPetsListRead) GetPetsListRead(ctx context.Context,
	getPetsListRequest *petstore.GetPetsListRequest,
	client petstore.GetPetsListClient) (*petstore.Pet, error) {
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
