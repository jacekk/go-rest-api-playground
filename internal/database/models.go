package database // import "github.com/jacekk/go-rest-api-playground/internal/database

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/raja/argon2pw"
)

type Post struct {
	gorm.Model
	AuthorID   uint   `gorm:"NOT NULL" validate:"required|uint"`
	Body       string `gorm:"type:text" validate:"required|minLen:10"`
	Category   PostCategory
	CategoryID uint   `json:"-" gorm:"NOT NULL" validate:"required|uint"`
	Title      string `gorm:"NOT NULL" validate:"required|minLen:2|maxLen:255"`
}

type PostCategory struct {
	gorm.Model
	Name string `gorm:"type:varchar(20)" validate:"in:[Uno,Dos,Tres,Cuatro]"` // @todo check whilelist
}

type Comment struct {
	gorm.Model
	Author string `gorm:"type:varchar(80)"`
	Body   string `gorm:"type:text;NOT NULL"`
}

type UserAccount struct {
	gorm.Model
	Email    string `gorm:"type:text;NOT NULL;unique_index" validate:"required|email"` // @todo check unique index
	Name     string `gorm:"type:varchar(50)" validate:"required|minLen:3|maxLen:50"`
	Password string `json:"-" gorm:"type:varchar(90)"`
}

func (self *UserAccount) EncryptPasswordIfSet(plainPass string) {
	if plainPass == "" {
		return
	}

	hashedPassword, err := argon2pw.GenerateSaltedHash(plainPass)

	if err != nil {
		log.Panicf("Hash generated returned error: %v \n", err)
	}

	self.Password = hashedPassword
}
