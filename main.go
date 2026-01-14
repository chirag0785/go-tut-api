package main

import (
	"github.com/chirag0785/go-tut-api/controllers"
	"github.com/chirag0785/go-tut-api/initializers"
	"github.com/chirag0785/go-tut-api/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB() 
}
func main() {
  r := gin.Default()
  r.POST("/posts", middleware.ValidatePost(), controllers.PostsCreate)
  r.GET("/posts", controllers.PostsIndex)
  r.GET("/posts/:id", controllers.PostsShow)
  r.POST("/posts/:id",controllers.PostsUpdate)
  r.DELETE("/posts/:id",controllers.PostsDelete)
  r.Run()
}