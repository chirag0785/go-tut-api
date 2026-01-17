package initializers

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectToDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"))
	for range 10 {
		var err error
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err == nil {
			DB=db
			log.Println("Connected to the database successfully")
			return
		}
		time.Sleep(2*time.Second)
	}
	log.Fatal("Could not connect to the database")
}