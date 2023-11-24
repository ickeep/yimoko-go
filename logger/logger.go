// Package logger  logger.go
package logger

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/ickeep/yimoko-go/config"
)

// GetLogger _
func GetLogger(service *config.Server, logger log.Logger) log.Logger {
	return log.With(logger,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", service.Id,
		"service.name", service.Name,
		"service.version", service.Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}

// GetStdLogger _
func GetStdLogger(conf *config.Config) log.Logger {
	logger := log.NewStdLogger(os.Stdout)
	return GetLogger(conf.Server, logger)
}
