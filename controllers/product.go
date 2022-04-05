package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rachitaryal/go-fiber-gorm/database"
	"github.com/rachitaryal/go-fiber-gorm/models"
	"github.com/rachitaryal/go-fiber-gorm/serializers"
)

// HELPER FUNCTIONS

func findProductById(product *models.Product, id int) error{
	database.Database.Db.Find(product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("Product Doesn't Exist")
	}
	return nil
}

// Product CRUD

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)
	responseUser := serializers.CreateResponseProduct(product)
	return c.Status(200).JSON(responseUser)
}

func GetProducts(c *fiber.Ctx) error{
	products := []models.Product{}
	database.Database.Db.Find(&products)
	responseProducts := []serializers.ProductSerializer{}
	for _, product := range products {
		responseProduct := serializers.CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}

	return c.Status(200).JSON(responseProducts)
}

func GetProductByID(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err != nil{
		return c.Status(400).JSON(map[string]string{"msg": ":id is invalid"})
	}
	product := models.Product{}
	if err := findProductById(&product, id); err != nil{
		return c.Status(400).JSON(map[string]string{
			"msg": "error with user",
			"error": err.Error(),
		})
	}

	responseProduct := serializers.CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}


func UpdateProduct(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateProduct struct {
		Name string `json:"name"`
		SerialNumber string `json:"serial_number"`
	} 

	product := models.Product{}
	if err := findProductById(&product, id); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var updateProduct UpdateProduct
	if err := c.BodyParser(&updateProduct); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	product.Name = updateProduct.Name
	product.SerialNumber = updateProduct.SerialNumber

	database.Database.Db.Save(&product)

	responseProduct := serializers.CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)

}


func DeleteProduct(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(200).JSON(err.Error())
	}
	var product models.Product
	if err := findProductById(&product, id); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(200).JSON(err.Error())
	}
	return c.Status(200).JSON("Product Deleted")
}