package user

import (
	"github.com/qq2575896094/admin-server/types"
)

type User interface {
	Create() types.HandlerFunc
	Login() types.HandlerFunc
	Logout() types.HandlerFunc
	GetUserInfo() types.HandlerFunc
}

type userHandler struct{}

func New() User {
	return &userHandler{}
}
