package main

import (
	"github.com/chirag0785/go-tut-api/initializers"
	"github.com/chirag0785/go-tut-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
