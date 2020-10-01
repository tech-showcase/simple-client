package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/tech-showcase/entertainment-service/cmd"
	"github.com/tech-showcase/entertainment-service/config"
	"github.com/tech-showcase/entertainment-service/model/movie"
	"github.com/tech-showcase/entertainment-service/service"
)

func init() {
	config.Instance = config.Read()
}

func main() {
	fmt.Println("Hi, I am Simple Client!")

	args := cmd.Parse()

	ctx := context.Background()

	movieClientEndpoint, err := movie.NewMovieClientEndpoint(config.Instance.APIGatewayAddress)
	if err != nil {
		panic(err)
	}
	movieService := service.NewMovieService(movieClientEndpoint)
	movieData, err := movieService.Search(ctx, args.Keyword, args.PageNumber)
	if err != nil {
		panic(err)
	}

	spew.Dump(movieData)
}
