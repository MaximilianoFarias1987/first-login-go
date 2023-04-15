package commons

import (
	"log"
	"login/models"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Getenv(key, defaultValue string) string {
	value, defined := os.LookupEnv(key)
	if !defined {
		return defaultValue
	}

	return value
}

func GetConnection() *gorm.DB {
	dsn := Getenv("CONNECTION_BD", "")
	log.Println(dsn)
	db, error := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	}

	return db
}

func Migrate() {
	db := GetConnection()
	log.Println("Start migration")

	db.AutoMigrate(&models.User{})
}
