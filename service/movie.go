package service

import (
	"github.com/tech-showcase/entertainment-service/model"
)

type (
	movieService struct{}
	MovieService interface {
		Search(string, int) (model.MovieListPerPage, error)
	}

	MovieServiceMiddleware func(MovieService) MovieService
)

func NewMovieService() MovieService {
	instance := movieService{}

	return &instance
}

func (instance *movieService) Search(keyword string, pageNumber int) (movieData model.MovieListPerPage, err error) {
	return
}
