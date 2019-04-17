package database // import "github.com/jacekk/go-rest-api-playground/internal/database

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	// Author     UserAccount
	// AuthorID   uint
	Body       string `gorm:"type:text" validate:"required|minLen:10"`
	Category   PostCategory
	CategoryID uint   `json:"-" gorm:"NOT NULL" validate:"required"`
	Title      string `gorm:"NOT NULL" validate:"required|minLen:2|maxLen:255"`
}

type PostCategory struct {
	gorm.Model
	Name string `gorm:"type:varchar(20)" validate:"in:[Uno,Dos,Tres,Cuatro]"`
}

type Comment struct {
	gorm.Model
	Author string `gorm:"type:varchar(80)"`
	Body   string `gorm:"type:text;NOT NULL"`
}

type UserAccount struct {
	gorm.Model
	Email    string `gorm:"type:text;NOT NULL"`
	Name     string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(100)"`
}

func (p *Post) UnmarshalJSON(bytes []byte) error {
	type PostAlias Post
	var category PostCategory

	post := &struct {
		Category string
		*PostAlias
	}{
		PostAlias: (*PostAlias)(p),
	}
	err := json.Unmarshal(bytes, &post)

	if err != nil {
		return err
	}

	DB.FirstOrCreate(&category, PostCategory{Name: post.Category})
	p.CategoryID = category.ID
	p.Category = category

	return nil
}
