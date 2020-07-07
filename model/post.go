package model

import (
	"github.com/jinzhu/gorm"
)

// Post gorm model struct
type Post struct {
	gorm.Model
	ID    int    `gorm:"AUTO_INCREMENT; primary_key"`
	Title string `gorm:"size:255"`
	Body  string `gorm:"type:text"`
}
