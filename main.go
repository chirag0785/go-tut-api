package main

import (
	"log"
	"os"

	"github.com/chirag0785/go-tut-api/controllers"
	"github.com/chirag0785/go-tut-api/initializers"
	"github.com/chirag0785/go-tut-api/middleware"
	"github.com/chirag0785/go-tut-api/migrate"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	if os.Getenv("RUN_MIGRATIONS") == "true" {
		if err := migrate.MigrateToDB(); err != nil {
			log.Fatal("Failed to migrate database:", err)
		}
	}
}
func main() {
  r := gin.Default()
  r.POST("/posts", middleware.Authenticate(), middleware.ValidatePost(), controllers.PostsCreate)
  r.GET("/posts", controllers.PostsIndex)
  r.GET("/posts/:id", middleware.Authenticate(), controllers.PostsShow)
  r.POST("/posts/:id", middleware.Authenticate(), controllers.PostsUpdate)
  r.DELETE("/posts/:id", middleware.Authenticate(), controllers.PostsDelete)

  r.POST("/users", controllers.UsersCreate)
  r.POST("/users/login", controllers.UsersLogin)

  r.GET("/posts/me",middleware.Authenticate(),controllers.PostsUserPosts)
  r.Run()
}