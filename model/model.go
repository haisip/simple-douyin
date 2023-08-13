package model

// 导入包
import (
	"time"
)

// User 定义User结构体，对应user表
type User struct {
	ID              int64     `gorm:"primaryKey;autoIncrement;not null" json:"id"`               // 主键、自增、不为空
	Name            string    `gorm:"unique;not null" json:"name"`                               // 唯一索引、不为空
	Password        string    `gorm:"" json:"-"`                                                 // 忽略密码字段
	FollowCount     int32     `gorm:"default:0;not null" json:"follow_count"`                    // 默认0、不为空
	FollowerCount   int32     `gorm:"default:0;not null" json:"follower_count"`                  // 默认0、不为空
	Avatar          string    `gorm:"" json:"avatar,omitempty"`                                  // 如果为空则忽略
	BackgroundImage string    `gorm:"" json:"background_image,omitempty"`                        // 如果为空则忽略
	Signature       string    `gorm:"" json:"signature,omitempty"`                               // 如果为空则忽略
	TotalFavorited  int32     `gorm:"default:0;not null" json:"-"`                               // 忽略总收藏数字段
	WorkCount       int32     `gorm:"default:0;not null" json:"-"`                               // 忽略作品数字段
	FavoriteCount   int32     `gorm:"default:0;not null" json:"-"`                               // 忽略收藏数字段
	CreateAt        time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;not null" json:"-"` // 忽略创建时间字段
	DelFlag         bool      `gorm:"default:0;not null" json:"-"`                               // 忽略删除标志字段

	Videos          []Video     `gorm:"foreignKey:AuthorID" json:"-"` // 忽略视频关联字段
	Comments        []Comment   `gorm:"foreignKey:UserID" json:"-"`   // 忽略评论关联字段
	Followers       []UserUser  `gorm:"foreignKey:Followed" json:"-"` // 忽略粉丝关联字段
	Followeds       []UserUser  `gorm:"foreignKey:Follower" json:"-"` // 忽略关注关联字段
	FavoritedVideos []UserVideo `gorm:"foreignKey:UserID" json:"-"`   // 忽略收藏视频关联字段
}

// Video 定义Video结构体，对应video表
type Video struct {
	ID            int64     `gorm:"primaryKey;autoIncrement;not null"`                // 主键、自增、不为空
	AuthorID      int64     `gorm:"not null"`                                         // 不为空
	PlayURL       string    `gorm:"not null"`                                         // 不为空
	CoverURL      string    `gorm:"not null"`                                         // 不为空
	FavoriteCount int32     `gorm:"default:0;not null"`                               // 默认0、不为空
	CommentCount  int32     `gorm:"default:0;not null"`                               // 默认0、不为空
	Title         string    `gorm:"not null"`                                         // 不为空
	Author        User      `gorm:"foreignKey:AuthorID"`                              // 关联用户表，外键为AuthorID
	CreateAt      time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;not null"` // 默认当前时间
	DelFlag       bool      `gorm:"default:0;not null"`                               // 0正常 1删除

	Comments []Comment `gorm:"foreignKey:VideoID"` // 关联评论表，外键为VideoID
	//FavoritedUsers []UserVideo `gorm:"foreignKey:VideoID"`                // 关联用户视频表，外键为VideoID
}

// Comment 定义Comment结构体，对应comment表
type Comment struct {
	ID       int64     `gorm:"primaryKey;autoIncrement;not null"`                // 主键自增id、不为空
	UserID   int64     `gorm:"not null"`                                         // 外键，不为空
	VideoID  int64     `gorm:"index;not null"`                                   // 外键，索引，不为空
	Content  string    `gorm:"not null"`                                         // 不为空
	CreateAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;not null"` // 默认当前时间
	DelFlag  bool      `gorm:"default:0;not null"`                               // 0正常 1删除

	User  User  `gorm:"foreignKey:UserID"`  // 关联用户表，外键为UserID
	Video Video `gorm:"foreignKey:VideoID"` // 关联视频表，外键为VideoID
}

// UserUser 定义UserUser结构体，对应user_user表
type UserUser struct {
	ID           int64 `gorm:"primaryKey;autoIncrement;not null"` // 主键自增id、不为空
	Follower     int64 `gorm:"uniqueIndex:er_ed_ui;not null"`     // 外键、联合唯一索引er_ed_ui、不为空
	Followed     int64 `gorm:"uniqueIndex:er_ed_ui;not null"`     // 外键、联合唯一索引er_ed_ui、不为空
	FollowerUser User  `gorm:"foreignKey:Follower"`               // 关联用户表，外键为Follower
	FollowedUser User  `gorm:"foreignKey:Followed"`               // 关联用户表，外键为Followed
}

// UserVideo 定义UserVideo结构体，对应user_video表
type UserVideo struct {
	ID      int64 `gorm:"primaryKey;autoIncrement;not null"`  // 主键自增id
	UserID  int64 `gorm:"uniqueIndex:user_video_ui;not null"` // 联合唯一索引user_video_ui、不为空
	VideoID int64 `gorm:"uniqueIndex:user_video_ui;not null"` // 联合唯一索引user_video_ui、不为空

	User  User  `gorm:"foreignKey:UserID"`  // 关联用户表，外键为UserID
	Video Video `gorm:"foreignKey:VideoID"` // 关联视频表，外键为VideoID
}
