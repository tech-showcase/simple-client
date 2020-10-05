package movie

import (
	"context"
	"github.com/go-kit/kit/endpoint"
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
		search  endpoint.Endpoint
	}
	ClientEndpoint interface {
		Search(ctx context.Context, keyword string, pageNumber int) (movieData ListPerPage, err error)
	}
)

func NewMovieClientEndpoint(address string) (ClientEndpoint, error) {
	instance := clientEndpoint{}

	if !strings.HasPrefix(address, "http") {
		address = "http://" + address
	}

	u, err := url.Parse(address)
	if err != nil {
		return nil, err
	}
	instance.address = u

	instance.search = makeSearchMovieClientEndpoint(u)

	return &instance, nil
}

func (instance *clientEndpoint) Search(ctx context.Context, keyword string, pageNumber int) (movieData ListPerPage, err error) {
	req := SearchMovieRequest{
		Keyword:    keyword,
		PageNumber: pageNumber,
	}
	response, err := instance.search(ctx, req)
	if err != nil {
		return ListPerPage{}, err
	}

	res, _ := response.(SearchMovieResponse)
	return res.ListPerPage, nil
}
