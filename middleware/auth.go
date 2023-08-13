package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/model"
	"simple-douyin/utils"
)

// todo 需要修改认证方式
func authenticateUser(userID int64) (model.User, error) {
	// todo redis 查询用户
	// todo sql 数据库查询用户
	// todo  用户放到redis
	var user = model.User{ID: userID}
	return user, nil
}

// TokenAuthMiddleware Token验证中间件
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			var requestData struct {
				Token string `json:"token"`
			}
			if err := c.ShouldBindJSON(&requestData); err != nil || requestData.Token == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing token"})
				c.Abort()
				return
			}
			token = requestData.Token
		}

		userID, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		user, err := authenticateUser(userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Set("user", user)
		c.Next()
	}
}
