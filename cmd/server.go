package main // import "github.com/jacekk/go-rest-api-playground"

import (
	"fmt"
	"github.com/jacekk/go-rest-api-playground/internal/database"
	"github.com/jacekk/go-rest-api-playground/internal/routing"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"runtime"
)

var projectDir string

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	projectDir = filepath.Dir(filepath.Dir(fileName))

	godotenv.Load(filepath.Join(projectDir, ".env"))      // gitignored; precedence over this one:
	godotenv.Load(filepath.Join(projectDir, ".env.dist")) // defaults; under VCS
}

func main() {
	fmt.Println("Server starting ...")
	db := database.InitDB(projectDir)
	defer db.Close()
	port := os.Getenv("SERVER_PORT")
	routing.InitRouter(port)
}
