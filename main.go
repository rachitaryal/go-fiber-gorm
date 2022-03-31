package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rachitaryal/go-fiber-gorm/database"
)

func main(){
	//connect to db
	database.ConnectDB()

	// creating a new app 
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"go": "fiber",
		})
	})
	port := ":3000"
	log.Fatal(app.Listen(port))

}