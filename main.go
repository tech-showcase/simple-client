package main

import (
	"context"
	"fmt"
	"github.com/tech-showcase/entertainment-service/cmd"
	"github.com/tech-showcase/entertainment-service/global"
	"github.com/tech-showcase/entertainment-service/helper"
	"github.com/tech-showcase/entertainment-service/transport"
	"net/http"
)

func main() {
	fmt.Println("Hi, I am Simple Client!")

	args := cmd.Parse()
	portStr := fmt.Sprintf(":%d", args.Port)

	ctx := context.Background()

	logger := helper.NewLogger(portStr)

	instances := make([]string, 0)
	searchMovieInstance, _ := helper.JoinURL(global.Configuration.APIGatewayAddress, "/movie/search")
	instances = append(instances, searchMovieInstance)

	http.Handle("/search", transport.NewSearchMovieHTTPServer(ctx, instances, logger))
	http.ListenAndServe(portStr, nil)
}
