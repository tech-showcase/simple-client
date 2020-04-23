package transport

import (
	"context"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/tech-showcase/entertainment-service/service"
	"net/http"
)

func NewSearchMovieHTTPServer(ctx context.Context, instances []string, logger log.Logger) http.Handler {
	movieService := service.NewMovieService()
	movieService = service.MakeMovieServiceMiddleware(ctx, instances, logger)(movieService)

	searchMovieHandler := httptransport.NewServer(
		makeSearchMovieEndpoint(ctx, movieService),
		decodeSearchMovieRequest,
		encodeResponse,
	)

	return  searchMovieHandler
}
