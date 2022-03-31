package database

import (
	"log"
	"os"

	"github.com/rachitaryal/go-fiber-gorm/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


type DbInstance struct {
	Db *gorm.DB
}
var Database DbInstance

func ConnectDB(){
	// create an api.db database if it is not already there
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil{
		log.Fatalln("Failed to connect to database:", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to Database")
	db.Logger = logger.Default.LogMode(logger.Info)

	// Adding automigration for all the models and running migrations
	log.Println("Running Migrations")
	err = db.AutoMigrate(&models.User{FirstName: "Rachit", LastName: "Aryal"}, &models.Product{}, &models.Order{})
	if err != nil{
		log.Fatalln("AutoMigration Failed: ", err.Error())
	}
	log.Println("\nAutoMigration Completed")
	Database = DbInstance{
		Db: db,
	}
}