package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jacekk/go-rest-api-playground/internal/database"
	"github.com/jacekk/go-rest-api-playground/internal/routing"
	"github.com/joho/godotenv"
)

var projectDir string

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	projectDir = filepath.Dir(filepath.Dir(fileName))

	godotenv.Load(filepath.Join(projectDir, ".env"))     // gitignored; precedence over `dist.env`, but not over system ENVs;
	godotenv.Load(filepath.Join(projectDir, "dist.env")) // defaults; under VCS;
}

func main() {
	fmt.Println("Server starting ...")
	db := database.InitDB(projectDir)
	defer db.Close()
	port := os.Getenv("SERVER_PORT")
	routing.InitRouter(port)
}
