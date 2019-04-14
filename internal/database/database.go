package database // import "github.com/jacekk/go-rest-api-playground/internal/database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"path/filepath"
)

var DB *gorm.DB

func InitDB(projectDir string) *gorm.DB {
	var err error
	dbPath := filepath.Join(projectDir, "./storage/database.sqlite")
	DB, err = gorm.Open("sqlite3", dbPath)

	if err != nil {
		panic("Failed to create or connect to database.")
	}

	DB.AutoMigrate(&Post{}, &Comment{})

	return DB
}
