package log

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Params struct {
	Request *http.Request `json:"-"`
	// TimeStamp shows the time after the server returns a response.
	TimeStamp time.Time `json:"-"`
	// StatusCode is HTTP response code.
	StatusCode int `json:"statusCode"`
	// Latency is how much time the server cost to process a certain request.
	Latency time.Duration `json:"latency"`
	// ClientIP equals Context's ClientIP method.
	ClientIP string `json:"clientIP"`
	// Method is the HTTP method given to the request.
	Method string `json:"method"`
	// Path is a path the client requests.
	Path string `json:"path"`
}

func formatter(params Params) string {
	j, _ := json.Marshal(params)
	fmt.Println(string(j), params)
	return string(j)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		params := Params{
			Request: c.Request,
			Latency: time.Now().Sub(start),
		}

		// Stop timer
		params.ClientIP = c.ClientIP()
		params.Method = c.Request.Method
		params.StatusCode = c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		params.Path = path

		logrus.Infof(formatter(params))

	}
}
