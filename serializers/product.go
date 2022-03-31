package serializers

import (
	"github.com/rachitaryal/go-fiber-gorm/models"
)

type ProductSerializer struct{
	ID uint `json:"id"`
	Name string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(product models.Product)ProductSerializer{
	return ProductSerializer{
		ID: product.ID,
		Name: product.Name,
		SerialNumber: product.SerialNumber,
	}
}