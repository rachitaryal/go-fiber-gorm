package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rachitaryal/go-fiber-gorm/database"
	"github.com/rachitaryal/go-fiber-gorm/models"
	"github.com/rachitaryal/go-fiber-gorm/serializers"
)

// HELPER FUNCTIONS

func findUserById(user *models.User, id int) error {
	database.Database.Db.Find(user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("User Doesn't Exist")
	}
	return nil
}

// CRUD OPERATIONS BELOW

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
	users := []models.User{}
	database.Database.Db.Find(&users)
	responseUsers := []serializers.UserSerializer{}
	for _, user := range users {
		responseUser := serializers.CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func GetUserByID(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err != nil{
		return c.Status(400).JSON(map[string]string{"msg": ":id is invalid"})
	}
	user := models.User{}
	if err := findUserById(&user, id); err != nil{
		return c.Status(400).JSON(map[string]string{
			"msg": "error with user",
			"error": err.Error(),
		})
	}

	responseUser := serializers.CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	} 

	user := models.User{}
	if err := findUserById(&user, id); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var updateUser UpdateUser
	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user.FirstName = updateUser.FirstName
	user.LastName = updateUser.LastName

	database.Database.Db.Save(&user)

	responseUser := serializers.CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)

}

func DeleteUser(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(200).JSON(err.Error())
	}
	var user models.User
	if err := findUserById(&user, id); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(200).JSON(err.Error())
	}
	return c.Status(200).JSON("User Deleted")
}