package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"path/filepath"
	"simple-douyin/model"
	"time"
)

func Publish(c *gin.Context) {
	currentUserInter, _ := c.Get("user")
	currentUser := currentUserInter.(model.User)
	title := c.PostForm("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "缺少title"})
		return
	}
	fmt.Println(1)
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	fmt.Println(1)

	finalName := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(data.Filename)) // 使用时间戳防止重复
	savePath := filepath.Join("./public/", finalName)

	// todo 写入文件中
	if err := c.SaveUploadedFile(data, savePath); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	tx := model.DB.Begin()

	video := model.Video{
		AuthorID: currentUser.ID,
		PlayURL:  "http://192.168.254.86:8080/static/" + finalName,
		CoverURL: "",
		Title:    title,
	}
	userVideo := model.UserVideo{
		UserID: currentUser.ID,
		Flag:   true,
	}
	if err := tx.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&video).Error; err != nil {
		tx.Rollback()
		return
	}
	userVideo.VideoID = video.ID
	if err := tx.Create(&userVideo).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

func PublishList(c *gin.Context) {
	CurrentUserID, _ := c.Get("user_id")
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "missing query param ot value"})
		return
	}

	// todo 获取用户的全部视频列表
	var videoArr []model.Video
	if err := model.DB.Table("video").
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.
				Select("user.*, CASE WHEN uu.flag = 1 THEN true ELSE false END AS is_follow").
				Joins("LEFT JOIN user_user AS uu ON  uu.followed = user.id  AND uu.follower = ?", CurrentUserID)
		}).
		Order("video.create_at DESC").
		Where("video.author_id = ?", userID).
		Find(&videoArr).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed"})
		return
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoArr,
	})
}
