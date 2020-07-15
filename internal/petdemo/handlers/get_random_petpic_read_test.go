package handlers

import (
	"context"
	"testing"

	"github.com/anz-bank/sysl-template/internal/gen/flickr"
	"github.com/anz-bank/sysl-template/internal/gen/petdemo"
	"github.com/anz-bank/sysl-template/internal/gen/petstore"
	"github.com/stretchr/testify/require"
)

func TestGetRandomPetPicListRead(t *testing.T) {
	req := petdemo.GetRandomPetPicListRequest{}
	client := petdemo.GetRandomPetPicListClient{
		GetPetList: func(ctx context.Context, req *petstore.GetPetListRequest) (*petstore.Pet, error) {
			setJSONResponseContentType(ctx)
			dummyPet := petstore.Pet{
				ID:   1,
				Name: "dummy pet dog",
				Tag:  newString("dog"),
			}
			return &dummyPet, nil
		},
		GetRestList: func(ctx context.Context, req *flickr.GetRestListRequest) (*flickr.PhotoResource, error) {
			return &flickr.PhotoResource{
				Page:    newFloat64(1),
				Pages:   newFloat64(1),
				Perpage: newFloat64(3),
				Photos: []flickr.Photo{
					{
						ID: newString("I am peter rabbit"),
						Tags: &flickr.Photo_tags{
							Tag: []flickr.Tag{
								{
									Raw: newString("rabbit"),
								},
							},
						},
						Urls: &flickr.Photo_urls{
							URL: []flickr.URL{
								{
									Type: newString("http://www.example.com/rabbit/pr"),
								},
							},
						},
					},
				},
				Total: newFloat64(2),
			}, nil
		},
	}

	// get the random pet
	response, err := GetRandomPetPicListRead{}.GetRandomPetPicListRead(newRequestContext(), &req, client)

	// check err
	require.Nil(t, err)

	// verify pet
	require.Contains(t, response.Name, "rabbit")
}
