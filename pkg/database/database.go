package database

import (
	"log"
	"time"

	"github.com/cleyton1986/client-cleancode-solid/internal/entities"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitializeDatabase() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open("postgres", "host=db port=5432 user=user dbname=devdb password=password sslmode=disable")
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database. Retry in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entities.User{})
	return db, nil
}
