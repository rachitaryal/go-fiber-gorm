package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rachitaryal/go-fiber-gorm/database"
	"github.com/rachitaryal/go-fiber-gorm/models"
	"github.com/rachitaryal/go-fiber-gorm/serializers"
)


func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := serializers.CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error{
	return c.Status(200).JSON(map[string]string{"msg": "get user working"})
}

// func GetUserByID(c *fiber.Ctx) error{}

// func UpdateUser(c *fiber.Ctx) error{}

// func DeleteUser(c *fiber.Ctx) error{}