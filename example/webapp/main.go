package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gocraft/web"
	"github.com/hnakamur/gocraft-web-logger"
)

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
	router := web.New(Context{}). // Create your router
					Middleware(logger.LoggerMiddleware).  // Use some included middleware
					Middleware(web.ShowErrorsMiddleware). // ...
					Middleware((*Context).SetHelloCount). // Your own middleware!
					Get("/", (*Context).SayHello)         // Add a route
	http.ListenAndServe("localhost:3000", router) // Start the server!
}
