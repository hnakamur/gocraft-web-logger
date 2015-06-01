package logger

import (
	"log"
	"os"
	"time"

	"github.com/gocraft/web"
)

// Logger can be set to your own logger. Logger only applies to the LoggerMiddleware.
var Logger = log.New(os.Stdout, "", 0)

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

	Logger.Printf("%s %d%s %d %s\n", startTime.Format(time.RFC3339), duration, durationUnits, rw.StatusCode(), req.URL.Path)
}
