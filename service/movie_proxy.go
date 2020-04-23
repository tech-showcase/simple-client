package service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/tech-showcase/entertainment-service/model"
)

type (
	movieProxy struct {
		ctx                 context.Context
		searchMovieEndpoint endpoint.Endpoint
	}
)

func (instance *movieProxy) Search(keyword string, pageNumber int) (movieData model.MovieListPerPage, err error) {
	req := model.SearchMovieRequest{
		Keyword:    keyword,
		PageNumber: pageNumber,
	}
	resp, err := instance.searchMovieEndpoint(instance.ctx, req)
	if err != nil {
		return
	}

	movieData = resp.(model.SearchMovieResponse).MovieListPerPage

	return
}
