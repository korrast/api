package main

import (
	"fmt"

	"main/api"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Loading env file")
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}

	fmt.Println("Starting API")

	api.InitApi()
}
