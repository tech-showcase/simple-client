package movie

import (
	"github.com/tech-showcase/entertainment-service/endpoint"
	"github.com/tech-showcase/entertainment-service/service"
)

type (
	Endpoint struct {
		Search endpoint.HTTPEndpoint
	}
)

func NewMovieEndpoint(movieService service.MovieService) (movieEndpoint Endpoint) {
	movieEndpoint.Search = endpoint.HTTPEndpoint{
		Endpoint: makeSearchMovieEndpoint(movieService),
		Decoder:  decodeSearchMovieRequest,
		Encoder:  encodeResponse,
	}

	return movieEndpoint
}
