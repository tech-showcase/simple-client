package global

import (
	"github.com/tech-showcase/entertainment-service/config"
)

var Configuration = config.Config{}

func init() {
	var err error
	Configuration, err = config.Parse()
	if err != nil {
		panic(err)
	}
}
