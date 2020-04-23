package model

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type (
	MovieClientEndpoints struct {
		SearchMovieClientEndpoint endpoint.Endpoint
	}

	SearchMovieRequest struct {
		Keyword    string `json:"keyword"`
		PageNumber int    `json:"page_number"`
	}
	SearchMovieResponse struct {
		MovieListPerPage
	}
	MovieListPerPage struct {
		Response     string      `json:"Response"`
		Search       []MovieItem `json:"Search"`
		TotalResults string      `json:"totalResults"`
	}
	MovieItem struct {
		Poster string `json:"Poster"`
		Title  string `json:"Title"`
		Type   string `json:"Type"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
	}
)

func MakeSearchMovieHTTPClient(ctx context.Context, instance string) endpoint.Endpoint {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		panic(err)
	}

	return httptransport.NewClient(
		"POST",
		u,
		encodeHTTPRequest,
		decodeSearchMovieHTTPResponse,
	).Endpoint()
}

func encodeHTTPRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func decodeSearchMovieHTTPResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var req SearchMovieResponse
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
