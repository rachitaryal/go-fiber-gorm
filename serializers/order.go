package serializers

import (
	"github.com/rachitaryal/go-fiber-gorm/models"
)

// type Order struct{
// 	ID uint `json:"id"`
// 	ProductReferer int `json:"product_id"`
// 	Product Product `gorm:"foreignKey:ProductReferer"`
// 	UserReferer int `json:"user_id"`
// 	User User `gorm:"foreignKey:UserReferer"`
// }


func CreateResponseOrder(order models.Order){}