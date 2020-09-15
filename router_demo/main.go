package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r := gin.Default()

	//static asset

	r.Static("/assets", "./assets")
	r.StaticFS("/static", http.Dir("./static"))
	r.StaticFile("/favicon.ico", "./favicon.ico")

	//http method demo
	r.GET("/get", func(c *gin.Context) {
		c.String(200, "get")
	})

	r.POST("/post", func(c *gin.Context){
		c.String(200, "post")
	})
	
	r.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.String(200, "delete")
	})

	r.Any("/any", func(c *gin.Context) {
		c.String(200, "any")
	})

	r.Run()
}
