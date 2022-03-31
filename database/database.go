package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


type DbInstance struct {
	Db *gorm.DB
}
var Database DbInstance

func ConnectDB(){
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil{
		log.Fatalln("Failed to connect to database:", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to Database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	// todo: add migrations code here
	Database = DbInstance{
		Db: db,
	}
}