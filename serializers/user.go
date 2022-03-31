package serializers

import (
	"github.com/rachitaryal/go-fiber-gorm/models"
)

type UserSerializer struct {
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func CreateResponseUser(userModel models.User)UserSerializer{
	return UserSerializer{
		ID: userModel.ID,
		FirstName: userModel.FirstName,
		LastName: userModel.LastName,
	}
}