package config

import (
	"log"

	"github.com/tolumadamori/sch_manager/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect to the DB. Credentials are passed using environment variables.
func ConnectDB() (*gorm.DB, error) {

	dsn := "root:secret@tcp(mysql:3306)/sch-manager-db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the Users DB")
	}

	if err = db.AutoMigrate(&models.Student{}, &models.Course{}); err != nil {
		log.Println(err)
	}

	return db, err
}
