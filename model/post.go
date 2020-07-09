package model

import (
	"github.com/jinzhu/gorm"
)

// Post gorm model struct
type Post struct {
	gorm.Model
	Title string `gorm:"size:255"`
	Body  string `gorm:"type:text"`
}
