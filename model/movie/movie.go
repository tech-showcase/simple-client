package movie

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/tech-showcase/entertainment-service/helper"
	"net/url"
	"strings"
)

type (
	ListPerPage struct {
		Response     string `json:"Response"`
		Search       []Item `json:"Search"`
		TotalResults string `json:"totalResults"`
	}
	Item struct {
		Poster string `json:"Poster"`
		Title  string `json:"Title"`
		Type   string `json:"Type"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
	}

	clientEndpoint struct {
		address *url.URL
	}
	ClientEndpoint interface {
		Search(ctx context.Context, keyword string, pageNumber int) (movieData ListPerPage, err error)
	}
)

func NewMovieClientEndpoint(address string) ClientEndpoint {
	instance := clientEndpoint{}

	if !strings.HasPrefix(address, "http") {
		address = "http://" + address
	}

	u, err := url.Parse(address)
	if err != nil {
		panic(err)
	}
	instance.address = u

	return &instance
}

func (instance *clientEndpoint) Search(ctx context.Context, keyword string, pageNumber int) (movieData ListPerPage, err error) {
	searchMovieURL, _ := helper.JoinURL(instance.address, "/movie/search")

	movieClientEndpoint := httptransport.NewClient(
		"POST",
		searchMovieURL,
		encodeHTTPRequest,
		decodeSearchMovieHTTPResponse,
	).Endpoint()

	req := SearchMovieRequest{
		Keyword:    keyword,
		PageNumber: pageNumber,
	}
	response, err := movieClientEndpoint(ctx, req)
	if err != nil {
		return ListPerPage{}, err
	}

	res, _ := response.(SearchMovieResponse)
	return res.ListPerPage, nil
}
