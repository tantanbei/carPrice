package main

import (
	"carPrice/car"
	"fmt"
	"net/http"

	"github.com/gocraft/web"
)

var n = 888

type Context struct {
	HelloCount int
}

func (c *Context) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HelloCount = 3
	next(rw, req)
}

func (c *Context) SayHello(rw web.ResponseWriter, req *web.Request) {
	carPrice := car.CarPriceById(n)
	bs, err := carPrice.ParseToJson()
	if err != nil {
		panic(err)
	}
	fmt.Fprint(rw, string(bs))

}

func main() {
	str := fmt.Sprint("/car/", n)
	router := web.New(Context{}). // Create your router
					Middleware(web.LoggerMiddleware).     // Use some included middleware
					Middleware(web.ShowErrorsMiddleware). // ...
					Middleware((*Context).SetHelloCount). // Your own middleware!
					Get(str, (*Context).SayHello)         // Add a route
	http.ListenAndServe("localhost:3000", router) // Start the server!

}
