package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	godotenv.Load("../.env")      // gitignored; precedence over this one:
	godotenv.Load("../.env.dist") // defaults; under VCS
}

func main() {
	fmt.Printf("Hello %s!", os.Getenv("NAME"))
}
