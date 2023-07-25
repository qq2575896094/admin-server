package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 定义系统返回报文格式
type Response struct {
	StatusCode int    `json:"statusCode"`
	Data       any    `json:"data"`
	Message    string `json:"message"`
}

const (
	SuccessCode            = 200
	SuccessMessage         = "success"
	ErrorCode              = 0
	ErrorAuthorizationCode = 401
	ErrorMessage           = "fail"
)

func Resp(statusCode int, data any, message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		statusCode,
		data,
		message,
	})
}

// OkDefault 成功, 无返回数据
func OkDefault(c *gin.Context) {
	Resp(SuccessCode, map[string]interface{}{}, SuccessMessage, c)
}

// Ok 成功, 有返回数据
func Ok(data any, c *gin.Context) {
	Resp(SuccessCode, data, SuccessMessage, c)
}

// Fail 失败, 返回失败信息
func Fail(statusCode int, message string, c *gin.Context) {
	Resp(SuccessCode, map[string]interface{}{}, message, c)
}

// FailDefault Fail 失败, 无返回数据
func FailDefault(c *gin.Context) {
	Resp(ErrorCode, map[string]interface{}{}, ErrorMessage, c)
}

func FailAuthorization(message string, c *gin.Context) {
	Resp(ErrorAuthorizationCode, map[string]interface{}{}, message, c)
}

// FailWithMessage 失败, 返回失败信息
func FailWithMessage(message string, c *gin.Context) {
	Resp(ErrorCode, map[string]interface{}{}, message, c)
}
