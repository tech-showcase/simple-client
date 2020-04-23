package service

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/sony/gobreaker"
	"github.com/tech-showcase/entertainment-service/model"
	"golang.org/x/time/rate"
	"time"
)

func MakeMovieServiceMiddleware(ctx context.Context, instance []string, logger log.Logger) MovieServiceMiddleware {
	if len(instance) < 1 {
		logger.Log("proxy_to", "none")
		return func(next MovieService) MovieService { return next }
	}

	logger.Log("proxy_to", fmt.Sprint(instance))

	searchMovieEndpoint := model.MakeSearchMovieHTTPClient(ctx, instance[0])
	searchMovieEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(searchMovieEndpoint)
	searchMovieEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))(searchMovieEndpoint)

	return func(next MovieService) MovieService {
		return &movieProxy{
			ctx:                 ctx,
			searchMovieEndpoint: searchMovieEndpoint,
		}
	}
}
