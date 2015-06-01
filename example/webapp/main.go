package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gocraft/web"
	"github.com/hnakamur/gocraft-web-logger"
	"github.com/monochromegane/conflag"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("webapp")

type Context struct {
	HelloCount int
}

func (c *Context) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HelloCount = 3
	next(rw, req)
}

func (c *Context) SayHello(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, strings.Repeat("Hello ", c.HelloCount), "World!")
}

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config.toml", "config file path")
	var logFormat string
	flag.StringVar(&logFormat, "logformat", "time:%{time:2006-01-02T15:04:05.000}\tlevel:%{level:.4s}\t%{message}", "log format")
	var logFileName string
	flag.StringVar(&logFileName, "logfilename", "-", "log file (- for stdout)")
	var listenAddress string
	flag.StringVar(&listenAddress, "listenaddress", ":3000", "listen address (host:port)")

	flag.Parse()
	if args, err := conflag.ArgsFrom(configFile); err == nil {
		flag.CommandLine.Parse(args)
	}
	flag.Parse()

	logWriter := os.Stderr
	if logFileName != "-" {
		logFile, err := os.Create(logFileName)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("cannot open log file: %s\n", err))
			os.Exit(1)
		}
		defer logFile.Close()
		logWriter = logFile
	}
	backend := logging.NewLogBackend(logWriter, "", 0)
	formatter, err := logging.NewStringFormatter(logFormat)
	if err != nil {
		os.Stderr.WriteString("invalid logFormat\n")
		os.Exit(1)
	}
	logging.SetBackend(logging.NewBackendFormatter(backend, formatter))

	router := web.New(Context{}). // Create your router
					Middleware(logger.LoggerMiddlewareFactory(log)). // Use some included middleware
					Middleware(web.ShowErrorsMiddleware).            // ...
					Middleware((*Context).SetHelloCount).            // Your own middleware!
					Get("/", (*Context).SayHello)                    // Add a route
	http.ListenAndServe(listenAddress, router) // Start the server!
}
