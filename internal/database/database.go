package database // import "github.com/jacekk/go-rest-api-playground/internal/database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"path/filepath"
)

var DB *gorm.DB

func getDbConfig(projectDir string) (connType string, connData string) {
	host := os.Getenv("PGHOST")

	if host == "" {
		connType = "sqlite3"
		connData = filepath.Join(projectDir, "./storage/database.sqlite")
		return
	}

	user := os.Getenv("PGUSER")
	dbName := os.Getenv("PGDATABASE")
	pass := os.Getenv("PGPASS")
	dataTpl := "host=%s user=%s dbname=%s password=%s sslmode=disable"
	connData = fmt.Sprintf(dataTpl, host, user, dbName, pass)
	connType = "postgres"
	return
}

func InitDB(projectDir string) *gorm.DB {
	var err error
	connType, connData := getDbConfig(projectDir)
	DB, err = gorm.Open(connType, connData)

	if err != nil {
		msg := "Failed to create or connect to database of type '%s' --> %s"
		panic(fmt.Sprintf(msg, connType, err.Error()))
	}

	DB.AutoMigrate(&Post{}, &Comment{})

	return DB
}
