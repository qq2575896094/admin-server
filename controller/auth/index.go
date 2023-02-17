package auth

import "github.com/qq2575896094/admin-server/types"

type Auth interface {
	Create() types.HandlerFunc
}

type authHandler struct{}

func New() Auth {
	return &authHandler{}
}
