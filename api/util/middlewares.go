package util

import (
	"github.com/gin-gonic/gin"
	"github.com/guiacarneiro/eterniza/logger"
	"net/http"
	"time"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		c.Next()

		timeStamp := time.Now()
		latency := timeStamp.Sub(start)
		if path != "/api/v1/contingencia/status" {
			data, exists := c.Get("user")
			if exists {
				switch data.(type) {
				case TokenData:
					user := data.(TokenData)
					logger.LogInfo.Printf("%3d %13v | %15s | %15s | %-7s %#v",
						c.Writer.Status(),
						latency,
						c.ClientIP(),
						user.Login,
						c.Request.Method,
						path)
					return
				}
			}
			logger.LogInfo.Printf("%3d %13v | %15s | %-7s %#v",
				c.Writer.Status(),
				latency,
				c.ClientIP(),
				c.Request.Method,
				path)

		}
	}
}

func UserAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := ExtractTokenData(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, struct {
				Mensagem string `json:"message"`
				Success  bool   `json:"success"`
			}{
				Mensagem: "Unauthorized",
				Success:  false,
			})
			logger.LogWarning.Println("Não autorizado: " + err.Error())
			return
		}
		c.Set("user", *data)
		c.Next()
	}
}

func NoCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Next()
	}
}
