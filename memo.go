package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Memo struct {
	ID        int `gorm:"primaryKey"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time

	UserUUID string `gorm:"index;not null"`
	//User     User   `gorm:"foreignKey:UserUUID;references:UUID;constraint:OnDelete:CASCADE"`
}

func Show_memos(c *gin.Context) {
	println("/memos")
	var memos []Memo
	user := GetProfile(c)
	if (User{}) == user {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ログインしてください"})
		return
	} else {
		db.Where("user_uuid = ?", user.UUID).Find(&memos)
		c.JSON(http.StatusOK, gin.H{"memos": memos})
	}
}

func Create_memo(c *gin.Context) {
	println("/add_memo")
	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの解析に失敗しました"})
		return
	}
	user := GetProfile(c)
	if (User{}) == user {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ログインしてください"})
		return
	} else {
		fmt.Println(user)
		println(req.Title)
		println(req.Content)
		db.Create(&Memo{Title: req.Title, Content: req.Content, UserUUID: user.UUID})
		c.JSON(http.StatusOK, gin.H{"message": "メモを作成しました"})
	}
}
