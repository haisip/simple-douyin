package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/model"
	"strconv"
)

type CommentListResponse struct {
	Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment model.Comment `json:"comment,omitempty"`
}

func CommentAction(c *gin.Context) {
	// todo 评论操作
	// todo 将评论添加到对应的数据库
	// todo 视频表中 评论数+1
	actionType := c.Query("action_type")
	videoID := c.Query("video_id")
	videoId, _ := strconv.ParseInt(videoID, 10, 64) // todo 判断是否符合规范，这里没做任何处理
	userID, _ := c.Get("user_id")
	userId := userID.(int64)

	if actionType == "1" {
		// todo 发布评论
		content := c.Query("comment_text") // todo 做参数合法判断
		comment := model.Comment{UserID: userId, VideoID: videoId, Content: content}
		if err := model.DB.Create(&comment).Error; err == nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 0},
				Comment:  comment,
			})
		} else {
			c.JSON(http.StatusBadRequest, err.Error()) // 插入失败
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else if actionType == "2" {
		// todo 删除评论
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusBadRequest, "video_id value error")
		return
	}

}

func CommentList(c *gin.Context) {
	// todo 获取评论列表
	videoID := c.Query("video_id") // todo 判断是否符合规范，这里没做任何处理

	var comments []model.Comment
	if err := model.DB.Preload("User"). // todo 优化异常处理
		Where("video_id = ?", videoID).
		Order("create_at DESC").
		Debug().
		Find(&comments).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: 0},
			CommentList: comments,
		})
	}
}
