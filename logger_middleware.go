package logger

import (
	"time"

	"github.com/gocraft/web"
	"github.com/op/go-logging"
)

var log *logging.Logger

func SetLogger(logger *logging.Logger) {
	log = logger
}

// LoggerMiddleware is generic middleware that will log requests to Logger (by default, Stdout).
func LoggerMiddleware(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	startTime := time.Now()

	next(rw, req)

	duration := time.Since(startTime).Nanoseconds()
	var durationUnits string
	switch {
	case duration > 2000000:
		durationUnits = "ms"
		duration /= 1000000
	case duration > 1000:
		durationUnits = "μs"
		duration /= 1000
	default:
		durationUnits = "ns"
	}

	log.Info("duration:%d%s\tstatus:%d\tmessage:%s\n", duration, durationUnits, rw.StatusCode(), req.URL.Path)
}
