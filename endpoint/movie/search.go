package movie

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/tech-showcase/entertainment-service/model/movie"
	"github.com/tech-showcase/entertainment-service/service"
	"net/http"
)

type (
	SearchMovieRequest struct {
		movie.SearchMovieRequest
	}
	SearchMovieResponse struct {
		movie.SearchMovieResponse
	}
)

func makeSearchMovieEndpoint(svc service.MovieService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SearchMovieRequest)

		movieData, err := svc.Search(ctx, req.Keyword, req.PageNumber)
		if err != nil {
			return SearchMovieResponse{}, err
		}

		return SearchMovieResponse{
			SearchMovieResponse: movie.SearchMovieResponse{
				ListPerPage: movieData,
			},
		}, nil
	}
}

func decodeSearchMovieRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req SearchMovieRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
