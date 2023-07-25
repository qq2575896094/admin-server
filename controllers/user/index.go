package user

import (
	"github.com/qq2575896094/admin-server/models"
)

type User interface {
	RegisterHandler() models.HandlerFunc
	LoginHandler() models.HandlerFunc
	LogoutHandler() models.HandlerFunc
	GetUserInfoHandler() models.HandlerFunc
}

type userHandler struct{}

func New() User {
	return &userHandler{}
}
