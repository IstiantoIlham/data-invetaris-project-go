package database

import (
	. "data-invetaris/config"
	"data-invetaris/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	config := Config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}
	err = DB.AutoMigrate(&models.Product{}, &models.Inventory{}, &models.Order{})
	if err != nil {
		log.Println("Cannot migrate table")
	}
	log.Println("Success connected to database")
}
