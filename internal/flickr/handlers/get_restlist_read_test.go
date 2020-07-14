package handlers

import (
	"testing"

	"github.com/anz-bank/sysl-template/internal/gen/flickr"
	"github.com/stretchr/testify/require"
)

func TestReadGetPetsList(t *testing.T) {
	req := flickr.GetRestListRequest{}
	client := flickr.GetRestListClient{}
	req.Method = newString("flickr%2Ephotos%2Esearch")
	req.Tags = newString("dog")

	// get the pet photo list
	response, err := GetRestListRead{}.GetRestListRead(newRequestContext(), &req, client)

	// err check
	require.Nil(t, err)
	require.True(t, len(response.Photos) > 0)

	// verify one of the pet photos
	require.Equal(t, *response.Photos[1].ID, "I am JJ's dog Bingo")
	url := response.Photos[1].Urls.URL[0].Type
	require.Equal(t, *url, "http://www.example.com/dog/bingo")
}
