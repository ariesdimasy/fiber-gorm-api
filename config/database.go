package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ariesdimasy/fiber-gorm-api/models"

	"gorm.io/driver/mysql"
	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnecDb() {
	dsn := "root:@tcp(127.0.0.1:3306)/mp_backend?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	fmt.Println("RUN")
	if err != nil {
		fmt.Println("Failed Connect to database")
		log.Fatal("Failed to Connect to the database! \n ", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	DB = db
}
