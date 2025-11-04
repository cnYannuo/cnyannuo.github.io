package db

import (
	"gorm.io/gorm"
)

// 用户模型
type User struct {
	gorm.Model
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

// 数据库全局变量已在main中定义
