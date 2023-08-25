package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"path/filepath"
	"simple-douyin/config"
	"simple-douyin/model"
	"time"
)

var (
	staticBaseUrl = ""
)

func init() {
	configLocal := config.GetConfig()
	staticBaseUrl = configLocal.StaticBaseUrl
}

func Publish(c *gin.Context) {
	// todo 保存视频到静态目录，保存到数据库（新建video字段、user_video、更新user表作品数量）
	currentUserID := c.GetInt64("user_id")
	title := c.PostForm("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "miss title"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	if filepath.Ext(data.Filename) != ".mp4" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "Invalid file format. Only .mp4 files are allowed.",
		})
		return
	}

	finalName := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(data.Filename)) // 使用时间戳防止重复
	savePath := filepath.Join("./public/", finalName)

	if err := c.SaveUploadedFile(data, savePath); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	tx := model.DB.Begin()
	if err := tx.Create(&model.Video{
		AuthorID: currentUserID,
		PlayURL:  staticBaseUrl + finalName,
		CoverURL: "",
		Title:    title,
	}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, "create error")
		return
	}
	if err := tx.Model(&model.User{}).Where("id = ?", currentUserID).Update("work_count", gorm.Expr("work_count + ?", 1)).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, "update error")
		return
	}
	tx.Commit()

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

func PublishList(c *gin.Context) {
	currentUserID, _ := c.Get("user_id")
	targetUserID := c.Query("user_id")
	if targetUserID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "missing query param ot user_id"})
		return
	}

	//videoArr := make([]model.Video, 10)
	var videoArr []model.Video
	if err := model.DB.Table("video").
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.
				Select("user.*, CASE WHEN uu.flag = 1 THEN true ELSE false END AS is_follow").
				Joins("LEFT JOIN user_user AS uu ON  uu.followed = user.id  AND uu.follower = ?", currentUserID)
		}).
		Order("video.create_at DESC").
		Where("video.author_id = ?", targetUserID).
		Debug().
		Find(&videoArr).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed"})
		return
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoArr,
	})
}
