package service

import (
	"context"
	"github.com/tech-showcase/entertainment-service/model/movie"
)

type (
	movieService struct {
		movieClientEndpoint movie.ClientEndpoint
	}
	MovieService interface {
		Search(context.Context, string, int) (movie.ListPerPage, error)
	}
)

func NewMovieService(movieClientEndpoint movie.ClientEndpoint) MovieService {
	instance := movieService{}
	instance.movieClientEndpoint = movieClientEndpoint

	return &instance
}

func (instance *movieService) Search(ctx context.Context, keyword string, pageNumber int) (movieData movie.ListPerPage, err error) {
	movieData, err = instance.movieClientEndpoint.Search(ctx, keyword, pageNumber)
	if err != nil {
		return movie.ListPerPage{}, err
	}
	return movieData, nil
}
