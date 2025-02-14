package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	Id       int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Role     string `gorm:"column:role;not null" json:"role"`
	Username string `gorm:"column:username;not null" form:"username" json:"username" binding:"required"`
	Password string `gorm:"column:password;not null" form:"password" json:"password" binding:"required"`
	Email    string `gorm:"column:email;not null" form:"email" json:"email"`
	Phone    string `gorm:"column:phone;not null" form:"phone" json:"phone"`
	// 外键。定义与Content模型的关联，表示一个用户可以有多个内容
	Contents []Content `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type ChangePasswordRequest struct {
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type Content struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"` // 标识每条留言，防止出现相同内容留言难以辨认的情况
	UserID    int       `gorm:"column:user_id;not null"`                      // 添加UserID字段作为外键,保证每个Content关联一个User
	User      User      `gorm:"-" json:"user"`                                // 不保存User到content表中
	Content   string    `gorm:"column:content;not null" form:"content" json:"content"`
	Timestamp time.Time `gorm:"column:timestamp;default:CURRENT_TIMESTAMP;type:datetime;not null"` // 打上时间戳
}

type MyClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
