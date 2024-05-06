package database

import (
	"github.com/cleyton1986/client-cleancode-solid/internal/entities"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitializeDatabase() (*gorm.DB, error) {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=user dbname=devdb password=password sslmode=disable")
    if err != nil {
        return nil, err
    }
    db.AutoMigrate(&entities.User{})
    return db, nil
}
