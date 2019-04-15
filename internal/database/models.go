package database // import "github.com/jacekk/go-rest-api-playground/internal/database

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title string `gorm:"NOT NULL" validate:"required|minLen:2|maxLen:255"`
	Body  string `gorm:"type:text" validate:"required|minLen:10"`
}

type Comment struct {
	gorm.Model
	Author string `gorm:"type:varchar(80)"`
	Body   string `gorm:"type:text;NOT NULL"`
}
