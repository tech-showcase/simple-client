package main

import (
	"context"
	"fmt"
	"github.com/tech-showcase/entertainment-service/cmd"
	"github.com/tech-showcase/entertainment-service/config"
	"github.com/tech-showcase/entertainment-service/helper"
	"github.com/tech-showcase/entertainment-service/transport"
	"net/http"
)

func init() {
	config.Instance = config.Read()
}

func main() {
	fmt.Println("Hi, I am Simple Client!")

	args := cmd.Parse()
	portStr := fmt.Sprintf(":%d", args.Port)

	ctx := context.Background()

	logger := helper.NewLogger(portStr)

	instances := make([]string, 0)
	searchMovieInstance, _ := helper.JoinURL(config.Instance.APIGatewayAddress, "/movie/search")
	instances = append(instances, searchMovieInstance)

	http.Handle("/search", transport.NewSearchMovieHTTPServer(ctx, instances, logger))
	http.ListenAndServe(portStr, nil)
}
