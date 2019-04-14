package main // import "github.com/jacekk/go-rest-api-playground"

import (
	"fmt"
	"github.com/jacekk/go-rest-api-playground/internal/routing"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	rootDir := filepath.Dir(fileName)

	godotenv.Load(filepath.Join(rootDir, "../.env"))      // gitignored; precedence over this one:
	godotenv.Load(filepath.Join(rootDir, "../.env.dist")) // defaults; under VCS
}

func main() {
	fmt.Println("Server starting ...")
	routing.InitRouter(os.Getenv("SERVER_PORT"))
}
