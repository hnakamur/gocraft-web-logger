package logger

import (
	"time"

	"github.com/gocraft/web"
	"github.com/op/go-logging"
)

func LoggerMiddlewareFactory(logger *logging.Logger) func(web.ResponseWriter, *web.Request, web.NextMiddlewareFunc) {
	return func(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
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

		logger.Info("duration:%d%s\tstatus:%d\tmessage:%s\n", duration, durationUnits, rw.StatusCode(), req.URL.Path)
	}
}
