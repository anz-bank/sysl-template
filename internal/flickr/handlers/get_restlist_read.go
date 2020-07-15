package handlers

import (
	"context"

	"github.com/anz-bank/sysl-template/internal/gen/flickr"
)

// Handler for reading photos from flickr
type GetRestListRead struct{}

// GetRestListRead reads photos from flickr
func (h GetRestListRead) GetRestListRead(ctx context.Context,
	getRestListRequest *flickr.GetRestListRequest,
	client flickr.GetRestListClient) (*flickr.PhotoResource, error) {
	setJSONResponseContentType(ctx)

	if *getRestListRequest.Method == "flickr.photos.search" {
		switch *getRestListRequest.Tags {
		case "dog":
			return &flickr.PhotoResource{
				Page:    newFloat64(1),
				Pages:   newFloat64(1),
				Perpage: newFloat64(3),
				Photos: []flickr.Photo{
					{
						ID: newString("white golden retriever"),
						Tags: &flickr.Photo_tags{
							Tag: []flickr.Tag{
								{
									Raw: newString("dog"),
								},
							},
						},
						Urls: &flickr.Photo_urls{
							URL: []flickr.URL{
								{
									Type: newString("http://www.example.com/dog/wgr"),
								},
							},
						},
					},
					{
						ID: newString("I am JJ's dog Bingo"),
						Tags: &flickr.Photo_tags{
							Tag: []flickr.Tag{
								{
									Raw: newString("dog"),
								},
							},
						},
						Urls: &flickr.Photo_urls{
							URL: []flickr.URL{
								{
									Type: newString("http://www.example.com/dog/bingo"),
								},
							},
						},
					},
				},
				Total: newFloat64(2),
			}, nil
		case "cat":
			return &flickr.PhotoResource{
				Page:    newFloat64(1),
				Pages:   newFloat64(1),
				Perpage: newFloat64(3),
				Photos: []flickr.Photo{
					{
						ID: newString("White cat"),
						Tags: &flickr.Photo_tags{
							Tag: []flickr.Tag{
								{
									Raw: newString("cat"),
								},
							},
						},
						Urls: &flickr.Photo_urls{
							URL: []flickr.URL{
								{
									Type: newString("http://www.example.com/cat/wc"),
								},
							},
						},
					},
					{
						ID: newString("Black cat"),
						Tags: &flickr.Photo_tags{
							Tag: []flickr.Tag{
								{
									Raw: newString("cat"),
								},
							},
						},
						Urls: &flickr.Photo_urls{
							URL: []flickr.URL{
								{
									Type: newString("http://www.example.com/cat/bc"),
								},
							},
						},
					},
				},
				Total: newFloat64(2),
			}, nil
		case "rabbit":
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
		}
	}

	return nil, nil
}
