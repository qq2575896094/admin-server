package types

import "github.com/gin-gonic/gin"

type Context = gin.Context

type HandlerFunc = gin.HandlerFunc

type Response struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
