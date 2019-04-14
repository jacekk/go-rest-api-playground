package database // import "github.com/jacekk/go-rest-api-playground/internal/database

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Body  string
}

type Comment struct {
	gorm.Model
	Author string
	Body   string
}
