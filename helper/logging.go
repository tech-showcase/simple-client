package helper

import (
	"github.com/go-kit/kit/log"
	"os"
)

func NewLogger(address string) log.Logger {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", address, "caller", log.DefaultCaller)

	return logger
}
