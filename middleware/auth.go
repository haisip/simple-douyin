package middleware

import "net/http"
import "github.com/gin-gonic/gin"

// 模拟一个简单的token验证函数
// todo 需要修改认证方式
func authenticateToken(token string) bool {
	// todo redis 查询用户
	// todo sql 数据库查询用户
	// todo  用户放到redis
	return true
}

// TokenAuthMiddleware Token验证中间件
// todo 需要修改认证方式
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
		flag := authenticateToken(token)
		if !flag {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}
		c.Next()
	}
}
