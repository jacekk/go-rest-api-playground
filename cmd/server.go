package main // import "github.com/jacekk/go-rest-api-playground"

import (
	"fmt"
	"github.com/jacekk/go-rest-api-playground/internal/routing"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	godotenv.Load("../.env")      // gitignored; precedence over this one:
	godotenv.Load("../.env.dist") // defaults; under VCS
}

func main() {
	fmt.Println("Server starting ...")
	routing.InitRouter(os.Getenv("SERVER_PORT"))
}
