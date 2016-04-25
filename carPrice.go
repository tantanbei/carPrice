package main

import (
	"carPrice/api"
	"net/http"

	"github.com/gocraft/web"
)

func main() {
	router := web.New(api.Api{}). // Create your router
					Middleware(web.LoggerMiddleware).                    // Use some included middleware
					Middleware(web.ShowErrorsMiddleware).                // ...
					Get("/chexiang/produce/id", (*api.Api).GetCarPrices) // Add a route

	//start serve
	http.ListenAndServe(":8080", router) // Start the server!www

}
