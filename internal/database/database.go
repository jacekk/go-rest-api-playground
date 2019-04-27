package database

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

func runMigrations() {
	DB.AutoMigrate(&Post{}, &Comment{}, &PostCategory{})
	if !DB.HasTable(&UserAccount{}) {
		DB.CreateTable(&UserAccount{})
	}
}

func InitDB(projectDir string) *gorm.DB {
	var err error
	connType, connData := getDbConfig(projectDir)
	fmt.Printf("Connecting to DB of type '%s' ... \n", connType)
	DB, err = gorm.Open(connType, connData)

	if err != nil {
		msg := "Failed to create or connect to database of type '%s' --> %s"
		panic(fmt.Sprintf(msg, connType, err.Error()))
	}

	runMigrations()

	return DB
}
