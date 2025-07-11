package main

import (
	"fmt"
	"log"

	"main/api"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Loading env file")
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}

	fmt.Println("Starting API")

	server, err := api.NewServer()
	if err != nil {
		log.Fatal("Failed to create server:", err)
	}
	defer server.Close()

	if err := server.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
