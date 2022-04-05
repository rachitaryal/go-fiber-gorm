package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rachitaryal/go-fiber-gorm/database"
	"github.com/rachitaryal/go-fiber-gorm/models"
	"github.com/rachitaryal/go-fiber-gorm/serializers"
)

// HELPER FUNCTIONS

func findOrderById(order *models.Order, id int) error {
	database.Database.Db.Find(order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("Order Doesn't Exist")
	}
	return nil
}


func CreateOrder(c *fiber.Ctx) error{
	var order models.Order
	if err := c.BodyParser(&order); err != nil{
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUserById(&user, order.UserReferer); err != nil{
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProductById(&product, order.ProductReferer); err != nil{
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	responseUser := serializers.CreateResponseUser(user)
	responseProduct := serializers.CreateResponseProduct(product)
	responseOrder := serializers.CreateResponseOrder(order, responseUser, responseProduct)


	return c.Status(200).JSON(responseOrder)

}

func GetOrders(c *fiber.Ctx) error {
	var orders []models.Order
	database.Database.Db.Find(&orders)

	var responseOrders []serializers.OrderSerializer
	for _, order := range orders{
		var user models.User
		if err := findUserById(&user, order.UserReferer); err != nil{
			return c.Status(400).JSON(err.Error())
		}
		var product models.Product
		if err := findProductById(&product, order.ProductReferer); err != nil{
			return c.Status(400).JSON(err.Error())
		}

		responseUser := serializers.CreateResponseUser(user)
		responseProduct := serializers.CreateResponseProduct(product)
		responseOrder := serializers.CreateResponseOrder(order, responseUser, responseProduct)
		responseOrders = append(responseOrders, responseOrder)

	}

	return c.Status(200).JSON(responseOrders)

}


func GetOrderByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil{
		return c.Status(400).JSON(map[string]string{"msg": ":id is invalid"})
	}
	order := models.Order{}
	if err := findOrderById(&order, id); err != nil{
		return c.Status(400).JSON(map[string]string{
			"msg": "error with order",
			"error": err.Error(),
		})
	}
	var user models.User
	if err := findUserById(&user, order.UserReferer); err != nil{
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProductById(&product, order.ProductReferer); err != nil{
		return c.Status(400).JSON(err.Error())
	}

	responseUser := serializers.CreateResponseUser(user)
	responseProduct := serializers.CreateResponseProduct(product)

	responseOrder := serializers.CreateResponseOrder(order, responseUser, responseProduct)
	return c.Status(200).JSON(responseOrder)
}