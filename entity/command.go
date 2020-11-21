package entity

import (
	"gorm.io/gorm"
)

type Command struct {
	gorm.Model
	Title string `gorm:"column:title;type:VARCHAR(255)"`
	Text  string `gorm:"column:text;type:VARCHAR(255)"`
}
