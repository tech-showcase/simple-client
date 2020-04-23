package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/tech-showcase/entertainment-service/model"
	"github.com/tech-showcase/entertainment-service/presenter"
	"github.com/tech-showcase/entertainment-service/service"
	"net/http"
)

func makeSearchMovieEndpoint(ctx context.Context, svc service.MovieService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(presenter.SearchMovieRequest)

		movieData, err := svc.Search(req.Keyword, req.PageNumber)
		if err != nil {
			return presenter.SearchMovieResponse{}, nil
		}
		return presenter.SearchMovieResponse{
			SearchMovieResponse: model.SearchMovieResponse{
				MovieListPerPage: movieData,
			},
		}, nil
	}
}

func decodeSearchMovieRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req presenter.SearchMovieRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
