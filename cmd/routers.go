package main

import "github.com/gin-gonic/gin"

func (app *application) routers() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/create", app.create)
	r.GET("/view", app.view)
	return r
}
