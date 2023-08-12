package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"simple-douyin/config"
	"simple-douyin/model"
	"time"
)

var JwtSecret string

func init() {
	JwtSecret = config.GetConfig().JWTSecret
}

type UserCustomClaims struct {
	ID             int64  `json:"id"`
	Username       string `json:"username"`
	StandardClaims jwt.StandardClaims
}

func (c *UserCustomClaims) Valid() error {
	vErr := new(jwt.ValidationError)
	now := time.Now().Unix()
	if c.StandardClaims.ExpiresAt > 0 {
		if now > c.StandardClaims.ExpiresAt {
			vErr.Inner = fmt.Errorf("token is expired")
			vErr.Errors |= jwt.ValidationErrorExpired
		}
	}
	if vErr.Errors == 0 {
		return nil
	}
	return vErr
}

func GenerateToken(user *model.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Second)
	claims := UserCustomClaims{
		ID:       user.ID,
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims).SignedString([]byte(JwtSecret))
	return token, err
}

func ParseToken(tokenString string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(*UserCustomClaims); ok && token.Valid { // 如果解析成功并且token有效
		return claims.ID, nil // 从claims中获取用户ID
	}
	return 0, err
}
