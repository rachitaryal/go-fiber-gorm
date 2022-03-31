package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rachitaryal/go-fiber-gorm/controllers"
	"github.com/rachitaryal/go-fiber-gorm/database"
)

func welcome(c *fiber.Ctx) error {
	return c.Status(200).SendString("Let's Go!")
}

func setupRoutes(app *fiber.App){
	// welcome endpoint
	app.Get("/api", welcome)

	// user endpoints
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users", controllers.GetUsers)
	app.Get("/api/users/:id", controllers.GetUserByID)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)
}

func main(){
	//connect to db
	database.ConnectDB()

	// creating a new app 
	app := fiber.New()
	setupRoutes(app)
	port := ":3000"
	log.Fatal(app.Listen(port))

}