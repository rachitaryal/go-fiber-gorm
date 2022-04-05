package serializers

import (
	"time"

	"github.com/rachitaryal/go-fiber-gorm/models"
)

type OrderSerializer struct{
	ID uint `json:"id"`
	User UserSerializer `json:"user"`
	Product ProductSerializer `json:"product"`
	CreatedAt time.Time `json:"order_date"`
}


func CreateResponseOrder(order models.Order, user UserSerializer, product ProductSerializer) OrderSerializer{
	return OrderSerializer{
		ID: order.ID,
		User: user,
		Product: product,
		CreatedAt: order.CreatedAt,
	}
}