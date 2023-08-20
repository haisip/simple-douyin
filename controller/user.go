package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"simple-douyin/model"
	"simple-douyin/utils"
	"strconv"
)

type UsnPwdRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Login(c *gin.Context) {
	var loginRequest UsnPwdRequest
	if err := c.ShouldBindQuery(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if result := model.DB.Where(&model.User{Name: loginRequest.Username}).First(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation error"})
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: ""},
		UserId:   user.ID,
		Token:    token,
	})
}

func Register(c *gin.Context) {
	var usnPwdRequest UsnPwdRequest
	if err := c.ShouldBindQuery(&usnPwdRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if err := model.DB.Where(&model.User{Name: usnPwdRequest.Username}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usnPwdRequest.Password), bcrypt.DefaultCost)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
				return
			}

			newUser := model.User{Name: usnPwdRequest.Username, Password: string(hashedPassword)}
			if err := model.DB.Create(&newUser).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
				return
			}

			token, err := utils.GenerateToken(&newUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
				return
			}

			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0, StatusMsg: ""},
				UserId:   newUser.ID,
				Token:    token,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User already exists"})
}

type UserInfoResponse struct {
	Response
	User model.User `json:"user"`
}

func UserInfo(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}
	CurrentUserID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not exist"})
	}

	var user model.User

	if err := model.DB.Table("user").
		Select("user.*, CASE WHEN uu.flag = 1 THEN true ELSE false END AS is_follow").
		Joins("LEFT JOIN user_user AS uu ON uu.followed = ?  AND uu.follower = ?", userID, CurrentUserID).
		Where("user.id = ?", userID).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, Response{StatusCode: 1, StatusMsg: "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, UserInfoResponse{
		Response: Response{StatusCode: 0},
		User:     user,
	})
}
