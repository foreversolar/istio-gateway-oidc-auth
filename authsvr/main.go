package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func validateTokenWithIDP(token string) bool {
	// 示例逻辑，可以根据实际情况调整
	if len(token) > 5 {
		return true
	} else {
		return false
	}
}

func checkTokenHandler(c *gin.Context) {
	// 打印所有请求头到控制台，用于调试
	log.Println("Request headers:", c.Request.Header)

	// 检查请求头中是否有Authorization字段
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		// 如果Authorization字段存在，处理逻辑
		log.Println("Authorization header found:", authHeader)

		isValid := validateTokenWithIDP(authHeader)
		if isValid {
			c.JSON(http.StatusOK, gin.H{"status": "pass"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Forbidden: Authorization header invalid as " + authHeader})
		}
	} else {
		// 如果Authorization字段不存在，返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Forbidden: Authorization header required"})
	}
}

func main() {
	// 创建一个 Gin 路由器
	router := gin.Default()

	// 定义路由和处理函数
	router.GET("/", checkTokenHandler)

	// 启动 HTTP 服务器
	router.Run(":4000") // 监听并在 0.0.0.0:4000 上启动服务
}
