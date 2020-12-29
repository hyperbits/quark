package models

import (
	"github.com/jinzhu/gorm"
)

type Template struct {
	gorm.Model
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `gorm:"type:text" json:"content"`
}
