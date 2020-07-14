package handlers

import (
	"testing"

	"github.com/anz-bank/sysl-template/internal/gen/petdemo"
	"github.com/stretchr/testify/require"
)

func TestGetRandomPetPicListRead(t *testing.T) {
	req := petdemo.GetRandomPetPicListRequest{}
	client := petdemo.GetRandomPetPicListClient{}

	// get the random pet
	response, err := GetRandomPetPicListRead{}.GetRandomPetPicListRead(newRequestContext(), &req, client)

	// check err
	require.Nil(t, err)

	// verify pet
	require.Contains(t, response.Name, "dog")
}
