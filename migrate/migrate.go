package migrate

import (
	"github.com/chirag0785/go-tut-api/initializers"
	"github.com/chirag0785/go-tut-api/models"
)

func MigrateToDB() error {
	return initializers.DB.AutoMigrate(&models.User{}, &models.Post{})
}
