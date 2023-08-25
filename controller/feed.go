package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"simple-douyin/model"
	"strconv"
	"time"
)

var maxVideoNum = 30

func Feed(c *gin.Context) {
	userID, _ := c.Get("user_id")
	lastTimeStr := c.Query("latest_time")

	lastTime, _ := strconv.ParseInt(lastTimeStr, 10, 64)

	videoArr := make([]model.Video, 30)
	if userID == nil {
		query := model.DB.Table("video").Preload("Author").
			Order("video.create_at DESC").
			Limit(maxVideoNum)
		if lastTime > 0 {
			query = query.Where("video.create_at > ?", lastTime)
		}
		if err := query.Find(&videoArr).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed"})
			return
		}
	} else {
		query := model.DB.Table("video").
			Joins("LEFT JOIN user_video AS uv ON video.id = uv.video_id AND uv.user_id = ? AND uv.flag = 1", userID).
			Preload("Author", func(db *gorm.DB) *gorm.DB {
				return db.
					Select("user.*, CASE WHEN uu.flag = 1 THEN true ELSE false END AS is_follow").
					Joins("LEFT JOIN user_user AS uu ON  uu.followed = user.id  AND uu.follower = ?", userID)
			}).
			Select("video.*, uv.flag AS is_favorite").
			Order("video.create_at DESC").
			Limit(maxVideoNum)
		if lastTime > 0 {
			query = query.Where("video.create_at > ?", lastTime)
		}

		if err := query.
			Find(&videoArr).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed"})
			return
		}
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoArr,
		NextTime:  time.Now().Unix(),
	})
}
