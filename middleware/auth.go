package middleware

import (
	"net/http"
)
import "github.com/gin-gonic/gin"
import "simple-douyin/utils"

// todo 需要修改认证方式
func authenticateUser(userID uint64) (bool, int64) {
	// todo redis 查询用户
	// todo sql 数据库查询用户
	// todo  用户放到redis
	return true, 0
}

// TokenAuthMiddleware Token验证中间件
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData map[string]interface{}
		err := c.ShouldBindJSON(&requestData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			c.Abort()
			return
		}
		token, ok := requestData["token"].(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Token not provided"})
			c.Abort()
			return
		}

		// todo 需要修改认证方式
		userID, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}

		c.Set("UserID", userID)
		c.Next()
	}
}
