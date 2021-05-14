package main

import (
	"gee/gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/index", func(context *gee.Context) {
		context.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("v1")
	{
		v1.GET("/", func(context *gee.Context) {
			context.HTML(http.StatusOK, "<h1>Index Gee v1</h1>")
		})

		v1.GET("/helo", func(context *gee.Context) {
			context.HTML(http.StatusOK, "<h1>Index Gee v1</h1>")
		})
	}

	r.Run(":9999")
}