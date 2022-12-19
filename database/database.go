package database

import (
	"log"
	"os"

	"github.com/alphanumericentity/fiber-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func Connect() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("could not connect to Db \n", err.Error())
		os.Exit(2)
	}

	log.Println("connected to database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("rinning migrations")
	// migrations
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	Database = DbInstance{Db: db}

}
